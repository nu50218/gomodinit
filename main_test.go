package main

import "testing"

func Test_fixRemoteRepoName(t *testing.T) {
	type args struct {
		repo string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "github_https",
			args: args{
				repo: "https://github.com/nu50218/gomodinit.git",
			},
			want:    "github.com/nu50218/gomodinit",
			wantErr: false,
		},
		{
			name: "github_ssh",
			args: args{
				repo: "git@github.com:nu50218/gomodinit.git",
			},
			want:    "github.com/nu50218/gomodinit",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fixRemoteRepoName(tt.args.repo)
			if (err != nil) != tt.wantErr {
				t.Errorf("fixRemoteRepoName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fixRemoteRepoName() = %v, want %v", got, tt.want)
			}
		})
	}
}
