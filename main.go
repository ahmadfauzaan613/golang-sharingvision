package main

import (
	"fmt"
	"golang-sharingvision/article"
	"golang-sharingvision/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// connect ke DB article
	dsn := "root:@tcp(127.0.0.1:3306)/article?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("GAGAL")
	}
	fmt.Println("sukses")

	// migration
	db.AutoMigrate(&article.Posts{})

	postRepository := article.NewRepository(db)
	postService := article.NewService(postRepository)
	handlerPosts := handler.PostHandler(postService)

	// Router
	router := gin.Default()
	api := router.Group("/v1/article")

	api.POST("/create", handlerPosts.NewPostHandler)
	api.GET("/showdata", handlerPosts.GetAll)
	api.GET("/showdata/:id", handlerPosts.GetID)
	api.PUT("/update/:id", handlerPosts.UpdatePostHandler)
	api.DELETE("/delete/:id", handlerPosts.Deleteid)

	router.Run()
}
