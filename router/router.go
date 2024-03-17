package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Initializes Routes

	InitializeRouter(router)
	
	//Run the server
	router.Run(":8080") // listen and serve on
}
