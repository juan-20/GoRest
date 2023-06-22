package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	supa "github.com/nedpals/supabase-go"
	"github.com/spf13/viper"
)

type User struct {
	Name       string    `json:"name" validate:"required"`
	Id         string    `json:"id"`
	Created_at time.Time `json:"created_at"`
	Phone      string    `json:"phone"`
}

func main() {
	// Load the configuration file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Failed to read config file:", err)
	}

	supabaseUrl := viper.GetString("API_KEY")
	supabaseKey := viper.GetString("API_URL")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	if supabase != nil {
		fmt.Println("Erro")
	}

	fmt.Println(supabaseUrl, supabaseKey)

	// Create a new Gin router
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		var results []map[string]interface{}
		err := supabase.DB.From("users").Select("*").Execute(&results)
		fmt.Println(results)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"users": results})
	})

	r.POST("/user", func(c *gin.Context) {
		var user User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validate := validator.New()
		if err := validate.Struct(user); err != nil {
			// Return validation errors
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		uuid := uuid.New().String()
		currentTime := time.Now()
		row := User{
			Name:       user.Name,
			Id:         uuid,
			Created_at: currentTime,
			Phone:      user.Phone,
		}
		fmt.Println(row) // Inserted rows

		var results []User
		err := supabase.DB.From("users").Insert(row).Execute(&results)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"message": results})
	})

	// Run the server
	r.Run(":8000")
}
