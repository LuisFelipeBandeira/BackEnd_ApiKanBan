package repositories

import (
	"database/sql"
	"errors"

	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/configuration"
	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/models"
	_ "github.com/go-sql-driver/mysql"
)

func GetUsersRepository() (*sql.Rows, error) {
	db, errConnectDb := configuration.ConnectDb()
	if errConnectDb != nil {
		return nil, errConnectDb
	}

	defer db.Close()

	result, errSelect := db.Query("select * from users")
	if errSelect != nil {
		return nil, errSelect
	}

	return result, nil
}

func GetUserByIDRepository(id int) (*sql.Row, error) {
	db, errConnectDb := configuration.ConnectDb()
	if errConnectDb != nil {
		return nil, errConnectDb
	}

	defer db.Close()

	var count int

	db.QueryRow("SELECT COUNT(*) FROM Users WHERE id = ?", id).Scan(&count)
	if count < 1 {
		return nil, errors.New("nenhum usuario encontrado com o ID informado")
	}

	sqlRow := db.QueryRow("select * from users where id = ?", id)

	return sqlRow, nil
}

func NewUserRepository(user *models.User) (sql.Result, error) {
	db, errConnectDb := configuration.ConnectDb()
	if errConnectDb != nil {
		return nil, errConnectDb
	}

	defer db.Close()

	var count int

	db.QueryRow("SELECT COUNT(*) FROM Users WHERE username = ?", user.Username).Scan(&count)
	if count >= 1 {
		return nil, errors.New("username ja cadastrado em nosso banco de dados")
	}

	statement, errPrepare := db.Prepare("insert into users (name, username, email, password, isAdm) values (?, ?, ?, ?, ?);")
	if errPrepare != nil {
		return nil, errPrepare
	}

	defer statement.Close()

	result, errExec := statement.Exec(user.Name, user.Username, user.Email, user.Password, user.AdmPermission)
	if errExec != nil {
		return nil, errExec
	}

	return result, nil
}

func DeleteUserRepository(id int) (sql.Result, error) {
	db, errConnectDb := configuration.ConnectDb()
	if errConnectDb != nil {
		return nil, errConnectDb
	}

	defer db.Close()

	var count int

	db.QueryRow("SELECT COUNT(*) FROM Users WHERE Id = ?", id).Scan(&count)
	if count < 1 {
		return nil, errors.New("nenhum usuario encontrado com o ID informado")
	}

	statement, errPrepare := db.Prepare("delete from users where Id = ?;")
	if errPrepare != nil {
		return nil, errPrepare
	}

	defer statement.Close()

	result, errExec := statement.Exec(id)
	if errExec != nil {
		return nil, errExec
	}

	return result, nil
}

func UpdateUserRepository(id int, user models.UpdateUser) error {
	db, errConnectDb := configuration.ConnectDb()
	if errConnectDb != nil {
		return errConnectDb
	}

	defer db.Close()

	var countResult int

	db.QueryRow("SELECT Count(*) FROM Users WHERE Id = ?", id).Scan(&countResult)
	if countResult < 1 {
		return errors.New("usuario nao encontrato")
	}

	if user.Name != "" {
		statement, errPrepare := db.Prepare("update Users set name = ? where id = ?")
		if errPrepare != nil {
			return errPrepare
		}

		defer statement.Close()

		_, errExec := statement.Exec(user.Name, id)
		if errExec != nil {
			return errExec
		}
	}

	if user.Password != "" {
		statement, errPrepare := db.Prepare("update Users set password = ? where id = ?")
		if errPrepare != nil {
			return errPrepare
		}

		defer statement.Close()

		user.EncriptPassword()

		_, errExec := statement.Exec(user.Password, id)
		if errExec != nil {
			return errExec
		}
	}

	if user.Username != "" {
		statement, errPrepare := db.Prepare("update Users set username = ? where id = ?")
		if errPrepare != nil {
			return errPrepare
		}

		defer statement.Close()

		_, errExec := statement.Exec(user.Username, id)
		if errExec != nil {
			return errExec
		}
	}

	return nil
}

func LoginRepository(userLogin models.LoginUser) (int, error) {
	db, errConnect := configuration.ConnectDb()
	if errConnect != nil {
		return 0, errConnect
	}

	var resultCount int

	db.QueryRow("select Count(*) from Users where username = ?", userLogin.Username).Scan(&resultCount)

	if resultCount < 1 {
		return 0, errors.New("nenhum usuario encontrado com o username informado")
	}

	var passwordDb string
	var userId int

	db.QueryRow("select id, password from Users where username = ?", userLogin.Username).Scan(&userId, &passwordDb)

	userLogin.EncriptPassword()

	if userLogin.Password != passwordDb {
		return 0, errors.New("incorrect password")
	}

	return userId, nil
}
