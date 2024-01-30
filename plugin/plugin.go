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

	args.SourceUsername = "AWS"
	// args.Password = "eyJwYXlsb2FkIjoieWcrckN4VlM3NEdicUUzeERIK3FGQ2dYQ0p1ZjlYQlN5RngwOVQwemtzaytldG5ZQndkLzRWcUNNSE1LZUV0V0drM280UnFMVWx5UDRnVHc1Nzhha05HRGljd2hXZjlDWWsxUnFDWEt6akZualVOd1FJc09YbUhiT3dkUUI2bDhuc24vYzI3SjNoem1abmVFandaQ2dQNjB4UTZiU1A2ODd0aHFMRW1KTVk4S0lpTU9jR0R4OUxYZFUyMmV1M3BGcWt4WWt2MEd3MXBnVTR2SWVwNzhpWlRuMXRkaWRRSnRBZEhiTlZIRVlVb3NRMlFJaWxITkdpUUhXeUlkSmM3cU5RKzVOY2JTY29PM002VTJYWVo5ZnJrcTZzaVUrMURhNE9Balp4K01LSEt1QzRoRFFCa0tyaStkN3J0NFZjeENkMDB4LzFTbzdCRUVtVHBEb2d2QVA5Tzl2RkF1c3g3eU1lRTRYNFpSRG8yODZvV3BSSmNQaHlBa3FBeWRxR0NGK3hLREdqN1N0K25OdHRhZStlUWpCU2RrMEx4U01XU1NKSUUxSW1FV1oxeU1mSmh0ekh5NEtlZHZobHIwcVJmRjI0c2p4WlZmWVlsRnZSYlM3K3RMNEdYRjJOZU9qd2h2S2lrQzVkS1I0ZkI2QzRVWEVhcEkxUWpzVmdySysrd3R6Vm16ZFovT0h4dWo0SnUrNGxYUlkrSmk3dENDMVFmeit6YzROcFhuYmlSS045c05KRTQxSFBVZmhVWXdJR05xam4vZFRtTko1UmYxQkpCQlhzU215dGt1ZElHZDVYQVhtRXhkV2FHY3ZWR0tFTXM5K1k5VmhpeUFURlhNRzRZR1RNYno1WU1NSjBmWFBvM3pLWU51aC8waE4xOWhlTlo1QndwVWxmZTFwYzVEMnhaS2h2b2lHQm1ySGVJcnBrNTVTSmZpVElBUXJMU2JMRDc2L0VudFl4U3NFNUMwWnF5TzhhRm1zdm02c05WWFZpMmIrVEVONk05d2lneWsydGNjZEdMUjFjeFJHNzNvYXUxVEdBUE13V0VCdE5PeWltaENvM0FsN3AwRzFnbjluZXVnT0RyMThyQndPS2d3d2VZYXQ2Yk5RUUVKUi80eWRJMVFhY2FwRXoweXNoc0ZNbjd2bldvZFR4RWlKQSs3UlZLY29PMXM4STRwQWhzQ3B5S056Vk9oM1E5dUZnZHB3RlRUNFBhQ3hSVVZOeTMrRzUwV0NGdXA0UkRBQnlQSEEyR2xpb1cvUXJ2VUxtRzg0dXFNU3V2eUpxc3JlUDNNTWkyVjVDdm1KTFROeTlxaEU4S0pmUEVMdXRCVGMvWW9tamF1Qm1iSjFMSks0bGRlMlNCZklRM2srREhjQWRoZE1PS0NUeVdsYldRPSIsImRhdGFrZXkiOiJBUUlCQUhoZnJGUTJhSUl3RjVWYzc5UkhFUEduY1VXdXIyL0p6TTh0L1p6b0V0UnJBZ0gwMEZWa1BTaWZjdnEyMm91UDI4cm9BQUFBZmpCOEJna3Foa2lHOXcwQkJ3YWdiekJ0QWdFQU1HZ0dDU3FHU0liM0RRRUhBVEFlQmdsZ2hrZ0JaUU1FQVM0d0VRUU05a0JNd3QzaEFGU2NjWHlmQWdFUWdEdlhQQmtnWm5taGVWM1JQQ0FuOTFQcUtUSHRNcXMweGYyK3RWaW9ncW5UZ2dWZmQ4UW9FZ2hRWDYrdmZSUE01MkFXU3NnR25OQ2ZMRWdnbHc9PSIsInZlcnNpb24iOiIyIiwidHlwZSI6IkRBVEFfS0VZIiwiZXhwaXJhdGlvbiI6MTcwNjY0MjUwMX0="
	args.Destination = "akshitagrawal/helm-test"
	args.Source = "751209113259.dkr.ecr.eu-north-1.amazonaws.com/myrepo"
	args.Username = "akshitagrawal"
	args.Password = "Akshit@2010@"
	args.Overwrite = true
	args.AwsAccessKeyID = "AKIA25Z4NM2VWF4ZXXV4"
	args.AwsSecretAcessKey = "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"
	args.AwsRegion = "eu-north-1"

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
