package usecase

import (
	"errors"
	"time"

	"github.com/alexander-pastana/go-api-lab/model"
	"github.com/alexander-pastana/go-api-lab/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repository repository.UserRepository
	secretKey string
}

func NewUserUseCase(repo repository.UserRepository, secretKey string) UserUseCase {
	return UserUseCase{
		repository: repo,
		secretKey: secretKey,
	}
}

func (uu *UserUseCase) SignUp(user model.User) error {
	// 1. Valida se o usuário já existe
	userExists, err := uu.repository.GetUserByName(user.Name)
	if err != nil {
		return err
	}
	if userExists != nil {
		return errors.New("usuário já cadastrado")
	}

	// 2. Criptografa a senha do usuário
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 3. Substitui a senha limpa pela hash e salva no banco
	user.Password = string(passwordHash)

	err = uu.repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (uu *UserUseCase) SignIn(user model.User) (string, error) {
	userExists, err := uu.repository.GetUserByName(user.Name)
	if err != nil {
		return "", err
	}
	if userExists == nil {
		return "", errors.New("usuário ou senha inválidos")
	}

	//Verifica se a hash passada é a cadastrada no banco
	err = bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(user.Password))
		if err != nil {
		return "", errors.New("usuário ou senha inválidos")
	}

	//registros de claims utilizados para assinar o token jwt
	claim := jwt.MapClaims{
		"sub": userExists.ID, // Identifica o usuário
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Data atual + 24 horas, convertida em segundos
	}

	//envelopar esses claims (metodo de assinatura, claims registrados)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(uu.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (uu *UserUseCase) ValidateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(uu.secretKey), nil
	})
	if err != nil {
		return err
	}

	return nil
}