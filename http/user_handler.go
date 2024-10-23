package http

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermelinosp/crud1/application/services/validations"
	"github.com/guilhermelinosp/crud1/application/usecases/users"
	"github.com/guilhermelinosp/crud1/config/logs"
	"github.com/guilhermelinosp/crud1/domain/dtos/requests"
	"net/http"
)

func InitUserHandler(r *gin.RouterGroup) {
	r.GET("/getAllUsers", func(c *gin.Context) {
		logs.Info("Init UserHandler.FindUsers")

		// Code for finding user
	})

	r.GET("/getUserById/:id", func(c *gin.Context) {
		logs.Info("Init UserHandler.FindUserById")

		// Code for finding user
	})

	r.GET("/getUserByEmail/:email", func(c *gin.Context) {
		logs.Info("Init UserHandler.FindUserByEmail")

		// Code for finding user
	})

	r.POST("/create", func(c *gin.Context) {
		logs.Info("Init UserHandler.CreateUser")

		var request requests.UserRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			logs.Error("Error trying to validate user info", err)
			validation := validations.ValidatorRequest(err)
			c.JSON(validation.Code, gin.H{"error": validation.Message})
			return
		}

		if err := users.CreateTask(&request); err != nil {
			logs.Error("Error creating user", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
		})

		logs.Info("User created successfully")
	})

	r.PUT("/updateUserById/:id", func(c *gin.Context) {
		logs.Info("Init UserHandler.UpdateUser")

		// Code for finding user
	})

	r.DELETE("/deleteUserById/:id", func(c *gin.Context) {
		logs.Info("Init UserController.DeleteUser")

		// Code for finding user
	})
}
