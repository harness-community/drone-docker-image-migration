// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/sirupsen/logrus"
)

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`

	Username    string `envconfig:"PLUGIN_USERNAME"`
	Password    string `envconfig:"PLUGIN_PASSWORD"`
	Source      string `envconfig:"PLUGIN_SOURCE"`
	Destination string `envconfig:"PLUGIN_DESTINATION"`

	// Optional
	SourceUsername    string `envconfig:"PLUGIN_SOURCE_USERNAME"`
	SourcePassword    string `envconfig:"PLUGIN_SOURCE_PASSWORD"`
	Overwrite         bool   `envconfig:"PLUGIN_OVERWRITE"`
	AwsAccessKeyID    string `envconfig:"PLUGIN_AWS_ACCESS_KEY_ID"`
	AwsSecretAcessKey string `envconfig:"PLUGIN_AWS_SECRET_ACCESS_KEY"`
	AwsRegion         string `envconfig:"PLUGIN_AWS_REGION"`
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {
	err := validateArgs(&args)
	if err != nil {
		return err
	}

	if args.Username == "AWS" && args.Password == "" {
		args.Password, err = getAWSPassword(args.AwsAccessKeyID, args.AwsSecretAcessKey, args.AwsRegion)
		if err != nil {
			return err
		}
	}

	if args.SourceUsername == "AWS" && args.SourcePassword == "" {
		args.SourcePassword, err = getAWSPassword(args.AwsAccessKeyID, args.AwsSecretAcessKey, args.AwsRegion)
		if err != nil {
			return err
		}
	}

	err = migrateImage(&args)
	if err != nil {
		return err
	}

	logrus.Infof("successfully migrated image %s to %s", args.Source, args.Destination)

	return nil
}

func validateArgs(args *Args) error {
	if args.Username == "" {
		return fmt.Errorf("missing username")
	}
	if args.Password == "" && args.Username != "AWS" {
		return fmt.Errorf("missing password")
	}
	if args.Source == "" {
		return fmt.Errorf("missing source")
	}
	if args.Destination == "" {
		return fmt.Errorf("missing destination")
	}
	return nil
}

func migrateImage(args *Args) error {
	var sourceAuth, destAuth authn.Authenticator

	if args.SourceUsername != "" && args.SourcePassword != "" {
		sourceAuth = authn.FromConfig(authn.AuthConfig{
			Username: args.SourceUsername,
			Password: args.SourcePassword,
		})
	} else {
		sourceAuth = authn.FromConfig(authn.AuthConfig{
			Username: args.Username,
			Password: args.Password,
		})
	}

	destAuth = authn.FromConfig(authn.AuthConfig{
		Username: args.Username,
		Password: args.Password,
	})

	if !args.Overwrite {
		_, err := crane.Head(args.Destination, crane.WithAuth(destAuth))
		if err == nil {
			return fmt.Errorf("image already exists in destination registry and overwrite is not enabled")
		}
	}

	originImage, err := crane.Pull(args.Source, crane.WithAuth(sourceAuth), crane.Insecure)
	if err != nil {
		return fmt.Errorf("failed to pull image %s: %w", args.Source, err)
	}

	err = crane.Push(originImage, args.Destination, crane.WithAuth(destAuth), crane.WithNoClobber(!args.Overwrite))
	if err != nil {
		return fmt.Errorf("failed to push image %s: %w", args.Destination, err)
	}

	return nil
}
