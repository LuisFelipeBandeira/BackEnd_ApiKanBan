package repositories

import (
	"database/sql"
	"errors"
	"time"

	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/configuration"
	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/models"
	_ "github.com/go-sql-driver/mysql"
)

func GetUsersRepository() ([]models.User, error) {
	db, errConnectDb := configuration.ConnectDb()
	if errConnectDb != nil {
		return nil, errConnectDb
	}

	defer db.Close()

	result, errSelect := db.Query("select * from users")
	if errSelect != nil {
		return nil, errSelect
	}

	defer result.Close()

	var users []models.User

	for result.Next() {
		var (
			user         models.User
			createdAtRaw []byte
			errParseTime error
		)
		if errScan := result.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Password, &user.AdmPermission, &createdAtRaw); errScan != nil {
			return []models.User{}, errScan
		}

		user.CreatedAt, errParseTime = time.Parse("2006-01-02 15:04:05", string(createdAtRaw))
		if errParseTime != nil {
			return []models.User{}, errParseTime
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserByIDRepository(id int) (models.User, error) {
	db, errConnectDb := configuration.ConnectDb()
	if errConnectDb != nil {
		return models.User{}, errConnectDb
	}

	defer db.Close()

	var count int

	db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", id).Scan(&count)
	if count < 1 {
		return models.User{}, errors.New("nenhum usuario encontrado com o ID informado")
	}

	var (
		user         models.User
		createdAtRaw []byte
		errParseTime error
	)

	if errScan := db.QueryRow("select * from users where id = ?", id).Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Password, &user.AdmPermission, &createdAtRaw); errScan != nil {
		return models.User{}, errScan
	}

	user.CreatedAt, errParseTime = time.Parse("2006-01-02 15:04:05", string(createdAtRaw))
	if errParseTime != nil {
		return models.User{}, errParseTime
	}

	return user, nil
}

func NewUserRepository(user models.User) (sql.Result, error) {
	db, errConnectDb := configuration.ConnectDb()
	if errConnectDb != nil {
		return nil, errConnectDb
	}

	defer db.Close()

	var count int

	db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ? OR email = ?", user.Username, user.Email).Scan(&count)
	if count >= 1 {
		return nil, errors.New("username e(ou) email ja cadastrado em nosso banco de dados")
	}

	statement, errPrepare := db.Prepare("insert into users (name, username, email, password, isAdm, created_at) values (?, ?, ?, ?, ?, ?);")
	if errPrepare != nil {
		return nil, errPrepare
	}

	defer statement.Close()

	result, errExec := statement.Exec(user.Name, user.Username, user.Email, user.Password, user.AdmPermission, user.CreatedAt)
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

	db.QueryRow("SELECT Count(*) FROM users WHERE Id = ?", id).Scan(&countResult)
	if countResult < 1 {
		return errors.New("usuario nao encontrato")
	}

	if user.Name != "" {
		statement, errPrepare := db.Prepare("update users set name = ? where id = ?")
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
		statement, errPrepare := db.Prepare("update users set password = ? where id = ?")
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
		statement, errPrepare := db.Prepare("update users set username = ? where id = ?")
		if errPrepare != nil {
			return errPrepare
		}

		defer statement.Close()

		_, errExec := statement.Exec(user.Username, id)
		if errExec != nil {
			return errExec
		}
	}

	if user.Email != "" {
		statement, errPrepare := db.Prepare("update users set email = ? where id = ?")
		if errPrepare != nil {
			return errPrepare
		}

		defer statement.Close()

		_, errExec := statement.Exec(user.Email, id)
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

	db.QueryRow("select Count(*) from users where email = ?", userLogin.Email).Scan(&resultCount)

	if resultCount < 1 {
		return 0, errors.New("nenhum usuario encontrado com o email informado")
	}

	var (
		passwordDb string
		userId     int
	)

	db.QueryRow("select id, password from users where email = ?", userLogin.Email).Scan(&userId, &passwordDb)

	userLogin.EncriptPassword()

	if userLogin.Password != passwordDb {
		return 0, errors.New("incorrect password")
	}

	return userId, nil
}

func UserIsAdm(userId int) (bool, error) {
	db, errToConnectDb := configuration.ConnectDb()
	if errToConnectDb != nil {
		return false, errToConnectDb
	}

	defer db.Close()

	var fieldUserIsAdm int

	if errScan := db.QueryRow("select isAdm from users where id = ?", userId).Scan(&fieldUserIsAdm); errScan != nil {
		return false, errScan
	}

	if fieldUserIsAdm == 1 {
		return true, nil
	}

	return false, nil
}
