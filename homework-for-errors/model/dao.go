package model

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/shunshizhen/golang-learning-camp-homework/homework-for-errors/config"
)

// Query query the database
func Query(field string) error {
	err := mockErr(field)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("%d, not found", config.NotFound)
		}
		return fmt.Errorf("%d, not found", config.SystemErr)
	}
	return nil
}

// IsNoRow check errorcode == 40001
func IsNoRow(err error) bool {
	return strings.HasPrefix(err.Error(), fmt.Sprintf("%d", config.NotFound))
}

func mockErr(data string)error{
	return errors.New("sql")
}
