package user

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ServiceTest struct {
	suite.Suite
	service    *service
	controller *gomock.Controller
	repository *MockRepository
}

func (t *ServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.Suite.T())
	t.repository = NewMockRepository(t.controller)
	t.service = &service{
		repository: t.repository,
	}
}

func (t *ServiceTest) TearDownTest() {
	t.controller.Finish()
	t.repository = nil
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTest))
}

func (t *ServiceTest) Test_GetUser_WhenRepositorySuccess_ThenReturnResult() {
	userID := "1"
	user := User{
		Name: "name",
	}

	t.repository.EXPECT().FindOne(gomock.Any(), userID).Return(user, nil)

	result, err := t.service.GetUser(context.Background(), userID)

	t.Nil(err)
	t.Equal(user, result)
}

func (t *ServiceTest) Test_GetUser_WhenRepositoryError_ThenReturnError() {
	userID := "1"

	t.repository.EXPECT().FindOne(gomock.Any(), userID).Return(User{}, errors.New("db error"))

	_, err := t.service.GetUser(context.Background(), userID)

	t.NotNil(err)
}

func (t *ServiceTest) Test_NewService_WhenCall_ThenReturnCounter() {
	service := &service{}
	t.IsType(NewService(t.repository), service)
}
