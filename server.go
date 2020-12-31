package main

import (
	"github.com/jck4/jck.sh/models"
	"github.com/gin-gonic/gin"

)

func main() {
	router := gin.Default()
	models.ConnectDataBase()
	addRoutes(router)
	router.Run("172.22.241.220:8080")
}

// func main() {
// 	r := gin.Default()

// 	addRoutes(r)
		
// 	scopes := []string{
// 		"repo",
// 		// You have to select your own scope from here -> https://developer.github.com/v3/oauth/#scopes
// 	}
// 	secret := []byte("secret")  
// 	sessionName := "jcksession"

// 	r.Use(github.Session(sessionName))
// 	github.Setup("http://172.22.241.220:8080/auth/", "creds.json", scopes, secret)
// 	r.GET("/login", github.LoginHandler)
	
	
// 	// Connect to database
// 	models.ConnectDataBase()
// 	r.Run("172.22.241.220:8080")
// }
