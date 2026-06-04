package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ponderada-cleancode/domain"
	"ponderada-cleancode/handler"
	"ponderada-cleancode/repository"
	"ponderada-cleancode/service"
)

func main() {
	db := connectDatabase()

	figurinhaRepo := repository.NewFigurinhaRepository(db)
	figurinhaSvc := service.NewFigurinhaService(figurinhaRepo)
	figurinhaHandler := handler.NewFigurinhaHandler(figurinhaSvc)

	r := setupRouter(figurinhaHandler)

	log.Println("Servidor rodando em http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func connectDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("figurinhas.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco:", err)
	}
	if err := db.AutoMigrate(&domain.Figurinha{}); err != nil {
		log.Fatal("Falha na migração:", err)
	}
	return db
}

func setupRouter(figurinhaHandler *handler.FigurinhaHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API de figurinhas funcionando",
			"routes": []string{
				"POST /figurinha",
				"GET /figurinha",
				"GET /figurinha/:id",
				"PUT /figurinha/:id",
				"DELETE /figurinha/:id",
			},
		})
	})
	figurinhas := r.Group("/figurinha")
	{
		figurinhas.POST("", figurinhaHandler.Create)
		figurinhas.GET("", figurinhaHandler.List)
		figurinhas.GET("/:id", figurinhaHandler.GetByID)
		figurinhas.PUT("/:id", figurinhaHandler.Update)
		figurinhas.DELETE("/:id", figurinhaHandler.Delete)
	}
	return r
}
