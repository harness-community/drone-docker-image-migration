package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

func main() {
	// Get environment variables
	sourceRegistry := os.Getenv("PLUGIN_SOURCE_REGISTRY")
	sourceUsername := os.Getenv("PLUGIN_SOURCE_USERNAME")
	sourcePassword := os.Getenv("PLUGIN_SOURCE_PASSWORD")
	sourceNamespace := os.Getenv("PLUGIN_SOURCE_NAMESPACE")

	destinationRegistry := os.Getenv("PLUGIN_DESTINATION_REGISTRY")
	destinationUsername := os.Getenv("PLUGIN_DESTINATION_USERNAME")
	destinationPassword := os.Getenv("PLUGIN_DESTINATION_PASSWORD")
	destinationNamespace := os.Getenv("PLUGIN_DESTINATION_NAMESPACE")

	imageName := os.Getenv("PLUGIN_IMAGE_NAME")
	imageTag := os.Getenv("PLUGIN_IMAGE_TAG")

	accessKeyId := os.Getenv("PLUGIN_AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("PLUGIN_AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("PLUGIN_AWS_REGION")

	// Authenticate and transfer image
	authenticateAndTransferImage(sourceRegistry, sourceUsername, sourcePassword, sourceNamespace, destinationRegistry, destinationUsername, destinationPassword, destinationNamespace, imageName, imageTag, accessKeyId, secretAccessKey, region)
}

func authenticateAndTransferImage(sourceRegistry, sourceUsername, sourcePassword, sourceNamespace, destinationRegistry, destinationUsername, destinationPassword, destinationNamespace, imageName, imageTag, accessKeyId, secretAccessKey, region string) {

	if sourceUsername == "AWS" && sourcePassword == "" {
		// Get token from AWS
		var err error
		sourcePassword, err = getToken(accessKeyId, secretAccessKey, region)
		if err != nil {
			fmt.Println("Error getting token:", err)
			os.Exit(1)
		}
	}

	if sourceUsername != "" && sourcePassword != "" {
		execCommand("skopeo", "login", "--username", sourceUsername, "--password", sourcePassword, sourceRegistry)
	}

	if destinationUsername == "AWS" && destinationPassword == "" {
		// Get token from AWS
		var err error
		destinationPassword, err = getToken(accessKeyId, secretAccessKey, region)
		if err != nil {
			fmt.Println("Error getting token:", err)
			os.Exit(1)
		}
	}

	if destinationUsername != "" && destinationPassword != "" {
		execCommand("skopeo", "login", "--username", destinationUsername, "--password", destinationPassword, destinationRegistry)
	}

	sourceImage := fmt.Sprintf("docker://%s/%s/%s:%s", sourceRegistry, sourceNamespace, imageName, imageTag)
	destinationImage := fmt.Sprintf("docker://%s/%s", destinationRegistry, destinationNamespace)

	execCommand("skopeo", "copy", sourceImage, destinationImage)
}

func execCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing %s: %s\n", name, err)
		fmt.Println(string(output))
		os.Exit(1)
	}

	fmt.Println(string(output))
}

func getToken(accessKeyId, secretAccessKey, region string) (string, error) {
	// Initialize a session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
			AccessKeyID:     accessKeyId,
			SecretAccessKey: secretAccessKey,
		}),
	})

	if err != nil {
		fmt.Println("Error creating session:", err)
		return "", err
	}

	// Create an ECR client
	svc := ecr.New(sess)

	// Get the authorization token
	input := &ecr.GetAuthorizationTokenInput{}
	result, err := svc.GetAuthorizationToken(input)
	if err != nil {
		fmt.Println("Error getting authorization token:", err)
		return "", err
	}

	var awsToken string

	for _, data := range result.AuthorizationData {
		token, err := base64.StdEncoding.DecodeString(*data.AuthorizationToken)
		if err != nil {
			fmt.Println("Error decoding token:", err)
			return "", err
		}

		awsToken = string(token)
	}

	return strings.Split(awsToken, ":")[1], nil
}
