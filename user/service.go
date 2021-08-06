/*
Melakukan mapping dari struct input ke struct User
*/

package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// menyimpan method / busines logic User
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
}

// ketergantungan / referensi ke repository
// memanggil repository
type service struct {
	repository Repository
}

// bikin obj service passing repository
func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	// mapping struc input ke struct User
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHas, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHas)
	user.Role = "user" // hard coded

	// simpan struct User melalui repository
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) LoginUser(input LoginUserInput) (User, error) {
	userLogin, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return userLogin, err
	}

	if userLogin.ID == 0 {
		return userLogin, errors.New("no user found on that email")
	}

	// check pw
	err = bcrypt.CompareHashAndPassword([]byte(userLogin.PasswordHash), []byte(input.Password))
	if err != nil {
		return userLogin, err
	}

	return userLogin, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	// belum ada user yang sesuai email input
	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	// dapatkan user berdasarkan ID
	// user update attributes file name
	// simpan perubahan avatar file name

	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, err
	}

	// rubah field avatar
	user.AvatarFileName = fileLocation

	// simpan ke db
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}
