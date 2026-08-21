package userpg

import "github.com/chishkin-afk/interval-learning-algo/backend/internal/modules/auth/domain/user"

func toRecord(user *user.User) *userRecord {
	return &userRecord{
		ID: user.ID(),
		Username: user.Username(),
		Email: user.Email().String(),
		PasswordHash: user.PasswordHash().String(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}
}

func toUser(record *userRecord) *user.User {
	return user.From(
		record.ID,
		record.Username,
		user.Email(record.Email),
		user.PasswordHash(record.PasswordHash),
		record.CreatedAt,
		record.UpdatedAt,
	)
}