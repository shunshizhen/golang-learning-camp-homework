package biz

import (
	"github.com/pkg/errors"
	"github.com/shunshizhen/golang-learning-camp-homework/homework-for-errors/model"
)

// Find biz`s find`
func Find() error {
	err := model.Query("t")
	if err != nil {
		if model.IsNoRow(err) {
			// TODO
			return errors.Wrap(err, "find fild:")
		}
		return err
	}
	return nil
}
