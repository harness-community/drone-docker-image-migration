// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import "testing"

func TestValidateArgs(t *testing.T) {
	type args struct {
		Args Args
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			args: args{Args: Args{
				SourceRegistry:       "registry.example.com",
				SourceUsername:       "user",
				SourcePassword:       "pass",
				SourceNamespace:      "namespace",
				ImageName:            "image",
				ImageTag:             "tag",
				DestinationRegistry:  "registry.example.com",
				DestinationUsername:  "user",
				DestinationPassword:  "pass",
				DestinationNamespace: "namespace",
			}},
			wantErr: false,
		},
		{
			name: "invalid source registry",
			args: args{Args: Args{
				SourceRegistry:      "",
				SourceUsername:      "user",
				SourcePassword:      "pass",
				ImageName:           "image",
				ImageTag:            "tag",
				DestinationRegistry: "registry.example.com",
				DestinationUsername: "user",
				DestinationPassword: "pass",
			}},
			wantErr: true,
		},
		{
			name: "invalid source username",
			args: args{Args: Args{
				SourceRegistry:      "registry.example.com",
				SourceUsername:      "",
				SourcePassword:      "pass",
				ImageName:           "image",
				ImageTag:            "tag",
				DestinationRegistry: "registry.example.com",
				DestinationUsername: "user",
				DestinationPassword: "pass",
			}},
			wantErr: true,
		},
		{
			name: "invalid source password",
			args: args{Args: Args{
				SourceRegistry:      "registry.example.com",
				SourceUsername:      "user",
				SourcePassword:      "",
				ImageName:           "image",
				ImageTag:            "tag",
				DestinationRegistry: "registry.example.com",
				DestinationUsername: "user",
				DestinationPassword: "pass",
			}},
			wantErr: true,
		},
		{
			name: "invalid image name",
			args: args{Args: Args{
				SourceRegistry:      "registry.example.com",
				SourceUsername:      "user",
				SourcePassword:      "pass",
				ImageName:           "",
				ImageTag:            "tag",
				DestinationRegistry: "registry.example.com",
				DestinationUsername: "user",
				DestinationPassword: "pass",
			}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := ValidateArgs(tt.args.Args); (err != nil) != tt.wantErr {
				t.Errorf("ValidateArgs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoginToRegistry(t *testing.T) {
	type args struct {
		Username string
		Password string
		Registry string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "invalid username",
			args:    args{Username: "", Password: "pass", Registry: "registry.example.com"},
			wantErr: true,
		},
		{
			name:    "invalid password",
			args:    args{Username: "user", Password: "", Registry: "registry.example.com"},
			wantErr: true,
		},
		{
			name:    "invalid registry",
			args:    args{Username: "user", Password: "pass", Registry: ""},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := LoginToRegistry(tt.args.Username, tt.args.Password, tt.args.Registry); (err != nil) != tt.wantErr {
				t.Errorf("LoginToRegistry() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAWSPassword(t *testing.T) {
	if _, err := GetAWSPassword(
		"accessKeyID",
		"secretAccessKey",
		"region",
	); err == nil {
		t.Errorf("GetAWSPassword() error = %v", err)
	}
}

func TestCopyImage(t *testing.T) {
	if err := CopyImage(
		"sourceImage",
		"destinationImage",
	); err == nil {
		t.Errorf("CopyImage() error = %v", err)
	}
}
