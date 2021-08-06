/*
Melakukan mapping dari struct input ke struct User
*/

package user

import "golang.org/x/crypto/bcrypt"

// menyimpan method / busines logic User
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
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
