package models

import (
	"errors"
	"strings"
	"time"
)

// Adicionar motivo, CNPJ, id cliente, nome fantasia
type Card struct {
	ID          int       `json:"id"`
	Board       string    `json:"board"`
	Desc        string    `json:"desc"`
	CreatedBy   string    `json:"createdby"`
	CreatedAt   time.Time `json:"createdat"`
	TicketOwner string    `json:"ticketowner"`
	FinishedBy  string    `json:"finishedby"`
	Finished    int       `json:"finished"`
	FinishedAt  time.Time `json:"finishedat"`
}

type UpdateCard struct {
	TicketOwner string `json:"ticketowner"`
	Desc        string `json:"desc"`
	Board       string `json:"board"`
}

func (card *Card) ValidAndFormat() error {
	switch {
	case card.Board == "":
		return errors.New("o board do card nao pode ser vazio")
	case card.Desc == "":
		return errors.New("a descricao do card nao pode ser vazia")
	default:
		card.Board = strings.TrimSpace(card.Board)
		card.Board = strings.ToLower(card.Board)
		card.Desc = strings.TrimSpace(card.Desc)
		card.Desc = strings.ToLower(card.Desc)
		return nil
	}
}

func (card *UpdateCard) Format() {
	card.Board = strings.TrimSpace(card.Board)
	card.Board = strings.ToLower(card.Board)
	card.Desc = strings.TrimSpace(card.Desc)
	card.Desc = strings.ToLower(card.Desc)
	card.TicketOwner = strings.TrimSpace(card.TicketOwner)
	card.TicketOwner = strings.ToLower(card.TicketOwner)
}
