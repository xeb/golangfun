package oauth

import (
	"bufio"
	"bytes"
	"code.google.com/p/goauth2/oauth"
	"errors"
	"fmt"
	"log"
	"os"
)

var (
	scope       = "https://www.googleapis.com/auth/userinfo.profile"
	redirectURL = "oob"
	authURL     = "https://accounts.google.com/o/oauth2/auth"
	tokenURL    = "https://accounts.google.com/o/oauth2/token"
	requestURL  = "https://www.googleapis.com/oauth2/v1/userinfo"
	code        = ""
	cachefile   = "cache.json"

	clientId, _ = readLine("/etc/flowdle/clientid")
	secret, _   = readLine("/etc/flowdle/secret")
)

type OAuthResult struct {
	Success bool
	AuthURL string
	Debug   string
}

func readLine(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text(), scanner.Err()
}

func Try(c string) {
	r, _ := TryOAuth()
	fmt.Printf(" Result is %s ", r)
}

func TryOAuth() (result *OAuthResult, e error) {

	result = &OAuthResult{}

	// Set up a configuration.
	config := &oauth.Config{
		ClientId:     clientId,
		ClientSecret: secret,
		RedirectURL:  redirectURL,
		Scope:        scope,
		AuthURL:      authURL,
		TokenURL:     tokenURL,
		TokenCache:   oauth.CacheFile(cachefile),
	}

	// Set up a Transport using the config.
	transport := &oauth.Transport{Config: config}

	// Try to pull the token from the cache; if this fails, we need to get one.
	token, err := config.TokenCache.Token()
	if err != nil {
		if clientId == "" || secret == "" {
			// TODO: error
			result.Success = false
			e = errors.New("Cannot find clientId and/or secret.  Check /etc/flowdle")
			return
		}
		if code == "" {
			url := config.AuthCodeURL("")
			result.Success = false
			result.AuthURL = url
			return
		}

		token, err = transport.Exchange(code)
		if err != nil {
			log.Fatal("Exchange:", err)
			e = err
		}
	}

	transport.Token = token

	r, err := transport.Client().Get(requestURL)
	if err != nil {
		log.Fatal("Get:", err)
		e = err
	}
	defer r.Body.Close()

	// Send final carriage return, just to be neat.
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	s := buf.String()
	result.Debug = s
	result.Success = true
	return
}
