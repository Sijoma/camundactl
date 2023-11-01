package auth0

import "golang.org/x/oauth2"

const deviceCodeGrantType = "urn:ietf:params:oauth:grant-type:device_code"
const clientCredentialsGrantType = "client_credentials"

var auth0Info = map[string]Auth0App{
	"dev": {
		clientId:    "sj38M3i2kotkPV9tjtFmwIhgn7ZbPJlB",
		audience:    "cloud.dev.ultrawombat.com",
		endpointURL: "console.cloud.dev.ultrawombat.com",
		loginURL:    "weblogin.cloud.dev.ultrawombat.com",
	},
	"int": {
		clientId:    "FziYKy02BxMN0wdhequc77KB6tqOAXrB",
		audience:    "cloud.ultrawombat.com",
		endpointURL: "console.cloud.ultrawombat.com",
		loginURL:    "weblogin.cloud.ultrawombat.com",
	},
	"prod": {
		clientId:    "da5b7lU3Kgv6PeM57KmZRPZjFagoEldW",
		audience:    "cloud.camunda.io",
		endpointURL: "console.cloud.camunda.io",
		loginURL:    "weblogin.cloud.camunda.io",
	},
}

func NewAuth0App(stage string) Auth0App {
	app := auth0Info[stage]
	app.grantType = deviceCodeGrantType
	return app
}

func NewAuth0m2m(stage string, clientID string) Auth0App {
	app := auth0Info[stage]
	app.clientId = clientID
	app.grantType = clientCredentialsGrantType
	return app
}

type Auth0App struct {
	clientId    string
	audience    string
	endpointURL string
	loginURL    string
	grantType   string
}

func (a Auth0App) ClientCredentials(clientID, clientSecret string) ClientCredentialsRequest {
	return ClientCredentialsRequest{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     a.endpointURL,
		Audience:     a.audience,
		GrantType:    a.grantType,
	}

}

// ClientCredentialsRequest request body that can be used with client credentials flow
type ClientCredentialsRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Endpoint     string
	Audience     string `json:"audience"`
	GrantType    string `json:"grant_type"`
}

func (a Auth0App) AuthURL() string {
	return "https://" + a.loginURL + "/oauth/device/code"
}

func (a Auth0App) Audience() string {
	return a.audience
}

func (a Auth0App) TokenURL() string {
	return "https://" + a.loginURL + "/oauth/token"
}

func (a Auth0App) EndpointURL() string {
	return a.endpointURL
}

func (a Auth0App) AccountsURL() string { return "https://accounts." + a.audience }

func (a Auth0App) Scopes() []string {
	return []string{"openid", "profile", "email"}
}

func (a Auth0App) Oauth() *oauth2.Config {
	return &oauth2.Config{
		ClientID: a.clientId,
		Endpoint: oauth2.Endpoint{
			AuthURL:   a.AuthURL(),
			TokenURL:  a.TokenURL(),
			AuthStyle: 0,
		},
		Scopes: a.Scopes(),
	}
}

func (a Auth0App) GrantType() string {
	return a.grantType
}
