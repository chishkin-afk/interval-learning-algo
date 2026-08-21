package userpg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"uuid"

	"github.com/Masterminds/squirrel"
	"github.com/chishkin-afk/interval-learning-algo/backend/internal/modules/auth/domain/user"
	"github.com/lib/pq"
)

// DB defines the interface for database operations used by the repository.
// It abstracts over *sql.DB and *sql.Tx to allow seamless transaction support.
type DB interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

// TxDB extends DB with transaction management capabilities.
type TxDB interface {
	DB
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type scanner interface {
	Scan(dest ...any) error
}

type userRepository struct {
	log *slog.Logger
	db  TxDB
	sb  squirrel.StatementBuilderType
}

// New creates a new userRepository instance with the given logger and database connection.
func New(log *slog.Logger, db TxDB) *userRepository {
	return &userRepository{
		log: log,
		db:  db,
		sb:  squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

// Save persists a new user into the database.
// Returns user.ErrAlreadyExists if a user with the same unique constraints already exists.
func (ur *userRepository) Save(ctx context.Context, user *user.User) error {
	ur.log.Debug("saving user into db",
		slog.String("user_id", user.ID().String()),
	)

	if err := ur.save(ctx, ur.db, user); err != nil {
		return fmt.Errorf("can't save user: %w",
			handleError(err),
		)
	}

	return nil
}

func (ur *userRepository) save(ctx context.Context, db DB, user *user.User) error {
	query, args, err := ur.buildInsert(toRecord(user))
	if err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

// GetByID retrieves a user by their UUID.
// Returns user.ErrNotFound if no user with the given ID exists.
func (ur *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	ur.log.Debug("getting user from db by id",
		slog.String("user_id", id.String()),
	)

	user, err := ur.get(ctx, ur.db, id)
	if err != nil {
		return nil, fmt.Errorf("can't get user by id: %w",
			handleError(err),
		)
	}

	return user, nil
}

// GetByEmail retrieves a user by their email address.
// Returns user.ErrNotFound if no user with the given email exists.
func (ur *userRepository) GetByEmail(ctx context.Context, email user.Email) (*user.User, error) {
	ur.log.Debug("getting user from db by email",
		slog.String("email", email.String()),
	)

	user, err := ur.get(ctx, ur.db, email)
	if err != nil {
		return nil, fmt.Errorf("can't get user by email: %w",
			handleError(err),
		)
	}

	return user, nil
}

func (ur *userRepository) get[T uuid.UUID | user.Email](ctx context.Context, db DB, cond T) (*user.User, error) {
	query, args, err := ur.buildSelectCond(cond)
	if err != nil {
		return nil, err
	}

	return scanRow(db.QueryRowContext(ctx, query, args))
}

// Update applies the provided update function to a user within a read-committed transaction.
// It uses SELECT FOR UPDATE to prevent concurrent modifications.
// Returns the updated user or an error if the user is not found or the update fails.
func (ur *userRepository) Update(ctx context.Context, updFunc user.UpdateFunc, id uuid.UUID) (*user.User, error) {
	tx, rollback, err := ur.beginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't begin tx: %w", err)
	}
	defer rollback()

	updUser, err := ur.getForUpdate(ctx, tx, id)
	if err != nil {
		return nil, fmt.Errorf("can't get user for update: %w",
			handleError(err),
		)
	}

	if err := updFunc(updUser); err != nil {
		return nil, fmt.Errorf("can't call update func: %w", err)
	}

	ur.log.Debug("updating user in db",
		slog.String("user_id", id.String()),
	)

	if err := ur.update(ctx, tx, updUser); err != nil {
		return nil, fmt.Errorf("can't update user: %w",
			handleError(err),
		)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("can't commit tx: %w", err)
	}

	return updUser, nil
}

func (ur *userRepository) update(ctx context.Context, db DB, updUser *user.User) error {
	query, args, err := ur.buildUpdate(toRecord(updUser))
	if err != nil {
		return err
	}

	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	if count, _ := result.RowsAffected(); count == 0 {
		return user.ErrNotFound
	}

	return nil
}

func (ur *userRepository) getForUpdate(ctx context.Context, db DB, id uuid.UUID) (*user.User, error) {
	query, args, err := ur.buildSelectForUpdate(id)
	if err != nil {
		return nil, err
	}

	return scanRow(db.QueryRowContext(ctx, query, args...))
}

// Delete removes a user from the database by their UUID.
// Returns user.ErrNotFound if no user with the given ID exists.
func (ur *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	ur.log.Debug("deleting user from db",
		slog.String("user_id", id.String()),
	)

	if err := ur.delete(ctx, ur.db, id); err != nil {
		return fmt.Errorf("can't delete user from db: %w",
			handleError(err),
		)
	}

	return nil
}

func (ur *userRepository) delete(ctx context.Context, db DB, id uuid.UUID) error {
	query, args, err := ur.buildDelete(id)
	if err != nil {
		return err
	}

	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	if count, _ := result.RowsAffected(); count == 0 {
		return user.ErrNotFound
	}

	return nil
}

func (ur *userRepository) beginTx(ctx context.Context) (*sql.Tx, func(), error) {
	tx, err := ur.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, nil, err
	}

	return tx, func() {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			ur.log.Error("can't rollback tx",
				slog.String("error", err.Error()),
			)
		}
	}, nil
}

func scanRow(scnr scanner) (*user.User, error) {
	var record userRecord
	if err := scnr.Scan(
		&record.ID,
		&record.Username,
		&record.Email,
		&record.PasswordHash,
		&record.CreatedAt,
		&record.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return toUser(&record), nil
}

func handleError(err error) error {
	switch {
	case errors.Is(err, context.DeadlineExceeded),
		errors.Is(err, context.Canceled):
		return err
	case errors.Is(err, sql.ErrNoRows):
		return user.ErrNotFound
	}

	if err, ok := errors.AsType[*pq.Error](err); ok {
		switch err.Code {
		case "23505":
			return user.ErrAlreadyExists
		}
	}

	return err
}
