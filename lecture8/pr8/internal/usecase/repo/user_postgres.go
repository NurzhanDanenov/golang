package repo

import (
	"context"

	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/evrone/go-clean-template/pkg/postgres"
)

const UsersDBName = "users"

type UserRepo struct {
	*postgres.Postgres
	// DB      *sql.DB // Your database connection
	// Builder *sqlbuilder.PostgreSQL
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (ur *UserRepo) GetUsers(ctx context.Context) (users []*entity.User, err error) {
	res := ur.DB.WithContext(ctx).Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
	// sql, _, err := ur.Builder.
	// 	Select("id, name, email, age")
	// 	From("users")
	// 	ToSql()
	// if err != nil : nil, fmt
}

func (ur *UserRepo) CreateUser(ctx context.Context, user *entity.User) (int, error) {
	res := ur.DB.WithContext(ctx).Create(user)
	if res.Error != nil {
		return 0, res.Error
	}
	return user.Id, nil
	// sql, args, err := ur.Builder.Insert(UsersDBName).Columns("name", "email", "age", "password").
	// 	Values(user.Name, user.Email, user.Age, user.Password).
	// 	Suffix("returning id").ToSql()
	// if err != nil {
	// 	return 0, err
	// }

	// var insertedID int

	// err = ur.Pool.QueryRow(ctx, sql, args).Scan(&insertedID)
	// if err != nil {
	// 	return 0, err
	// }

	// return insertedID, nil
}

func (ur *UserRepo) GetUserByEmail(ctx context.Context, email string) (user *entity.User, err error) {
	res := ur.DB.Where("email = ?", email).WithContext(ctx).Find(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (ur *UserRepo) GetUserByID(ctx context.Context, id string) (user *entity.User, err error) {
	res := ur.DB.WithContext(ctx).Where("id = ?", id).Find(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}
