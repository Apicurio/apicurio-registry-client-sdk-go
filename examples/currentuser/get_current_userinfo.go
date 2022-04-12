package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/Apicurio/apicurio-registry-client-sdk-go"
)

func main() {
	cfg := registryclient.NewConfiguration()
	cfg.Servers = registryclient.ServerConfigurations{
		{
			URL: os.Getenv("REGISTRY_API_URL"),
		},
	}
	client := registryclient.NewAPIClient(cfg)

	ctxWithCreds := context.WithValue(context.Background(), "basicAuth", getHashedCredentials())
	userInfo, _, err := client.UsersApi.GetCurrentUserInfo(ctxWithCreds).Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", userInfo)
}

func getHashedCredentials() string {
	username := os.Getenv("EXAMPLE_USERNAME")
	password := os.Getenv("EXAMPLE_PASSWORD")

	usernameHash := sha256.Sum256([]byte(username))
	passwordHash := sha256.Sum256([]byte(password))

	return fmt.Sprintf("%s:%s", usernameHash, passwordHash)
}