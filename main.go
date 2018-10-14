package main

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/juniorsev3n/td-auth-res/config"
	"github.com/juniorsev3n/td-auth-res/controllers"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.POST("/login", loginHandler)
	router.GET("/reseller/:id", auth, inDB.GetReseller)
	router.GET("/reseller", auth, inDB.GetResellers)
	router.POST("/register", auth, inDB.CreateReseller)
	router.PUT("/reseller", auth, inDB.UpdateReseller)
	router.DELETE("/reseller/:id", auth, inDB.DeleteReseller)
	router.Run(":81")
}

func loginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"token":    tokenString,
	})
}

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString != "" {
		bearerToken := strings.Split(tokenString, " ")
		if len(bearerToken) == 2 {
			token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte("secret"), nil
			})
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"status": "valid 1",
				})
				return
			}
			if token.Valid {
				c.JSON(http.StatusOK, gin.H{
					"status": "valid",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status": "Invalid",
				})
				return
			}
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "Please input authorization",
		})
		return
	}

}
