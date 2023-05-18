package repositories

import (
	"database/sql"
	"errors"

	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/configuration"
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
		return nil, errors.New("user not found")
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
