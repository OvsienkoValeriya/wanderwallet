package services

import (
	"errors"
	"os"
	"strconv"
	"time"
	"wanderwallet/internal/models"
	"wanderwallet/internal/repository"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrTokenGeneration   = errors.New("failed to generate token")
)

type UserService struct {
	userRepo repository.UserRepositoryInterface
}

func NewUserService(userRepo repository.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

type RegisterResponse struct {
	User  *models.User
	Token string
}

func (s *UserService) Register(login, password string) (*RegisterResponse, error) {
	exists, err := s.userRepo.IsLoginExists(login)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrUserAlreadyExists
	}

	hashedPassword, err := s.hashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Login:    login,
		Password: hashedPassword,
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, err
	}

	token, err := s.generateToken(user.ID)

	if err != nil {
		return nil, ErrTokenGeneration
	}

	return &RegisterResponse{
		User:  user,
		Token: token,
	}, nil
}

func (s *UserService) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *UserService) generateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(int(userID)),
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func (s *UserService) comparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

type LoginResponse struct {
	User  *models.User
	Token string
}

func (s *UserService) Login(login, password string) (*LoginResponse, error) {
	user, err := s.userRepo.GetByLogin(login)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	if err := s.comparePassword(user.Password, password); err != nil {
		return nil, ErrInvalidPassword
	}

	token, err := s.generateToken(user.ID)
	if err != nil {
		return nil, ErrTokenGeneration
	}

	return &LoginResponse{
			User:  user,
			Token: token,
		},
		nil
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.GetByID(id)
}
