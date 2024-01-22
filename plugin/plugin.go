// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"encoding/base64"
	"fmt"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`

	SourceRegistry  string `envconfig:"PLUGIN_SOURCE_REGISTRY"`
	SourceUsername  string `envconfig:"PLUGIN_SOURCE_USERNAME"`
	SourcePassword  string `envconfig:"PLUGIN_SOURCE_PASSWORD"`
	SourceNamespace string `envconfig:"PLUGIN_SOURCE_NAMESPACE"`

	DestinationRegistry  string `envconfig:"PLUGIN_DESTINATION_REGISTRY"`
	DestinationUsername  string `envconfig:"PLUGIN_DESTINATION_USERNAME"`
	DestinationPassword  string `envconfig:"PLUGIN_DESTINATION_PASSWORD"`
	DestinationNamespace string `envconfig:"PLUGIN_DESTINATION_NAMESPACE"`

	ImageName string `envconfig:"PLUGIN_IMAGE_NAME"`
	ImageTag  string `envconfig:"PLUGIN_IMAGE_TAG"`

	AccessKeyID     string `envconfig:"PLUGIN_AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `envconfig:"PLUGIN_AWS_SECRET_ACCESS_KEY"`
	Region          string `envconfig:"PLUGIN_AWS_REGION"`
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {

	if err := ValidateArgs(args); err != nil {
		return err
	}

	if err := LoginToRegistry(
		args.SourceUsername,
		args.SourcePassword,
		args.SourceRegistry,
	); err != nil {
		return err
	}

	if args.DestinationUsername == "AWS" && args.DestinationPassword == "" {
		if args.AccessKeyID == "" || args.SecretAccessKey == "" || args.Region == "" {
			return fmt.Errorf("AWS credentials are not set")
		}

		var err error

		args.DestinationPassword, err = GetAWSPassword(
			args.AccessKeyID,
			args.SecretAccessKey,
			args.Region,
		)

		if err != nil {
			return fmt.Errorf("error getting AWS credentials: %s", err)
		}
	}

	if err := LoginToRegistry(
		args.DestinationUsername,
		args.DestinationPassword,
		args.DestinationRegistry,
	); err != nil {
		return fmt.Errorf("error logging in to destination registry: %s", err)
	}

	sourceImage := fmt.Sprintf("docker://%s/%s/%s:%s", args.SourceRegistry, args.SourceNamespace, args.ImageName, args.ImageTag)
	destinationImage := fmt.Sprintf("docker://%s/%s", args.DestinationRegistry, args.DestinationNamespace)

	if err := CopyImage(sourceImage, destinationImage); err != nil {
		return fmt.Errorf("error copying image: %s", err)
	}

	fmt.Println("Image copied successfully")

	return nil
}

// ValidateArgs validates the plugin arguments.
func ValidateArgs(args Args) error {
	if args.SourceRegistry == "" {
		return fmt.Errorf("source registry is not set")
	}
	if args.SourceUsername == "" {
		return fmt.Errorf("source username is not set")
	}
	if args.SourcePassword == "" {
		return fmt.Errorf("source password is not set")
	}
	if args.SourceNamespace == "" {
		return fmt.Errorf("source namespace is not set")
	}
	if args.DestinationRegistry == "" {
		return fmt.Errorf("destination registry is not set")
	}
	if args.DestinationUsername == "" {
		return fmt.Errorf("destination username is not set")
	}
	if args.DestinationNamespace == "" {
		return fmt.Errorf("destination namespace is not set")
	}
	if args.ImageName == "" {
		return fmt.Errorf("image name is not set")
	}
	if args.ImageTag == "" {
		return fmt.Errorf("image tag is not set")
	}
	return nil
}

func LoginToRegistry(username, password, regsitry string) error {
	cmd := exec.Command("skopeo", "login", "--username", username, "--password", password, regsitry)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func GetAWSPassword(accessKeyID, secretAccessKey, region string) (string, error) {
	// Initialize a session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
			AccessKeyID:     accessKeyID,
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

func CopyImage(source, destination string) error {
	cmd := exec.Command("skopeo", "copy", source, destination)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
