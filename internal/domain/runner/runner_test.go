package runner

import (
	"reflect"
	"testing"
)

func TestNewRunner(t *testing.T) {
	type args struct {
		name         string
		emailAddress string
	}
	tests := []struct {
		name    string
		args    args
		want    Runner
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRunner(tt.args.name, tt.args.emailAddress)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRunner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRunner() got = %v, want %v", got, tt.want)
			}
		})
	}
}
