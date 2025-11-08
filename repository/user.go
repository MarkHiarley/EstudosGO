package repository

import (
	"encoding/json"
	"estudosGo/model"
	"fmt"
	"os"
)

func GetUsers() ([]model.User, error) {
	file, err := os.ReadFile("users.json")
	if err != nil {
		return nil, err
	}

	var users []model.User
	if err := json.Unmarshal(file, &users); err != nil {
		return nil, err
	}
	fmt.Println(users)
	return users, nil
}
