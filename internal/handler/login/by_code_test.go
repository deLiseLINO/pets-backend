package login_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pets-backend/internal/handler"
	"pets-backend/internal/handler/login"
	mock_login "pets-backend/internal/handler/login/mocks"
	"pets-backend/internal/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandleByCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	email := "some-email@gmail.com"
	code := "3451"

	userID := uuid.New()
	username := "some-unique-name123"
	name := "Sam"
	token := "some-token"

	otpSvc := mock_login.NewMockByCodeOtpService(ctrl)
	otpSvc.EXPECT().
		VerifyCode(gomock.Any(), email, code).
		Return(nil)

	userSvc := mock_login.NewMockByCodeUserService(ctrl)
	userSvc.EXPECT().
		GetByEmail(gomock.Any(), email).
		Return(&models.User{
			ID:         userID,
			Email:      email,
			UniqueName: username,
			Name:       name,
		}, nil)

	ssoSvc := mock_login.NewMockByCodeSSOService(ctrl)
	ssoSvc.EXPECT().
		GenerateToken(userID.String()).
		Return(token, nil)

	data := login.ByCodeRequest{
		Email:   email,
		OtpCode: code,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(&data)
	assert.NoError(t, err)

	router := gin.Default()
	path := "/login/by_code"
	router.POST(path, login.HandleByCode(otpSvc, userSvc, ssoSvc))

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, path, &buf)

	router.ServeHTTP(w, req)

	response := handler.APIResponse[login.ByCodeResponse]{}

	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, username, response.Data.UserName)
	assert.Equal(t, name, response.Data.Name)
	assert.Equal(t, token, response.Data.Token)
}
