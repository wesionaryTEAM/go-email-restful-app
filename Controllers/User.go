package Controllers

import (
	"first-api/Models"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

//Email ... Structure for email
type Email struct {
	EmailTo string `json:"email_to" binding:"required"`
}

//SendEmail ... Send the SendEmail
func SendEmail(c *gin.Context) {
	setEmailConfig("First Email", c)
}

func setEmailConfig(body string, c *gin.Context) {
	envError := godotenv.Load()
	if envError != nil {
		fmt.Println("Error loading .env file")
	}
	var json Email
	c.BindJSON(&json)
	from := os.Getenv("EMAIL_FROM")
	pass := os.Getenv("PASSWORD")
	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_PORT")
	to := json.EmailTo

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there \n\n" +
		body

	err := smtp.SendMail(host + ":" + port,
		smtp.PlainAuth("", from, pass, os.Getenv("EMAIL_HOST")),
		from, []string{to}, []byte(msg))

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	fmt.Println("Email sent")
	c.JSON(http.StatusOK, "Email has been sent")
}

