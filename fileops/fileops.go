package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"github.com/Pallinder/go-randomdata"
)

func ReadBalanceFromFile(fileName string) (float64, error) {
	randomdata.Address()
	val, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("file doesnt exist")
	}
	balanceText := string(val)
	balanceNumb, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, errors.New("file doesnt have a balance")
	}
	return balanceNumb, nil
}
func WriteBalanceToFile(balance float64, balanceFile string) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}
