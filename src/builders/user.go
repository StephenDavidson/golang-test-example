package builders

import (
	"github.com/bxcodec/faker/v3"
)

type UserBuilder struct {
	Email              string  `faker:"email"`
	FirstName          string  `faker:"first_name"`
	Value			   string   `faker:"uuid_hyphenated"`
}

func GenerateUser() UserBuilder {
	userBuilder := UserBuilder{}
	faker.FakeData(&userBuilder)
	return userBuilder
}