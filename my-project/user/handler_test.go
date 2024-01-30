package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HandlerTest struct {
	suite.Suite
	controller *gomock.Controller
	handler    *handler
	service    *MockService
}

func (t *HandlerTest) SetupTest() {
	t.controller = gomock.NewController(t.Suite.T())
	t.service = NewMockService(t.controller)

	t.handler = &handler{
		service: t.service,
	}
}

func (t *HandlerTest) TearDownTest() {
	t.controller.Finish()
	t.handler = nil
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTest))
}

func (t *HandlerTest) TestHandler_Get_WhenServiceFoundUser_ThenReturnHttp200AndUser() {
	recorder := httptest.NewRecorder()
	userID := "1"
	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", userID), nil)
	mapCtx := map[string]string{pathUserID: userID}
	request = mux.SetURLVars(request, mapCtx)

	user := User{
		Name: "name",
	}

	gomock.InOrder(
		t.service.EXPECT().GetUser(gomock.Any(), userID).Return(user, nil),
	)

	t.handler.Get(recorder, request)

	expectBody, _ := json.Marshal(user)

	t.Equal(http.StatusOK, recorder.Code)
	t.Equal(expectBody, recorder.Body.Bytes())
}

func (t *HandlerTest) TestHandler_Get_WhenServiceNotFound_ThenReturnHttp404() {
	recorder := httptest.NewRecorder()
	userID := "1"
	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", userID), nil)
	mapCtx := map[string]string{pathUserID: userID}
	request = mux.SetURLVars(request, mapCtx)

	gomock.InOrder(
		t.service.EXPECT().GetUser(gomock.Any(), userID).Return(User{}, errors.New("not found user")),
	)

	t.handler.Get(recorder, request)

	expectBody, _ := json.Marshal(map[string]string{"error": "not found user"})

	t.Equal(http.StatusNotFound, recorder.Code)
	t.Equal(expectBody, recorder.Body.Bytes())
}

func (t *HandlerTest) TestHandler_Get_WhenPathIDNotFound_ThenReturnHttp400() {
	recorder := httptest.NewRecorder()
	userID := ""
	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", userID), nil)
	mapCtx := map[string]string{pathUserID: userID}
	request = mux.SetURLVars(request, mapCtx)

	gomock.InOrder(
		t.service.EXPECT().GetUser(gomock.Any(), userID).Times(0),
	)

	t.handler.Get(recorder, request)

	expectBody, _ := json.Marshal(map[string]string{"error": "bad request"})

	t.Equal(http.StatusBadRequest, recorder.Code)
	t.Equal(expectBody, recorder.Body.Bytes())
}

func (t *HandlerTest) TestHandler_Get_WhenPanic_ThenReturnHttp500() {
	recorder := httptest.NewRecorder()
	userID := "999"
	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", userID), nil)
	mapCtx := map[string]string{pathUserID: userID}
	request = mux.SetURLVars(request, mapCtx)

	gomock.InOrder(
		t.service.EXPECT().GetUser(gomock.Any(), userID).Times(0),
	)

	t.handler.Get(recorder, request)

	expectBody, _ := json.Marshal(map[string]string{"error": "internal server error"})

	t.Equal(http.StatusInternalServerError, recorder.Code)
	t.Equal(expectBody, recorder.Body.Bytes())
}

func (t *HandlerTest) TestHandler_NewHandler_WhenCall_ThenReturnHandler() {
	handler := &handler{}
	t.IsType(NewHandler(t.service), handler)
}
