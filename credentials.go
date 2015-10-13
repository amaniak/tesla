package tesla

import (
	"bufio"
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/bgentry/speakeasy"
	"log"
	"os"
)

type Credentials struct {
	Username     string
	Password     string
	ClientId     string
	ClientSecret string
}

func GetCredentials(path string) *Credentials {
	cred := new(Credentials)
	cred.FetchFromEnv(path)
	cred.FetchUserName()
	cred.FetchPassWord()
	return cred
}

func (cred *Credentials) FetchFromEnv(path string) {

	// Load config
	cfg, err := goconfig.LoadConfigFile(path)

	// Parse
	clientId, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "id")
	clientSecret, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "secret")

	// Check for errors
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Check for empty values
	if len(clientId) == 0 || len(clientSecret) == 0 {
		log.Fatal("[tesla]: No config found .. ")
		os.Exit(1)
	}

	// Set on Credentials
	cred.ClientId = clientId
	cred.ClientSecret = clientSecret

	// trace
	Ok("[OK] Oauth credentials found")

}

func (cred *Credentials) FetchPassWord() {
	// Grab password
	password, err := speakeasy.Ask("Enter password: ")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	cred.Password = password
}

func (cred *Credentials) FetchUserName() {
	//Grab username
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	cred.Username = username
}
