// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import "testing"

func TestValidateArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    Args
		wantErr bool
	}{
		{
			name: "valid",
			args: Args{
				Username:    "user",
				Password:    "pass",
				Source:      "source",
				Destination: "destination",
			},
			wantErr: false,
		},
		{
			name: "invalid username",
			args: Args{
				Username:    "",
				Password:    "pass",
				Source:      "src",
				Destination: "dest",
			},
			wantErr: true,
		},
		{
			name: "invalid password",
			args: Args{
				Username:    "user",
				Password:    "",
				Source:      "src",
				Destination: "dest",
			},
			wantErr: true,
		},
		{
			name: "invalid source",
			args: Args{
				Username:    "user",
				Password:    "pass",
				Source:      "",
				Destination: "dest",
			},
			wantErr: true,
		},
		{
			name: "invalid destination",
			args: Args{
				Username:    "user",
				Password:    "pass",
				Source:      "src",
				Destination: "",
			},
			wantErr: true,
		},
		{
			name: "aws source",
			args: Args{
				Username:    "AWS",
				Password:    "",
				Source:      "src",
				Destination: "dest",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := validateArgs(&tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateArgs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
