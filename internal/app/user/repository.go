package user

import (
	"context"

	"github.com/alfianvitoanggoro/go-postgres/db"
)

func GetAll(ctx context.Context) ([]User, error) {
	rows, err := db.Pool.Query(ctx, "SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func Create(ctx context.Context, user User) error {
	_, err := db.Pool.Exec(ctx, "INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}
