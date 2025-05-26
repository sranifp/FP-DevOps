package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"FP-DevOps/config"
	"FP-DevOps/controller"
	"FP-DevOps/dto"
	"FP-DevOps/entity"
	"FP-DevOps/repository"
	"FP-DevOps/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func SetUpRoutes() *gin.Engine {
	r := gin.Default()
	return r
}

func SetupControllerUser() controller.UserController {
	var (
		db             = config.SetUpDatabaseConnection()
		userRepo       = repository.NewUserRepository(db)
		jwtService     = config.NewJWTService()
		userService    = service.NewUserService(userRepo)
		userController = controller.NewUserController(userService, jwtService)
	)

	return userController
}

func InsertTestUser() ([]entity.User, error) {
	db := config.SetUpDatabaseConnection()
	users := []entity.User{
		{
			ID:       uuid.New(),
			Username: "admin",
			Password: "admin123",
		},
		{
			ID:       uuid.New(),
			Username: "user",
			Password: "user123",
		},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return nil, err
		}
	}

	return users, nil
}

func CleanUpTestUsers() {
	db := config.SetUpDatabaseConnection()
	if err := db.Exec("DELETE FROM users").Error; err != nil {
		panic(err)
	}
}

func Test_Register_OK(t *testing.T) {
	r := SetUpRoutes()
	uc := SetupControllerUser()
	CleanUpTestUsers()
	r.POST("/api/user", uc.Register)

	payload := dto.UserRequest{Username: "newuser", Password: "newuser123"}
	b, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/user", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var out struct {
		Data entity.User `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &out)
	assert.Equal(t, payload.Username, out.Data.Username)
}

func Test_Register_BadRequest(t *testing.T) {
	r := SetUpRoutes()
	uc := SetupControllerUser()
	r.POST("/api/user", uc.Register)

	// empty body
	req, _ := http.NewRequest(http.MethodPost, "/api/user", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_Login_BadRequest(t *testing.T) {
	r := SetUpRoutes()
	uc := SetupControllerUser()
	r.POST("/api/user/login", uc.Login)

	req, _ := http.NewRequest(http.MethodPost, "/api/user/login", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_Login_OK(t *testing.T) {
	r := SetUpRoutes()
	userController := SetupControllerUser()
	CleanUpTestUsers()
	InsertTestUser()

	// first register
	r.POST("/api/user", userController.Register)
	r.POST("/api/user/login", userController.Login)

	payload := dto.UserRequest{
		Username: "admin",
		Password: "admin123",
	}

	// login
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	type Resp struct {
		Data struct {
			Token string `json:"token"`
			Role  string `json:"role"`
		} `json:"data"`
	}
	var resp Resp
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.NotEmpty(t, resp.Data.Token)
}

// Test Me
func Test_Me_OK(t *testing.T) {
	r := SetUpRoutes()
	userController := SetupControllerUser()
	CleanUpTestUsers()
	r.GET("/api/user/me", func(c *gin.Context) {
		// insert and set user_id
		users, _ := InsertTestUser()
		c.Set("user_id", users[0].ID.String())
		userController.Me(c)
	})

	req, _ := http.NewRequest(http.MethodGet, "/api/user/me", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	type Resp struct {
		Data entity.User `json:"data"`
	}
	var resp Resp
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, "admin", resp.Data.Username)
}
