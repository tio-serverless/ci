package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {

	var c = `log="debug"
token="xyz123"
[repos]
  branch=["master","develop"]
  name=["c1","c2"]`

	f, _ := ioutil.TempFile("", "")

	f.WriteString(c)

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    config
		wantErr bool
	}{
		{
			name: "config",
			args: args{path: f.Name()},
			want: config{
				Log:   "debug",
				Token: "xyz123",
				Repos: struct {
					Branch []string `toml:"branch"`
					Name   []string `toml:"name"`
				}{
					Branch: []string{"master", "develop"},
					Name:   []string{"c1", "c2"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfig(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
