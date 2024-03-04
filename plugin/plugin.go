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
	Insecure          bool   `envconfig:"PLUGIN_INSECURE"`

	SourceAWSAccessKeyID    string `envconfig:"PLUGIN_SOURCE_AWS_ACCESS_KEY_ID"`
	SourceAWSSecretAcessKey string `envconfig:"PLUGIN_SOURCE_AWS_SECRET_ACCESS_KEY"`
	SourceAWSRegion         string `envconfig:"PLUGIN_SOURCE_AWS_REGION"`
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
		args.SourcePassword, err = getAWSPassword(args.SourceAWSAccessKeyID, args.SourceAWSSecretAcessKey, args.SourceAWSRegion)
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
		headOptions := appendInsecure([]crane.Option{crane.WithAuth(destAuth)}, args.Insecure)
		_, err := crane.Head(args.Destination, headOptions...)
		if err == nil {
			return fmt.Errorf("image already exists in destination registry and overwrite is not enabled")
		}
	}

	pullOptions := appendInsecure([]crane.Option{crane.WithAuth(sourceAuth)}, args.Insecure)
	originImage, err := crane.Pull(args.Source, pullOptions...)
	if err != nil {
		return fmt.Errorf("failed to pull image %s: %w", args.Source, err)
	}

	pushOptions := appendInsecure([]crane.Option{crane.WithAuth(destAuth)}, args.Insecure)
	err = crane.Push(originImage, args.Destination, pushOptions...)
	if err != nil {
		return fmt.Errorf("failed to push image %s: %w", args.Destination, err)
	}

	return nil
}

func appendInsecure(options []crane.Option, insecure bool) []crane.Option {
	if insecure {
		return append(options, crane.Insecure)
	}
	return options
}
