package repositories

import (
	"database/sql"
	"errors"

	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/configuration"
	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/models"
	_ "github.com/go-sql-driver/mysql"
)

func GetCardsRepository() (*sql.Rows, error) {
	db, errConnect := configuration.ConnectDb()
	if errConnect != nil {
		return nil, errConnect
	}

	defer db.Close()

	rows, errQuery := db.Query("select * from cards")
	if errQuery != nil {
		return nil, errQuery
	}

	return rows, nil
}

func GetCardByIdRepository(id int) (*sql.Row, error) {
	db, errConnect := configuration.ConnectDb()
	if errConnect != nil {
		return nil, errConnect
	}

	defer db.Close()

	var count int

	db.QueryRow("Select Count(*) FROM Cards WHERE Id = ?", id).Scan(&count)

	if count < 1 {
		return nil, errors.New("card not found")
	}

	sqlRow := db.QueryRow("select * from cards where id = ?", id)

	return sqlRow, nil
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

	statement, errPrepare := db.Prepare("Delete From Users Where Id = ?")
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

func NewCardRepository(card models.Card) (models.Card, error) {
	db, err := configuration.ConnectDb()
	if err != nil {
		return models.Card{}, err
	}

	defer db.Close()

	statement, errPrepare := db.Prepare("INSERT INTO Cards(Pipeline, CreatedBy, CreatedAt) VALUES (?, ?, ?, ?)")
	if errPrepare != nil {
		return models.Card{}, errPrepare
	}

	defer statement.Close()

	_, errExec := statement.Exec(card.BoardId, card.CreatedBy, card.CreatedAt)
	if errExec != nil {
		return models.Card{}, errExec
	}

	return card, nil
}

func FinishCardRepository(cardId int, user models.User) error {
	db, errConnectDatabase := configuration.ConnectDb()
	if errConnectDatabase != nil {
		return errConnectDatabase
	}

	defer db.Close()

	statement, errPrepare := db.Prepare("Update Users SET finishedby = ?, finished = ?, finishedat = ? WHERE Id = ?")
	if errPrepare != nil {
		return errPrepare
	}

	defer statement.Close()

	_, errExec := statement.Exec(user.Username, 1, "NOW()", cardId)
	if errExec != nil {
		return errExec
	}

	return nil
}

func UpdateCardRepository(id int, cardFieldsToUpdate models.UpdateCard, user models.User) error {
	db, errConnectDb := configuration.ConnectDb()
	if errConnectDb != nil {
		return errConnectDb
	}

	defer db.Close()

	if cardFieldsToUpdate.BoardId != 0 {
		statement, errPrepare := db.Prepare("update Cards set Pipeline = ? where id = ?")
		if errPrepare != nil {
			return errPrepare
		}

		defer statement.Close()

		_, errExec := statement.Exec(cardFieldsToUpdate.BoardId, id)
		if errExec != nil {
			return errExec
		}
	}

	if cardFieldsToUpdate.Desc != "" {
		statement, errPrepare := db.Prepare("update Cards set Description = ? where id = ?")
		if errPrepare != nil {
			return errPrepare
		}

		defer statement.Close()

		_, errExec := statement.Exec(cardFieldsToUpdate.Desc, id)
		if errExec != nil {
			return errExec
		}
	}

	if cardFieldsToUpdate.TicketOwnerId != 0 {
		if user.AdmPermission != 1 {
			return errors.New("usuário não possui permisão de ADM")
		}

		statement, errPrepare := db.Prepare("update Cards set TicketOwner = ? where id = ?")
		if errPrepare != nil {
			return errPrepare
		}

		defer statement.Close()

		_, errExec := statement.Exec(cardFieldsToUpdate.TicketOwnerId, id)
		if errExec != nil {
			return errExec
		}
	}

	return nil
}

func ReopenCardRepository(user models.User, cardToReopen models.Card) error {
	db, errConnect := configuration.ConnectDb()
	if errConnect != nil {
		return errConnect
	}

	defer db.Close()

	statement, errPrepare := db.Prepare("Update Cards SET FinishedAt = '', Finished = 0, FinishedBy = '', TicketOwnerid = ? WHERE Id = ?")
	if errPrepare != nil {
		return errPrepare
	}

	defer statement.Close()

	_, errExec := statement.Exec(user.Username, cardToReopen.ID)
	if errExec != nil {
		return errExec
	}

	return nil
}
