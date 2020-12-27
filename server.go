package main

import (
	"github.com/jck4/jck.sh/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
)

func main() {
	r := gin.Default()

	var store = sessions.NewCookieStore([]byte("secret"))

	r.Use(sessions.Sessions("goquestsession", store))

	addRoutes(r)

	// Connect to database
	models.ConnectDataBase()
	r.Run()
}

