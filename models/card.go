package models

import (
	"errors"
	"strings"
	"time"
)

// Adicionar motivo, CNPJ, id cliente, nome fantasia
type Card struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Desc          string    `json:"desc"`
	BoardId       uint      `json:"board_id"`
	CreatedBy     uint      `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
	TicketOwnerId uint      `json:"ticket_owner_id"`
	FinishedBy    uint      `json:"finished_by"`
	Finished      uint      `json:"finished"`
	FinishedAt    time.Time `json:"finished_at"`
}

type UpdateCard struct {
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	BoardId       int    `json:"board_id"`
	TicketOwnerId int    `json:"ticket_owner_id"`
}

func (card *Card) ValidAndFormat() error {
	switch {
	case card.Title == "":
		return errors.New("o titulo do card nao pode ser vazio")
	case card.Desc == "":
		return errors.New("a descricao do card nao pode ser vazia")
	case card.BoardId == 0:
		return errors.New("o board do card nao pode ser vazio")
	default:
		card.Desc = strings.ToLower(strings.TrimSpace(card.Desc))
		card.Title = strings.ToLower(strings.TrimSpace(card.Title))
		return nil
	}
}
