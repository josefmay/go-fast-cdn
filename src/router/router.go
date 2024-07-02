package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinanielsen/go-fast-cdn/src/middleware"
	"github.com/kevinanielsen/go-fast-cdn/ui"
)

// Router initializes the router and sets up middleware, routes, etc.
// It returns a *gin.Engine instance configured with the routes, middleware, etc.
func Router() *gin.Engine{
	/*
	Setups a router
	*/
	r := gin.Default()

	/*
	Configures CORS middleware
	*/
	r.Use(middleware.CORSMiddleware())

	// Adds and group the various routes for API
	AddApiRoutes(r)

	// Add the embedded ui routes
	ui.AddRoutes(r)

	/*
	Glorified wrapper of listen and serve without options
	*/
	// r.Run(":" + os.Getenv("PORT"))

	return r
}
