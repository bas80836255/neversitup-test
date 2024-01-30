package user

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type RepositoryTest struct {
	suite.Suite
	repository *repository
	controller *gomock.Controller
}

func (t *RepositoryTest) SetupTest() {
	t.controller = gomock.NewController(t.Suite.T())

	adminUser := User{
		ID:         0,
		Name:       "admin",
		CreateDate: time.Now(),
		CreateBy:   "system",
	}

	user1 := User{
		ID:         1,
		Name:       "user1",
		CreateDate: time.Now(),
		CreateBy:   "system",
	}
	t.repository = &repository{
		mapUsers: map[string]User{"0": adminUser, "1": user1},
	}
}

func (t *RepositoryTest) TearDownTest() {
	t.controller.Finish()
	t.repository = nil
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTest))
}

func (t *RepositoryTest) Test_FindOne_WhenFoundUser_ThenReturnUserAndNil() {
	userID := "0"

	result, err := t.repository.FindOne(context.Background(), userID)

	t.Nil(err)
	t.Equal(t.repository.mapUsers[userID], result)
}

func (t *RepositoryTest) Test_FindOne_WhenNotFoundUser_ThenReturnEmptyAndError() {
	userID := "3"

	_, err := t.repository.FindOne(context.Background(), userID)

	t.NotNil(err)
}

func (t *RepositoryTest) Test_NewRepository_WhenCall_ThenReturnCounter() {
	repo := &repository{}
	t.IsType(NewRepository(), repo)
}
