package googleAPI

import (
	"encoding/json"
	"log"
)

type CredentialsJSON struct {
	Installed Installed `json:"installed"`
}
type Installed struct {
	ClientID                string   `json:"client_id"`
	ProjectID               string   `json:"project_id"`
	AuthURI                 string   `json:"auth_uri"`
	TokenURI                string   `json:"token_uri"`
	AuthProviderX509CertURL string   `json:"auth_provider_x509_cert_url"`
	ClientSecret            string   `json:"client_secret"`
	RedirectURIs            []string `json:"redirect_uris"`
}

func getCredentials() []byte {
	/*
		b, err := ioutil.ReadFile("../../config/credentials.json")
		if err != nil {
			return nil, fmt.Errorf("Unable to read client secret file: %v", err)
		}
	*/
	var jsonKey = CredentialsJSON{
		Installed: Installed{
			ClientID:                "525434559900-la0n63r5b1vrediimigrfrr4u88cgaml.apps.googleusercontent.com",
			ProjectID:               "quickstart-1548473599903",
			AuthURI:                 "https://accounts.google.com/o/oauth2/auth",
			TokenURI:                "https://oauth2.googleapis.com/token",
			AuthProviderX509CertURL: "https://www.googleapis.com/oauth2/v1/certs",
			ClientSecret:            "naAMiEsLMvvZeIh_n3SHXxBi",
			RedirectURIs:            []string{"urn:ietf:wg:oauth:2.0:oob", "http://localhost"},
		},
	}
	b, err := json.Marshal(jsonKey)
	if err != nil {
		log.Println("error..")
	}
	return b
}