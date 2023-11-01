package console

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"

	"github.com/sijoma/camundactl/internal/config"
)

func (c *Console) Auth() {
	c.deviceCodeAuth()
}

func (c *Console) MachineLogin(clientID, clientSecret string) error {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(c.auth0.ClientCredentials(clientID, clientSecret))
	if err != nil {
		return fmt.Errorf("unable to encode auth0 config %w", err)
	}
	post, err := http.Post(c.auth0.TokenURL(), "application/json", body)
	if err != nil {
		return fmt.Errorf("unable to post to auth0 %w", err)
	}

	var token oauth2.Token
	err = json.NewDecoder(post.Body).Decode(&token)
	if err != nil {
		return err
	}

	if token.AccessToken != "" {
		fmt.Println(color.CyanString("Authentication successful "))
		fmt.Println(color.CyanString("You are great 💪! You are authenticated and can now start to use the CLI."))
		config.StoreAccessToken(c.stage, &token)
	}
	return nil
}

func (c *Console) deviceCodeAuth() {
	reqString := fmt.Sprintf("client_id=%s&scope=%s&audience=%s", c.auth0.Oauth().ClientID, c.auth0.Scopes(), c.auth0.Audience())
	payload := strings.NewReader(reqString)

	req, err := http.NewRequest("POST", c.auth0.AuthURL(), payload)
	if err != nil {
		fmt.Println("Error getting device code auth: ", err)
		return
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	client := http.Client{Timeout: time.Second * 10}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("There was an error performing the request", err)
		return
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response of device code request: ", err)
		return
	}

	var resp deviceCodeFlow
	err = json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Println("There was an error performing the device flow request:", err)
		return
	}

	fmt.Println(color.CyanString("You will now be taken to your browser for authentication"))
	fmt.Println(color.CyanString("User Code: " + resp.UserCode))
	fmt.Println(color.CyanString("Verification URL: " + resp.VerificationURIComplete))

	_ = open.Run(resp.VerificationURIComplete)

	for {
		token, err := c.requestToken(resp.DeviceCode)
		if err != nil {
			fmt.Println("polling token ... retrying. ", err)
		}
		if token.AccessToken != "" {
			fmt.Println(color.CyanString("Authentication successful "))
			fmt.Println(color.CyanString("You are great 💪! You are authenticated and can now start to use the CLI."))
			config.StoreAccessToken(c.stage, token)
			return
		}
		time.Sleep(time.Duration(resp.Interval) * time.Second)
	}
}

func (c *Console) requestToken(deviceCode string) (*oauth2.Token, error) {
	fmt.Println("Polling token ...")
	url := c.auth0.TokenURL()
	reqString := fmt.Sprintf("grant_type=%s&device_code=%s&client_id=%s", c.auth0.GrantType(), deviceCode, c.auth0.Oauth().ClientID)
	payload := strings.NewReader(reqString)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var token oauth2.Token
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}
	return &token, err
}

type deviceCodeFlow struct {
	DeviceCode              string `json:"device_code"`
	UserCode                string `json:"user_code"`
	VerificationURI         string `json:"verification_uri"`
	ExpiresIn               int    `json:"expires_in"`
	Interval                int    `json:"interval"`
	VerificationURIComplete string `json:"verification_uri_complete"`
}
