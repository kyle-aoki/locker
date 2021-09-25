package req

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"lkcli/filesystem"
	"lkcli/logger"
	"net/http"
)

func Post(uri string, class interface{}, withCredentials bool) []byte {
	host := filesystem.GetHost()
	url := host + uri

	postBody, _ := json.Marshal(class)

	client := http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postBody))
	if err != nil {
		logger.Exit("Failed to form post request.")
	}

	req.Header.Add("Content-Type", "application/json")

	if withCredentials {
		username, sessionToken := filesystem.GetCredentials()
		req.Header = http.Header{
			"username":     []string{username},
			"sessiontoken": []string{sessionToken},
			"Content-Type": []string{"application/json"},
		}
	}

	res, err := client.Do(req)
	if err != nil {
		logger.Exit("Something went wrong contacting the server. Server is down or check internet connection.")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Exit("Error parsing response body.")
	}

	return body
}
