package clase

import (
	"errors"
	"fmt"
)

type customError struct {
	status  int
	message string
}

func (err customError) Error() string {
	return fmt.Sprintf("Status: %d\nMessage: %v", err.status, err.message)
}

func Errores() {
	err1 := fmt.Errorf("UPS")
	err2 := errors.New(err1.Error() + "!")
	fmt.Println(err2.Error())
}
