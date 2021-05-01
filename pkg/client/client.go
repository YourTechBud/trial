package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Greet greets the user
func Greet(url, name string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", url, name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var m map[string]string
	_ = json.NewDecoder(resp.Body).Decode(&m)

	return m["greeting"], nil
}
