package models

import (
	"errors"
	"strings"
	"time"
)

type Board struct {
	ID        uint      `json:"id"`
	BoardName string    `json:"board_name"`
	IsActive  uint      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

func (b *Board) ValidAndFormat() error {
	if b.BoardName == "" {
		return errors.New("the board's name can not be empty")
	}

	b.BoardName = strings.ToLower(strings.TrimSpace(b.BoardName))
	return nil
}
