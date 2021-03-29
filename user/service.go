package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	// method register user, mwakili bisnis logic
	RegisterUser(input RegisteruserInput) (User, error)

	// method untuk login setelah itu kita akses dengan func
	Login(input LoginInput) (User, error)

	// untuk cek email available
	IsEmailAvailable(input CheckEmailInput) (bool, error)
}

type service struct {
	// akan mapping struct input ke struck User
	repository Repository //ambil interface repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisteruserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	// kita cari email pengguna
	user, err := s.repository.FindByEmail(email)
	// jika ada error
	if err != nil {
		return user, err
	}
	// jika id tidak ditemukan
	if user.ID == 0 {
		return user, errors.New("User not found")
	}
	// jika tidak error dan ditemukan ID nya
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil

}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return true, nil
}
