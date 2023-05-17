package repositories

import (
	"database/sql"

	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/configuration"
	_ "github.com/go-sql-driver/mysql"
)

func GetCardsRepository() (*sql.Rows, error) {
	db, errConnect := configuration.ConnectDb()
	if errConnect != nil {
		return nil, errConnect
	}

	if errPing := db.Ping(); errPing != nil {
		return nil, errPing
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

	if errPing := db.Ping(); errPing != nil {
		return nil, errPing
	}

	defer db.Close()

	sqlRow := db.QueryRow("select * from cards where id = ?", id)

	return sqlRow, nil
}
