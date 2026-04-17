package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User table structure
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

var DB *gorm.DB

func main() {
	// 1. Connect to MySQL (Change 'root:' if you set a password in XAMPP)
	dsn := "root:root@tcp(127.0.0.1:3306)/simple_login_project?charset=utf8mb4&parseTime=True&loc=Local" //if use docker mysql
	//dsn := "root:@tcp(127.0.0.1:3306)/simple_login_project?charset=utf8mb4&parseTime=True&loc=Local" //if use xampp mysql
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// 2. Auto-migrate (This creates the 'users' table automatically)
	DB.AutoMigrate(&User{})

	r := gin.Default()

	// 3. Enable CORS (Allows Vue.js to talk to this API)
	r.Use(cors.Default())

	// 4. API Routes
	r.POST("/register", RegisterUser)
	r.POST("/login", LoginUser)

	r.Run(":8080") // API runs on http://localhost:8080
}

func RegisterUser(c *gin.Context) {
	var input User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Requirement: Password Hashing
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	input.Password = string(hashedPassword)

	if err := DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already taken"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful!"})
}

func LoginUser(c *gin.Context) {
	var input User
	var user User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Find user by email
	if err := DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Compare hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful!", "user": user.Email})
}
