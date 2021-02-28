package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/sawasaki-narumi/gin-practice/basic"
)

func main() {
	fmt.Println("vim-go")

	r := gin.Default()
	api := r.Group("/")

	basic.InitializeRoutes(api)

	r.Run()
}
