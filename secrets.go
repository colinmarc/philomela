package main

import (
	"log"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type secretsData struct {
	SessionKey string
	TwitterAuth twitterSecretsData
	TwitterAnnounce twitterSecretsData
}

type twitterSecretsData struct {
	Key string
	Secret string
	AccessToken string
	AccessSecret string
}

func loadSecrets() secretsData {
	raw, err := ioutil.ReadFile("secrets.yaml")
	if err != nil {
		log.Fatal(err)
	}

	secrets := secretsData{}
	err = yaml.Unmarshal(raw, &secrets)
	if err != nil {
		log.Fatal(err)
	}

	return secrets
}
