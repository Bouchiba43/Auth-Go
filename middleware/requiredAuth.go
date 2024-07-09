package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Bouchiba43/Auth-Go/initializers"
	"github.com/Bouchiba43/Auth-Go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)


func RequireAuth(c *gin.Context){

	fmt.Println("In the middleware")
	
	//Get the cookie off request
	tokenString,err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//decode/validate it

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		
		//check the expiration date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		
		}
	
		//find the user in the token sub

		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//attach the user to request 

		c.Set("user", user)
	
		//continue
		c.Next()
		
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}


}