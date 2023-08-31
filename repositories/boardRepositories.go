package repositories

import (
	"net/http"
	"time"

	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/configuration"
	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/models"
)

func GetAllBoards() ([]models.Board, error) {
	db, errToConnectDb := configuration.ConnectDb()
	if errToConnectDb != nil {
		return []models.Board{}, errToConnectDb
	}

	defer db.Close()

	sqlRows, errToDoQuery := db.Query("select * from boards")
	if errToDoQuery != nil {
		return []models.Board{}, errToDoQuery
	}

	defer sqlRows.Close()

	var boards []models.Board

	for sqlRows.Next() {

		var (
			board               models.Board
			createdAtRaw        []byte
			errToParseCreatedAt error
		)

		if errScan := sqlRows.Scan(&board.ID, &board.BoardName, &board.IsActive, &createdAtRaw); errScan != nil {
			return []models.Board{}, errScan
		}

		board.CreatedAt, errToParseCreatedAt = time.Parse("2006-01-02 15:04:05", string(createdAtRaw))
		if errToParseCreatedAt != nil {
			return []models.Board{}, errToParseCreatedAt
		}

		boards = append(boards, board)
	}

	return boards, nil
}

func GetBoardById(boardid uint) (models.Board, int, error) {
	db, errToConnect := configuration.ConnectDb()
	if errToConnect != nil {
		return models.Board{}, http.StatusInternalServerError, errToConnect
	}

	defer db.Close()

	boardExist, errBoardExist := BoardExist(boardid)
	if errBoardExist != nil {
		return models.Board{}, http.StatusInternalServerError, errBoardExist
	}

	if !boardExist {
		return models.Board{}, http.StatusNotFound, nil
	}

	var (
		board                 models.Board
		createdAtRaw          []byte
		errToConvertCreatedAt error
	)

	if errScan := db.QueryRow("select * from board where id = ?", boardid).Scan(&board.ID, &board.BoardName, &board.IsActive, &createdAtRaw); errScan != nil {
		return models.Board{}, http.StatusInternalServerError, errScan
	}

	board.CreatedAt, errToConvertCreatedAt = time.Parse("2006-01-02 15:04:05", string(createdAtRaw))
	if errToConvertCreatedAt != nil {
		return models.Board{}, http.StatusInternalServerError, errToConvertCreatedAt
	}

	return board, http.StatusOK, nil
}

func BoardExist(boardid uint) (bool, error) {
	db, errToConnect := configuration.ConnectDb()
	if errToConnect != nil {
		return false, errToConnect
	}

	defer db.Close()

	var count int

	errScan := db.QueryRow("select count(*) from boards where id = ?", boardid).Scan(&count)
	if errScan != nil {
		return false, errScan
	}

	if count >= 1 {
		return true, nil
	}

	return false, nil
}
