package req

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"lkcli/filesystem"
	"log"
	"net/http"
)

func Post(uri string, class interface{}, withCredentials bool) string {
	host := filesystem.GetHost()
	url := host + uri
	
	postBody, _ := json.Marshal(class)

	client := http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postBody))
	req.Header.Add("Content-Type", "application/json");

	if err != nil {
		log.Fatal("Failed to make post request.")
	}

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
		log.Fatal("Failed to read post request.")
	}
	defer res.Body.Close()
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error parsing response body.")
	}

	return string(body)
}
