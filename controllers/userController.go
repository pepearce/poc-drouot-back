package controllers

import (
	"drouotBack/models"
	"drouotBack/services"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	userService     services.UserService           = services.NewUserService()
	useInputService services.CheckUserInputService = services.NewCheckUserInputService()
)

// GET /users
// Get all users
func GetAllUsers(c *gin.Context) {
	// services.FindUsers(c)
	users := userService.FindUsers()
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GET /users/:id
// Get a user by his id
func GetUserById(c *gin.Context) {
	// services.FindUser(c)
	user, err := userService.FindUser(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// POST /users
// Add a user
func AddUser(c *gin.Context) {
	// services.CreateUser(c)

	input, err := useInputService.CheckCreateInput(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userService.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"ID": user.ID, "firstName": user.FirstName, "lastName": user.LastName, "email": user.Email, "address": user.Address, "role": user.Role})
}

// PUT /users/:id
// Update user
func UpdateUser(c *gin.Context) {
	// services.UpdateUser(c)
	input, err := useInputService.CheckUpdateInput(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := userService.UpdateUser(c.Param("id"), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ID": user.ID, "firstName": user.FirstName, "lastName": user.LastName, "email": user.Email, "address": user.Address, "role": user.Role})
}

// DELETE /users/:id
// Delete user
func DeleteUser(c *gin.Context) {
	// services.DeleteUser(c)
	deleted, err := userService.DeleteUser(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": deleted})
}

// POST /login
func SignInUser(c *gin.Context) {
	session := sessions.Default(c)

	// services.FindUserByEmail(c)
	userInput, err := useInputService.CheckSignInInput(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := userService.SignInUser(userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session.Set("sessionId", user.UUID)
	session.Save()
	c.JSON(http.StatusOK, gin.H{"ID": user.ID, "firstName": user.FirstName, "lastName": user.LastName, "email": user.Email, "address": user.Address, "role": user.Role})
}

// /GET /signout
func SignOutUser(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"data": "User Sign out successfully"})
}

// /GET /user/auctions/:id
func FindUserAuctions(c *gin.Context) {
	var auctions []models.Auction
	models.DB.Where("userID = ?", c.Param("id")).Find(&auctions)
	c.JSON(http.StatusOK, gin.H{"data": auctions})
}
