package repositories

import (
	"database/sql"

	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/configuration"
	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/models"
	_ "github.com/go-sql-driver/mysql"
)

func GetUsersRepository() (*sql.Rows, error) {
	db, errConnectDb := configuration.ConnectDb()
	if errConnectDb != nil {
		return nil, errConnectDb
	}

	result, errSelect := db.Query("select * from users")
	if errSelect != nil {
		return nil, errSelect
	}

	return result, nil
}

func NewUserRepository(user *models.User) (sql.Result, error) {
	db, errConnectDb := configuration.ConnectDb()
	if errConnectDb != nil {
		return nil, errConnectDb
	}

	statement, errPrepare := db.Prepare("insert into users (name, username, password) values (?, ?, ?);")
	if errPrepare != nil {
		return nil, errPrepare
	}

	defer statement.Close()

	result, errExec := statement.Exec(user.Name, user.Username, user.Password)
	if errExec != nil {
		return nil, errExec
	}

	return result, nil
}
