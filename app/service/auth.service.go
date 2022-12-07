package service

import (
	"context"
	"fiber/app/model"
	"fiber/pkg/util"
	"fiber/platform/db"
	"fiber/postgres"
	"fmt"

	"github.com/google/uuid"
)

var invalidCred = "Invalid email or password"

func Login(loginData model.Login) (string, error) {
	// fetch user details based on email
	user, err := db.Query().GetUser(context.Background(), loginData.Email)

	// if nothing is returned it means the user does not exist
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", fmt.Errorf("%v", invalidCred)
		}
	}

	password := loginData.Password
	if password := util.CheckPasswordHash(password, user.Password); password {
		// generating token for user
		token, err := util.GenerateToken(postgres.User{
			ID:    user.ID,
			Email: user.Email,
		})
		if err != nil {
			return "", err
		}
		return token, err
	} else {
		return "", fmt.Errorf(invalidCred)
	}
}

func Register(user postgres.AddUserParams) (string, error) {
	userByEmail, _ := GetUser(user.Email)
	if userByEmail != (postgres.User{}) {
		return "", fmt.Errorf("email already registered")
	}

	// hashing user password
	err := util.HashPasswordParser(&user.Password)
	if err != nil {
		return "", err
	}

	// adding the user to database
	u, err := AddUser(user)
	if err != nil {
		return "", err
	}

	// generate token
	token, err := util.GenerateToken(u)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetAllUsers() ([]postgres.User, error) {
	return db.Query().GetUsers(context.Background())
}

func GetUser(identity string) (postgres.User, error) {
	return db.Query().GetUser(context.Background(), identity)
}

func AddUser(user postgres.AddUserParams) (postgres.User, error) {
	return db.Query().AddUser(context.Background(), user)
}

func GetUserById(id uuid.UUID) (postgres.GetUserByIdRow, error) {
	return db.Query().GetUserById(context.Background(), id)
}
