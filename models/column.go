package models

import (
	"errors"
	"strings"
	"time"
)

type Column struct {
	ID         uint      `json:"id"`
	ColumnName string    `json:"name"`
	BoardId    uint      `json:"board_id"`
	IsActive   uint      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
}

func (c *Column) ValidAndFormat() error {
	switch {
	case c.ColumnName == "":
		return errors.New("the column's name can not be empty")
	case c.BoardId == 0:
		return errors.New("the board_id can not be empty")
	default:
		c.ColumnName = strings.ToLower(strings.TrimSpace(c.ColumnName))
		return nil
	}
}
