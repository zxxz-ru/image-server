package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	FilePath string `json:"imageDir"`
}

func Prt() string {
	var c Config
	res, err := filepath.Abs("./utils/test.json")
	if err != nil {
		return fmt.Sprintf("Error during exec. err: %v\n", err)
	}
	bytes, err := ioutil.ReadFile(res)
	if err != nil {
		return fmt.Sprintf("Can not read file: %v\n", err)
	}
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return fmt.Sprintf("Can not Unmarshal bytes err: %v\n", err)
	}
	return fmt.Sprintf("%v\n", c.FilePath)

}
