package utils

import (
	"encoding/json"
	"github.com/fredrikssongustav/avanza-cli/avanza/authentication"
	"io/ioutil"
	"os"
)

func GetAuth() authentication.AuthenticationResponse {
	jsonFile, _ := os.Open(".session/.data")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var authAsJSON authentication.AuthenticationResponse

	json.Unmarshal(byteValue, &authAsJSON)

	return authAsJSON
}
