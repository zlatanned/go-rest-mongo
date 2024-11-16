package main

import (
	"fmt"						// For logging
	"github.com/gin-gonic/gin"  // Import Gin package for building web APIs
    "net/http"                  // Import for HTTP-related constants and functions
	"encoding/json"    			// To decode the JSON response from the API
	"github.com/joho/godotenv"	// To load env vars from .env
    "os"						// built-in os module
)

func main() {

	// Load .env file
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
        return
    }

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(response *gin.Context) {
		response.JSON(http.StatusOK, gin.H{
			"message": "Welcome To the homepage",
		})
	});

	r.GET("/title/:title_id", func(response *gin.Context) {
		movieID := response.Param("title_id");

		apiURL := os.Getenv("IMDB_HOST_URL") + "/?i=" + movieID + "&apikey=" + os.Getenv("IMDB_APIKEY");

		// Make a GET request to the movie API
		resp, err := http.Get(apiURL)

		if err != nil {
			response.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movie details"})
			return
		}
		defer resp.Body.Close();

		// Check for non-200 status codes
		if resp.StatusCode == http.StatusUnauthorized { // 401 Unauthorized for invalid API key
			response.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			return
		}

		if resp.StatusCode == http.StatusNotFound { // 404 Not Found for non-existent movie
			response.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
			return
		}

		if resp.StatusCode == http.StatusBadRequest { // 400 Bad Request for incorrect requests
			response.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, check the movie ID or API key"})
			return
		}

		// Decode the JSON response from the API
		var movieData map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&movieData); err != nil {
			response.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse movie details"})
			return
		}

		// Send the movie data as the JSON response
		response.JSON(http.StatusOK, movieData)
	});

	r.GET("/ping", func(response *gin.Context) {
		response.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	});

	// r.SetTrustedProxies(nil)
	fmt.Println("About to run ...")
	r.Run(":3061") // listen and serve on 0.0.0.0:3061
}