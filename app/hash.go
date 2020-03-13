package app

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Hasher interface {
	GenerateHash(password string) (string, error)
	ValidateHashPassword(hashedPassword string, password string) error
}

type BcryptHash struct {
	cost int
}

var _ Hasher = &BcryptHash{}

func NewBcryptHasher(cost int) Hasher {
	return &BcryptHash{cost}
}

func (b *BcryptHash) GenerateHash(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", fmt.Errorf("generate hash failed: %w", err)
	}
	return string(hashPass), nil
}

func (b *BcryptHash) ValidateHashPassword(hashedPassword string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return fmt.Errorf("validate hash failed: %w", err)
	}
	return nil
}
