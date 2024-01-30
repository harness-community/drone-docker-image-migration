// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/sirupsen/logrus"
)

func getAWSPassword(accessKeyID, secretAccessKey, region string) (string, error) {

	if accessKeyID == "" || secretAccessKey == "" || region == "" {
		return "", fmt.Errorf("aws credentials not provided")
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
			AccessKeyID:     accessKeyID,
			SecretAccessKey: secretAccessKey,
		}),
	})
	if err != nil {
		return "", fmt.Errorf("failed to create aws session: %w", err)
	}

	svc := ecr.New(sess)

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

	logrus.Info("successfully retrieved aws token\n")

	return strings.Split(awsToken, ":")[1], nil
}
