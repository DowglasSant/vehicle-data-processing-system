package model

type City struct {
	ID      int
	Name    string
	StateID int // Adicionado para referenciar o Estado
}
