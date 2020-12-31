package main

import (
	"github.com/gin-gonic/gin"
	// controllers "github.com/jck4/jck.sh/controllers"
	"github.com/gin-gonic/contrib/static"
	// "github.com/zalando/gin-oauth2/github"
	// "github.com/gin-gonic/contrib/sessions"
	// "crypto/rand"
	// "encoding/base64"
	// "github.com/golang/glog"
)

func addRoutes(router *gin.Engine) {

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

}

