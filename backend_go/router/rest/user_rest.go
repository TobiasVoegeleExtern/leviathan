package rest

import (
	"backend_go/db/dao"
	"backend_go/db/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, userDAO *dao.UserDAO) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", createUser(userDAO))
		userRoutes.GET("/", getUsers(userDAO))
		userRoutes.PUT("/:id", updateUser(userDAO))
		userRoutes.DELETE("/:id", deleteUser(userDAO))
		userRoutes.POST("/login", login(userDAO))
		userRoutes.POST("/authenticate", authenticateUser(userDAO))
		userRoutes.GET("/:identifier", getUser(userDAO))

	}
}

func createUser(userDAO *dao.UserDAO) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the email already exists
		existingUser, err := userDAO.GetByEmail(input.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check if email exists"})
			return
		}
		if existingUser != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already taken"})
			return
		}

		// Call DAO method with password included
		user, err := userDAO.Create(input.Name, input.Email, input.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
		c.JSON(http.StatusCreated, user)
	}
}

func getUsers(userDAO *dao.UserDAO) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := userDAO.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func updateUser(userDAO *dao.UserDAO) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Convert string ID to int
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		// Update user details (including password if provided)
		err = userDAO.Update(id, input.Name, input.Email, input.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	}
}

func deleteUser(userDAO *dao.UserDAO) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		err = userDAO.Delete(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}

func login(userDAO *dao.UserDAO) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Check if the user exists by email
		user, err := userDAO.GetByEmail(input.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
			return
		}

		// If user does not exist
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Check if the passwords match (no hashing)
		if user.Password != input.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// If everything matches, send success response (no JWT)
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	}
}

// New authentication endpoint that checks both email and username
func authenticateUser(userDAO *dao.UserDAO) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Identifier string `json:"identifier"` // Can be email or username
			Password   string `json:"password"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Authenticate user by email or username
		user, err := userDAO.AuthenticateUser(input.Identifier, input.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authenticate user"})
			return
		}

		// If user not found
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Return user ID and name upon successful authentication
		c.JSON(http.StatusOK, gin.H{
			"message": "Authentication successful",
			"user_id": user.ID,
			"name":    user.Name,
		})
	}
}

func getUser(userDAO *dao.UserDAO) gin.HandlerFunc {
	return func(c *gin.Context) {
		identifier := c.Param("identifier")

		user, err := userDAO.GetByIDOrEmail(identifier)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
