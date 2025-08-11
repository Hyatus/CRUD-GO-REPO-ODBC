package controllers

import (
	"errors" // Importing errors for error creation
	"net/http" // Importing the net/http package for HTTP response handling
	"strconv"  // Importing strconv for converting string to int
	// "fmt"      // Importing fmt for formatting strings
	"github.com/Hyatus/myapi/models"   // Importing the models package for user data structure
	"github.com/Hyatus/myapi/services" // Importing the services package for user-related operations
	"github.com/Hyatus/myapi/utils"    // Importing the utils package for response handling
	"github.com/gin-gonic/gin"         // Importing the gin framework for handling HTTP requests
)

func GetUsers(c *gin.Context){
	users,err := services.GetAllUsers() // Fetching all users from the service
	// How to handle error in the response
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch users") // Using the utility function to respond with an error
		// c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()}) // Responding with an error message
		return 
	}
	utils.RespondWithSuccess(c, users) // Using the utility function to respond with success
}

func GetUserByID(c *gin.Context){
	id := c.Param("id") // Extracting the user ID from the URL parameter
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User ID is required"})
		return
	}

	// convert id to uint and fetch user
	userID, err := ParseUintParam(c, "id") // Using the utility function to parse the user ID from the context
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error()) // Responding with a bad request error if parsing fails
		return
	}

	user, err := services.GetUserByID(userID) // Fetching the user by ID using the service
	if err != nil {
		if err.Error() == "User not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch user") // Using the utility function to respond with an error
		return
	}
	utils.RespondWithSuccess(c, user) // Using the utility function to respond with success
}

func CreateUser(c *gin.Context){
	var newUser models.User // Creating a new User instance 
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.CreateUser(newUser) // Creating a new user using the service
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create user")
		// c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()}) // Responding with an error message
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message":"Usuario creado con éxito"}) // Responding with the created user and status code 201
}

func UpdateUser(c *gin.Context){
	id := c.Param("id")
	var updatedUser models.User // Creating a new User instance for the updated user data
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Responding with a bad request error if binding fails
		return
	}
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User ID is required"})
		return
	}
	userID, err := ParseUintParam(c, "id") // Using the utility function to parse the user ID from the context
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error()) // Responding with a bad request error if parsing fails
		return
	}
	resultado, errorActualiza := services.UpdateUser(userID, updatedUser) // Updating the user using the service
	if errorActualiza != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Error al actualizar el usuario"}) // Responding with a not found error if the user does not exist
		return
	}
	
	utils.RespondWithSuccess(c, gin.H{"message": resultado}) // Using the utility function to respond with success
}

func DeleteUser(c *gin.Context){
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User ID is required"}) // Responding with a bad request error if the ID is missing
		return
	}
	userID, err := ParseUintParam(c, "id") // Using the utility function to parse the user ID from the context
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error()) // Responding with a bad request error if parsing fails
		return
	}
	err = services.DeleteUser(userID) // Deleting the user using the service
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to delete user")
		// c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()}) // Responding with an error message
		return
	}
	c.JSON(http.StatusNoContent, nil) // Responding with no content status after successful deletion
}

// HandleWrongPath handles requests to undefined routes and logs the wrong path
func HandleWrongPath(c *gin.Context, path string) {
	utils.Log.Warn("Wrong path accessed: ", path) // Logging the wrong path access
	c.JSON(404, gin.H{"message": "The path " + path + " does not exist"}) // Responding with a 404 error message
}

func ParseUintParam(c *gin.Context, paramName string) ( int , error) {
	paramValue := c.Param(paramName) // Extracting the parameter value from the context
	if paramValue == "" {
		return -1, errors.New("Parameter " + paramName + " is required") // Returning an error if the parameter is empty
	}

	intValue, err := strconv.Atoi(paramValue) // Converting the parameter value to an integer
	if err != nil {
		return -1, errors.New("Invalid " + paramName + " parameter") // Returning an error if conversion fails
	}
	return intValue, nil // Returning the integer value and nil error if successful
}