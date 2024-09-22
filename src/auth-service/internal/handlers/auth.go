package handlers

import (
	"net/http"

	"auth-service/internal/db"
	"auth-service/internal/models"
	"auth-service/internal/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login handler for authenticating users and generating JWT token.
func Login(c *gin.Context) {
    var loginDetails struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&loginDetails); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid input"})
        return
    }

    var user models.User
    if err := models.GetUserByUsername(db.DB, loginDetails.Username, &user); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid username or password"})
        return
    }

    // Compare hashed password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDetails.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid username or password"})
        return
    }

    // Create JWT token with the role included in the claims
    token, err := services.CreateJWT(user.ID, user.Username, user.Email, user.Role)  // Pass role to the token creation
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to create token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"access_token": token})
}

// Register handler for user registration
func Register(c *gin.Context) {
	var registerDetails struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Role	 string `json:"role"` //Possible roles: viewer, editor, admin
	}

	if err := c.ShouldBindJSON(&registerDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid input"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerDetails.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to hash password"})
		return
	}

	role := registerDetails.Role
	if role == "" {
		role = "viewer"
	}

	user := models.User{
		Username: registerDetails.Username,
		Password: string(hashedPassword),
		Email:    registerDetails.Email,
		Role:     role,
	}

	if err := models.CreateUser(db.DB, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "User registered successfully"})
}