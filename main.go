package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ponderada-cleancode/domain"
	"ponderada-cleancode/repository"
)

func main() {
	db := connectDatabase()

	fmt.Println("Banco conectado!")

	repo := repository.NewFigurinhaRepository(db)

	figurinha := domain.Figurinha{
		Numero:  "BRA 15",
		Tipo:    "comum",
		Posicao: "atacante",
	}

	err := repo.Create(&figurinha)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Figurinha criada!")

	lista, err := repo.FindAll()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nFigurinhas cadastradas:")

	for _, f := range lista {

		fmt.Printf(
			"ID=%d | Numero=%s| Tipo=%s | Posicao=%s\n",
			f.ID,
			f.Numero,
			f.Tipo,
			f.Posicao,
		)
	}
}

func connectDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("figurinhas.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados:", err)
	}

	if err := db.AutoMigrate(&domain.Figurinha{}); err != nil {
		log.Fatal("Erro ao executar a migração:", err)
	}

	return db
}
