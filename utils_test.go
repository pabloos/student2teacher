package main

import (
	"os"
	"testing"
)

func Test_endsWithSlash(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "withSlash",
			args: args{
				"path/",
			},
			want: true,
		},
		{
			name: "withoutSlash",
			args: args{
				"path",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := endsWithSlash(tt.args.path); got != tt.want {
				t.Errorf("endsWithSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_correctPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "withSlash",
			args: args{
				"path/",
			},
			want: "path/",
		},
		{
			name: "withoutSlash",
			args: args{
				"path",
			},
			want: "path/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := correctPath(tt.args.path); got != tt.want {
				t.Errorf("correctPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "just a dir",
			args: args{
				"dir/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createDir(tt.args.path)

			if !dirExists(tt.args.path) {
				t.Error("createDir() does not work")
			}

			os.Remove(tt.args.path) //dir is empty
		})
	}
}
