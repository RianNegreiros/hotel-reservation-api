package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost      = 12
	minPassLen      = 6
	maxPassLen      = 72
	minFirstNameLen = 2
	minLastNameLen  = 2
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params CreateUserParams) Validate() []string {
	errors := []string{}

	if len(params.Password) < minPassLen {
		errors = append(errors, fmt.Sprintf("password must be at least %d characters", minPassLen))
	}

	if len(params.Password) > maxPassLen {
		errors = append(errors, fmt.Sprintf("password must be at most %d characters", maxPassLen))
	}

	if len(params.FirstName) < minFirstNameLen {
		errors = append(errors, fmt.Sprintf("firstName must be at least %d characters", minFirstNameLen))
	}

	if len(params.LastName) < minLastNameLen {
		errors = append(errors, fmt.Sprintf("lastName must be at least %d characters", minLastNameLen))
	}

	if !isEmailValid(params.Email) {
		errors = append(errors, fmt.Sprintf("email is invalid"))
	}

	return errors
}

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(enc),
	}, nil
}
