package controllers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Bouchiba43/Auth-Go/controllers"
	"github.com/Bouchiba43/Auth-Go/initializers"
	"github.com/Bouchiba43/Auth-Go/models"
	"github.com/Bouchiba43/Auth-Go/repositories"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewUserController(t *testing.T) {
	// Create a new user controller
	userController := controllers.NewUserController(repositories.NewUserRepository(initializers.DB))

	// Check if the userController is not nil
	if userController == nil {
		t.Error("NewUserController() should not return nil")
	}
}

func TestSignup(t *testing.T) {
	// Create a new user controller
	userController := controllers.NewUserController(repositories.NewUserRepository(initializers.DB))

	// Create a new gin engine
	_, r := gin.CreateTestContext(httptest.NewRecorder())

	// Register the Signup route
	r.POST("/signup", userController.Signup)

	// Test cases
	tests := []struct {
		name           string
		requestBody    interface{}
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Invalid request body",
			requestBody:    "invalid JSON",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Failed to read body"}`,
		},
		{
			name: "Valid request body",
			requestBody: models.User{
				Email:    "test@example.com",
				Password: "password",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{}`, // Adjust according to your actual response
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request body
			var body io.Reader
			if tt.requestBody != nil {
				jsonBody, _ := json.Marshal(tt.requestBody)
				body = bytes.NewBuffer(jsonBody)
			} else {
				body = nil
			}

			// Create a new HTTP request
			req, err := http.NewRequest(http.MethodPost, "/signup", body)
			assert.NoError(t, err)

			// Set content type
			req.Header.Set("Content-Type", "application/json")

			// Create a response recorder
			w := httptest.NewRecorder()

			// Serve the HTTP request
			r.ServeHTTP(w, req)

			// Check the status code
			assert.Equal(t, tt.expectedStatus, w.Code)

			// Check the response body
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}
