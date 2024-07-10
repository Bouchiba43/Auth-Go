package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Bouchiba43/Auth-Go/models"
	"github.com/Bouchiba43/Auth-Go/repositories"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	Repository *repositories.UserRepository
}

func NewUserController(repository *repositories.UserRepository) *UserController {
	return &UserController{
		Repository: repository,
	}
}

func (controller *UserController) GetAll(c *gin.Context) {
	users, err := controller.Repository.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (controller *UserController) Signup(c *gin.Context) {
	//GET the email and password from the request

	var body struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	//Check if the email is already taken

	existingUser, err := controller.Repository.FindByEmail(body.Email)

	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find user"})
		return
	}

	//hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	//Create a new user

	user := models.User{Email: body.Email, Password: string(hash)}
	// result := initializers.DB.Create(&user)
	_, err = controller.Repository.Create(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}

	//Return the user

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (controller *UserController) Login(c *gin.Context) {

	//GET the email and password from the request

	var body struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	//look up the user by email

	existingUser, e := controller.Repository.FindByEmail(body.Email)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find user"})
		return
	}

	//compare the password with the hashed password in the database

	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(body.Password))

	//return an error if the email or password is incorrect

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	//generate a JWT token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": existingUser.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to generate token"})
		return

	}

	//send cookie with the token

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 60*60*24*30, "", "", false, true)

	//return nothing
	c.JSON(http.StatusOK, gin.H{})
}

func (controller *UserController) Logout(c *gin.Context) {

	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func (controller *UserController) Validate(c *gin.Context) {

	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"message": user})
}
