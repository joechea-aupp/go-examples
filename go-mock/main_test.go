package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUser(id int) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}

func TestUserService_GetUser(t *testing.T) {
	assert := assert.New(t)

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		userService := &UserService{repo: mockRepo}

		mockUser := &User{ID: 1, Name: "John Doe"}
		mockRepo.On("GetUser", 1).Return(mockUser, nil)

		user, err := userService.GetUser(1)
		assert.NoError(err)
		assert.Equal(mockUser, user)
		mockRepo.AssertExpectations(t)
	})
}
