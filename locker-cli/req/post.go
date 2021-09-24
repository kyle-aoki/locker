package req

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"lkcli/session"
	"log"
	"net/http"
)

var Host string = "http://localhost:8080"

func Post(uri string, class interface{}, withCredentials bool) string {
	url := Host + uri
	postBody, _ := json.Marshal(class)

	client := http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postBody))
	req.Header.Add("Content-Type", "application/json");

	if err != nil {
		log.Fatal("Failed to make post request.")
	}

	if withCredentials {
		username, sessionToken := getCredentials()
		req.Header = http.Header{
			"username":     []string{username},
			"sessiontoken": []string{sessionToken},
			"Content-Type": []string{"application/json"},
		}
	}
	
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Failed to read post request.")
	}
	defer res.Body.Close()
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error parsing response body.")
	}

	return string(body)
}

func getCredentials() (string, string) {
	usernamePath, sessionTokenPath := session.GetCredentialPaths()

	bytes, err := ioutil.ReadFile(usernamePath)
	if err != nil {
		log.Fatal("Failed to read username file at " + usernamePath)
	}
	username := string(bytes)

	bytes, err = ioutil.ReadFile(sessionTokenPath)
	if err != nil {
		log.Fatal("Failed to read session token file at " + sessionTokenPath)
	}
	sessionToken := string(bytes)

	return username, sessionToken
}
