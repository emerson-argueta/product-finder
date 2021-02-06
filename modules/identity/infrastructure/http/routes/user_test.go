package routes

import (
	"emerson-argueta/m/v2/modules/identity/usecase/controllers"
	"emerson-argueta/m/v2/shared/infrastructure"
	"emerson-argueta/m/v2/shared/infrastructure/http/authorization"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

var rec *httptest.ResponseRecorder
var echoRouter *echo.Echo
var identityHandler *IdentityHandler

func setUpHandler(e *echo.Echo) {

	logger := log.New(os.Stderr, "", log.LstdFlags)
	authorizationService := authorization.NewJWTService(infrastructure.GlobalConfig)
	controllers := controllers.NewMock(authorizationService, logger)

	identityHandler = &IdentityHandler{
		Echo:        e,
		Controllers: controllers,
		Logger:      &log.Logger{},
	}
}
func setUpRequest(e *echo.Echo, reqBody string) echo.Context {
	rec = httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := e.NewContext(req, rec)
	return c
}
func TestLogin(t *testing.T) {
	echoRouter = echo.New()
	setUpHandler(echoRouter)

	t.Run("loginSuccessTest", loginSuccessTest)
	t.Run("loginFailedPasswordTest", loginFailedPasswordTest)
	t.Run("loginNonExistingUserTest", loginNonExistingUserTest)
	t.Run("loginInvalidRequestTest", loginInvalidRequestTest)

	echoRouter.Close()
	rec = nil

}

func loginSuccessTest(t *testing.T) {
	succesfulLoginJSON := `{"email":"test@test.com","password":"password"}`
	c := setUpRequest(echoRouter, succesfulLoginJSON)
	if err := identityHandler.handleLogin(c); err != nil {
		t.Fatal("Login failed: ", err)
	}
	code := rec.Code
	if code != http.StatusOK {
		t.Log("Login was not successful with request body: ", succesfulLoginJSON)
		t.Fail()
	}
}
func loginFailedPasswordTest(t *testing.T) {
	failedLoginJSON := `{"email":"test@test.com","password":"password1"}`
	c := setUpRequest(echoRouter, failedLoginJSON)
	if err := identityHandler.handleLogin(c); err != nil {
		t.Fatal("Login failed: ", err)
	}
	code := rec.Code
	if code != http.StatusUnauthorized {
		t.Log("Login should fail due to incorrect password but did with request body: ", failedLoginJSON)
		t.Fail()
	}
}
func loginNonExistingUserTest(t *testing.T) {
	failedLoginJSON := `{"email":"thisuserdoesnotexist@test.com","password":"password"}`
	c := setUpRequest(echoRouter, failedLoginJSON)
	if err := identityHandler.handleLogin(c); err != nil {
		t.Fatal("Login failed: ", err)

	}
	code := rec.Code
	if code != http.StatusNotFound {
		t.Log("Login should fail due to non existing user but did not with request body: ", failedLoginJSON)
		t.Fail()
	}
}
func loginInvalidRequestTest(t *testing.T) {
	failedLoginJSON := `{"password":"password"}`
	c := setUpRequest(echoRouter, failedLoginJSON)
	if err := identityHandler.handleLogin(c); err != nil {
		t.Fatal("Login failed: ", err)
	}
	code := rec.Code
	if code != http.StatusBadRequest {
		t.Log("Login should fail due to a bad request body did not with request body: ", failedLoginJSON)
		t.Fail()
	}

}
