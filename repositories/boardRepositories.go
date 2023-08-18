package repositories

import (
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
		var board models.Board

		if errScan := sqlRows.Scan(&board.ID, &board.BoardName, &board.IsActive, &board.CreatedAt); errScan != nil {
			return []models.Board{}, errScan
		}

		boards = append(boards, board)
	}

	return boards, nil
}
