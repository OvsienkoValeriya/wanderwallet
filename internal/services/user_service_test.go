package services

import (
	"errors"
	"os"
	"testing"
	"wanderwallet/internal/mocks"
	"wanderwallet/internal/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	// Устанавливаем тестовый SECRET_KEY
	os.Setenv("SECRET_KEY", "test-secret-key-for-testing")
	code := m.Run()
	os.Exit(code)
}

func TestUserService_Register_Success(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	service := NewUserService(mockRepo)

	login := "testuser"
	password := "testpass123"

	// Настраиваем мок с помощью gomock
	mockRepo.EXPECT().IsLoginExists(login).Return(false, nil)
	mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil).Do(func(user *models.User) {
		user.ID = 1 // Имитируем автоинкремент ID
	})

	// Act
	response, err := service.Register(login, password)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, login, response.User.Login)
	assert.NotEmpty(t, response.Token)
	assert.NotEqual(t, password, response.User.Password) // Пароль должен быть захеширован
}

func TestUserService_Register_UserAlreadyExists(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	service := NewUserService(mockRepo)

	login := "existinguser"
	password := "testpass123"

	// Настраиваем мок - пользователь уже существует
	mockRepo.EXPECT().IsLoginExists(login).Return(true, nil)

	// Act
	response, err := service.Register(login, password)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrUserAlreadyExists, err)
	assert.Nil(t, response)
}

func TestUserService_Register_DatabaseError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	service := NewUserService(mockRepo)

	login := "testuser"
	password := "testpass123"
	dbError := errors.New("database connection error")

	// Настраиваем мок
	mockRepo.EXPECT().IsLoginExists(login).Return(false, dbError)

	// Act
	response, err := service.Register(login, password)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, dbError, err)
	assert.Nil(t, response)
}

func TestUserService_Login_Success(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	service := NewUserService(mockRepo)

	login := "testuser"
	password := "testpass123"

	// Создаем пользователя с захешированным паролем
	hashedPassword, _ := service.hashPassword(password)
	user := &models.User{
		Login:    login,
		Password: hashedPassword,
	}
	user.ID = 1

	// Настраиваем мок
	mockRepo.EXPECT().GetByLogin(login).Return(user, nil)

	// Act
	response, err := service.Login(login, password)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, user, response.User)
	assert.NotEmpty(t, response.Token)
}

func TestUserService_Login_UserNotFound(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	service := NewUserService(mockRepo)

	login := "nonexistentuser"
	password := "testpass123"

	// Настраиваем мок
	mockRepo.EXPECT().GetByLogin(login).Return(nil, gorm.ErrRecordNotFound)

	// Act
	response, err := service.Login(login, password)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrUserNotFound, err)
	assert.Nil(t, response)
}

func TestUserService_Login_InvalidPassword(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	service := NewUserService(mockRepo)

	login := "testuser"
	correctPassword := "correctpass123"
	wrongPassword := "wrongpass123"

	// Создаем пользователя с правильным паролем
	hashedPassword, _ := service.hashPassword(correctPassword)
	user := &models.User{
		Login:    login,
		Password: hashedPassword,
	}
	user.ID = 1

	// Настраиваем мок
	mockRepo.EXPECT().GetByLogin(login).Return(user, nil)

	// Act
	response, err := service.Login(login, wrongPassword)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidPassword, err)
	assert.Nil(t, response)
}

func TestUserService_GetUserByID_Success(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	service := NewUserService(mockRepo)

	userID := uint(1)
	expectedUser := &models.User{
		Login: "testuser",
	}
	expectedUser.ID = userID

	// Настраиваем мок
	mockRepo.EXPECT().GetByID(userID).Return(expectedUser, nil)

	// Act
	user, err := service.GetUserByID(userID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}
