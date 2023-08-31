package repositories

import (
	"errors"
	"time"

	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/configuration"
	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/models"
	_ "github.com/go-sql-driver/mysql"
)

func GetCardsRepository() ([]models.Card, error) {
	db, errConnect := configuration.ConnectDb()
	if errConnect != nil {
		return nil, errConnect
	}

	defer db.Close()

	rows, errQuery := db.Query("select * from cards")
	if errQuery != nil {
		return nil, errQuery
	}

	defer rows.Close()

	var cards []models.Card

	for rows.Next() {

		var (
			card         *models.Card
			createdAtRaw []byte
			errParseTime error
		)

		if errScan := rows.Scan(&card.ID, &card.Title, &card.Desc, &card.BoardId, &card.ColumnId, &card.CreatedBy, &createdAtRaw,
			&card.TicketOwnerId, &card.FinishedBy, &card.Finished, &card.FinishedAt); errScan != nil {
			return nil, errScan
		}

		card.CreatedAt, errParseTime = time.Parse("2006-01-02 15:04:05", string(createdAtRaw))
		if errParseTime != nil {
			return []models.Card{}, errParseTime
		}

		cards = append(cards, *card)
	}

	return cards, nil
}

func GetCardByIdRepository(id int) (models.Card, error) {
	db, errConnect := configuration.ConnectDb()
	if errConnect != nil {
		return models.Card{}, errConnect
	}

	defer db.Close()

	var count int

	db.QueryRow("Select Count(*) FROM Cards WHERE Id = ?", id).Scan(&count)

	if count < 1 {
		return models.Card{}, errors.New("card not found")
	}

	var (
		card                models.Card
		createdAtRaw        []byte
		errToParseCreatedAt error
	)

	if errScan := db.QueryRow("select * from cards where id = ?", id).Scan(&card.ID, &card.Title, &card.Desc, &card.BoardId, &card.ColumnId, &card.CreatedBy, &createdAtRaw,
		&card.TicketOwnerId, &card.FinishedBy, &card.Finished, &card.FinishedAt); errScan != nil {
		return models.Card{}, errScan
	}

	card.CreatedAt, errToParseCreatedAt = time.Parse("2006-01-02 15:04:05", string(createdAtRaw))
	if errToParseCreatedAt != nil {
		return models.Card{}, errToParseCreatedAt
	}

	return card, nil
}

func DeleteCardRepository(id int) error {
	db, errConnect := configuration.ConnectDb()
	if errConnect != nil {
		return errConnect
	}

	defer db.Close()

	var count int

	db.QueryRow("Select Count(*) FROM Cards WHERE Id = ?", id).Scan(&count)

	if count < 1 {
		return errors.New("user not found")
	}

	statement, errPrepare := db.Prepare("Delete From cards Where id = ?")
	if errPrepare != nil {
		return errPrepare
	}

	defer statement.Close()

	_, errExec := statement.Exec(id)
	if errExec != nil {
		return errExec
	}

	return nil
}

func NewCardRepository(card models.Card) error {
	db, err := configuration.ConnectDb()
	if err != nil {
		return err
	}

	defer db.Close()

	statement, errPrepare := db.Prepare("INSERT INTO Cards(title, description, board_id, column_id, created_by, created_at) VALUES (?, ?, ?, ?, ?, ?)")
	if errPrepare != nil {
		return errPrepare
	}

	defer statement.Close()

	_, errExec := statement.Exec(card.Title, card.Desc, card.BoardId, card.ColumnId, card.CreatedBy, card.CreatedAt)
	if errExec != nil {
		return errExec
	}

	return nil
}

func FinishCardRepository(cardId int, user models.User, finished_at time.Time) error {
	db, errConnectDatabase := configuration.ConnectDb()
	if errConnectDatabase != nil {
		return errConnectDatabase
	}

	defer db.Close()

	statement, errPrepare := db.Prepare("Update cards SET finished_by = ?, is_finished = ?, finished_at = ? WHERE Id = ?")
	if errPrepare != nil {
		return errPrepare
	}

	defer statement.Close()

	_, errExec := statement.Exec(user.ID, 1, finished_at, cardId)
	if errExec != nil {
		return errExec
	}

	return nil
}

func UpdateCardRepository(cardId int, cardFieldsToUpdate models.UpdateCard, user models.User) error {
	db, errConnectDb := configuration.ConnectDb()
	if errConnectDb != nil {
		return errConnectDb
	}

	defer db.Close()

	// Title         string `json:"title"`
	// Desc          string `json:"desc"`
	// BoardId       uint   `json:"board_id"`
	// ColumnId      uint   `json:"column_id"`
	// TicketOwnerId uint   `json:"ticket_owner_id"

	if cardFieldsToUpdate.Title != "" {
		statement, errPrepare := db.Prepare("update cards set title = ? where id = ?")
		if errPrepare != nil {
			return errPrepare
		}

		defer statement.Close()

		_, errExec := statement.Exec(cardFieldsToUpdate.Title, cardId)
		if errExec != nil {
			return errExec
		}
	}

	if cardFieldsToUpdate.Desc != "" {
		statement, errPrepare := db.Prepare("update cards set description = ? where id = ?")
		if errPrepare != nil {
			return errPrepare
		}

		defer statement.Close()

		_, errExec := statement.Exec(cardFieldsToUpdate.Desc, cardId)
		if errExec != nil {
			return errExec
		}
	}

	if cardFieldsToUpdate.BoardId != 0 {
		statement, errPrepare := db.Prepare("update cards set board_id = ? where id = ?")
		if errPrepare != nil {
			return errPrepare
		}

		defer statement.Close()

		_, errExec := statement.Exec(cardFieldsToUpdate.BoardId, cardId)
		if errExec != nil {
			return errExec
		}
	}

	if cardFieldsToUpdate.ColumnId != 0 {
		statement, errPrepare := db.Prepare("update cards set column_id = ? where id = ?")
		if errPrepare != nil {
			return errPrepare
		}

		defer statement.Close()

		_, errExec := statement.Exec(cardFieldsToUpdate.ColumnId, cardId)
		if errExec != nil {
			return errExec
		}
	}

	if cardFieldsToUpdate.TicketOwnerId != 0 {
		statement, errPrepare := db.Prepare("update cards set ticket_owner_id = ? where id = ?")
		if errPrepare != nil {
			return errPrepare
		}

		defer statement.Close()

		_, errExec := statement.Exec(cardFieldsToUpdate.TicketOwnerId, cardId)
		if errExec != nil {
			return errExec
		}
	}

	return nil
}
