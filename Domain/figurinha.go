package domain

import "time"

type FigurinhaType string

const (
	Comum   FigurinhaType = "comum"
	Brilhante FigurinhaType = "brilhante"
	Legends_ouro FigurinhaType = "legends_ouro"
	Legends_bronze FigurinhaType = "legends_bronze"
)


type FigurinhaPosicao string

const (
	Goleiro   FigurinhaPosicao = "goleiro"
	Zagueiro  FigurinhaPosicao = "zagueiro"
	MeioCampista FigurinhaPosicao = "meio_campista"
	Atacante  FigurinhaPosicao = "atacante"
)

type Figurinha struct {
    ID        uint      `json:"id"         gorm:"primaryKey;autoIncrement"`
    Numero    string    `json:"numero"     gorm:"not null"`
    Tipo      string    `json:"tipo"       gorm:"not null"`
    Posicao   string    `json:"posicao"    gorm:"not null"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
var TiposValidos = []string{"comum", "brilhante", "legends_ouro", "legends_bronze"} //enum para tipo e posicao
var PosicoesValidas = []string{"goleiro", "zagueiro", "meio_campista", "atacante"} //usar no service na hora de validar

type CreateFigurinhaRequest struct {
	Numero string `json:"numero"`
	Tipo string `json:"tipo"`
	Posicao string `json:"posicao"`
}

type UpdateFigurinhaRequest struct {
	Numero string `json:"numero"`
	Tipo string `json:"tipo"`
	Posicao string `json:"posicao"`
}
