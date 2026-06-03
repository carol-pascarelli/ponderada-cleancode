package main

import (
	"ponderada3/domain"
	"ponderada3/handler"
	"ponderada3/repository"
	"ponderada3/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("gastos.db"), &gorm.Config{})
	if err != nil {
		panic("falha ao conectar ao banco: " + err.Error())
	}
	db.AutoMigrate(&domain.Expense{})

	repo := repository.NewExpenseRepository(db)
	svc := service.NewExpenseService(repo)
	h := handler.NewExpenseHandler(svc)

	r := gin.Default()
	h.RegisterRoutes(r)
	r.Run(":8080")
}
