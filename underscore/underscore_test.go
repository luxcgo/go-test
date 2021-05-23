package underscore

import "testing"

func TestCamel(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Camel(tt.args.str); got != tt.want {
				t.Errorf("Camel() = %v, want %v", got, tt.want)
			}
		})
	}
}
