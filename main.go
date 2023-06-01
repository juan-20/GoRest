package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
	"github.com/spf13/viper"
)

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

	// Define your routes
	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.GET("/api/multiply/:value", func(c *gin.Context) {
		// Retrieve the value from the URL parameter
		valueParam := c.Param("value")
		fmt.Println(valueParam)
		// Convert the value to an integer
		value, err := strconv.Atoi(valueParam)
		// Validade header
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid value",
			})
			return
		}

		// Multiply the value by 10
		result := value * 10

		// Return the result as a JSON response
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})

	// Run the server
	r.Run(":8000")
}
