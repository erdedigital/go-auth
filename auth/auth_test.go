package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var (
	uri = "http://localhost:8000"
)

func TestRegister(t *testing.T) {
	url := fmt.Sprintf("%s/auth/register", uri)
	// fmt.Printf(url)
	method := "POST"

	payload := strings.NewReader(`{
		"email": "dede.yuyandy@gmail.com"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
