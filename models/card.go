package models

import (
	"errors"
	"strings"
)

type Card struct {
	ID         int    `json:"id"`
	Board      string `json:"board"`
	Desc       string `json:"desc"`
	CreatedBy  string `json:"createdby"`
	CreatedAt  string `json:"createdat"`
	FinishedBy string `json:"finishedby"`
	Finished   int    `json:"finished"`
	FinishedAt string `json:"finishedat"`
}

func (card *Card) ValidAndFormat() error {
	switch {
	case card.Board == "":
		return errors.New("o board do card nao pode ser vazio")
	case card.Desc == "":
		return errors.New("a descricao do card nao pode ser vazia")
	case card.CreatedBy == "":
		return errors.New("o criador do card nao pode ser vazio")
	default:
		card.Board = strings.TrimSpace(card.Board)
		card.Board = strings.ToLower(card.Board)
		card.Desc = strings.TrimSpace(card.Desc)
		card.Desc = strings.ToLower(card.Desc)
		card.CreatedBy = strings.TrimSpace(card.CreatedBy)
		card.CreatedBy = strings.ToLower(card.CreatedBy)
		return nil
	}
}
