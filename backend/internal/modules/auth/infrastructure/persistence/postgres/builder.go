package userpg

import (
	"uuid"

	"github.com/chishkin-afk/interval-learning-algo/backend/internal/modules/auth/domain/user"
)

var (
	allUserColumns = []string{
		"id",
		"username",
		"email",
		"password_hash",
		"created_at",
		"updated_at",
	}
)

func (ur *userRepository) buildInsert(record *userRecord) (string, []any, error) {
	return ur.sb.Insert("users").Columns(allUserColumns...).Values(
		record.ID,
		record.Username,
		record.Email,
		record.PasswordHash,
		record.CreatedAt,
		record.UpdatedAt,
	).ToSql()
}

func (ur *userRepository) buildSelectCond[T uuid.UUID | user.Email](cond T) (string, []any, error) {
	builder := ur.sb.Select(allUserColumns...).From("users").Limit(1)
	switch v := any(cond).(type) {
	case uuid.UUID:
		builder = builder.Where("id = ?", v)
	case user.Email:
		builder = builder.Where("email = ?", v)
	}

	return builder.ToSql()
}

func (ur *userRepository) buildSelectForUpdate(id uuid.UUID) (string, []any, error) {
	return ur.sb.Select(allUserColumns...).Where("id = ?", id).Suffix("FOR UPDATE").ToSql()
}

func (ur *userRepository) buildUpdate(updRecord *userRecord) (string, []any, error) {
	return ur.sb.Update("users").SetMap(map[string]any{
		"username":   updRecord.Username,
		"updated_at": updRecord.UpdatedAt,
	}).ToSql()
}

func (ur *userRepository) buildDelete(id uuid.UUID) (string, []any, error) {
	return ur.sb.Delete("users").Where("id = ?", id).ToSql()
}
