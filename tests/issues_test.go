package tests

import (
	"context"
	registryclient "github.com/Apicurio/apicurio-registry-client-sdk-go"
	"os"
	"testing"
)

func TestIssue117(t *testing.T) {
	client := registryclient.NewAPIClient(nil)
	fileBody, _ := os.CreateTemp("ascj", "d")
	client.ArtifactsApi.CreateArtifact(context.TODO(), "group1").Body("string_body").Execute()
	client.ArtifactsApi.CreateArtifact(context.TODO(), "group1").Body(fileBody).Execute()
}
