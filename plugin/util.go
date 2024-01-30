// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/sirupsen/logrus"
)

func getAWSPassword(accessKeyID, secretAccessKey, region string) (string, error) {
	if accessKeyID == "" || secretAccessKey == "" || region == "" {
		return "", fmt.Errorf("aws credentials not provided")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
	)
	if err != nil {
		return "", fmt.Errorf("failed to load aws config: %w", err)
	}

	svc := ecr.NewFromConfig(cfg)

	input := &ecr.GetAuthorizationTokenInput{}
	result, err := svc.GetAuthorizationToken(context.TODO(), input)
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

	logrus.Info("successfully retrieved aws token\n")

	return strings.Split(awsToken, ":")[1], nil
}
