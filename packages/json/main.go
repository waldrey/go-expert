package main

import (
	"encoding/json"
	"os"
)

type AccountBank struct {
	Number  int `json:"-"`
	Balance int `json:"balance" validate:"gt-0"`
}

func main() {
	account := AccountBank{Number: 1, Balance: 100}
	res, err := json.Marshal(account)
	if err != nil {
		println(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"n": 2, "b": 200}`)
	var accountX AccountBank
	err = json.Unmarshal(jsonPuro, &accountX)
	if err != nil {
		println(err)
	}
	println(accountX.Balance)
}
