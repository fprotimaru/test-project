package cmd

import (
	"bytes"
	"github.com/fprotimaru/test-project/models"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var SendSMS = cli.Command{
	Name:   "sendsms",
	Usage:  "send sms to user",
	Action: sendSms,
}

func sendSms(c *cli.Context) error {
	for _, user := range models.Users {
		sendSMS(&user)
	}
	return nil
}

func sendSMS(user *models.User) {
	endpoint := "https://rest.nexmo.com/sms/json"

	var values = url.Values{}
	values.Add("from", "fprotimaru")
	values.Add("text", "hello")
	values.Add("to", user.Phone)
	values.Add("api_key", "b5256cfd")
	values.Add("api_secret", "rGp7hBHLzVUqUSuB")

	res, err := http.Post(endpoint, "application/x-www-form-urlencoded", bytes.NewReader([]byte(values.Encode())))
	if err != nil {
		log.Println(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
}
