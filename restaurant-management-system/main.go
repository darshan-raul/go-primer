package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/darshan-raul/goprimer/restaurant/database"
)



func main(){
	port:= os.Getenv("PORT")

	if port == ""{
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	




	
	

}