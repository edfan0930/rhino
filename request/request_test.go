package request

import (
	"net/url"
	"reflect"
	"testing"
)

func TestNewMethod(t *testing.T) {
	u := ""

	type args struct {
		u string
	}
	tests := []struct {
		name string
		args args
		want *Method
	}{
		{
			name: "NewMethod",
			args: args{
				u: u,
			},
			want: &Method{
				URL:     u,
				Timeout: 10,
				Query:   url.Values{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMethod(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMethod_AddQueryParam(t *testing.T) {

	key := "ping"
	value := "pong"

	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		m    *Method
		args args
	}{
		{
			name: "Method.AddQueryParam",
			m:    NewMethod(""),
			args: args{
				key,
				value,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.AddQueryParam(tt.args.key, tt.args.value)

			if v := tt.m.Query.Get(key); v != value {
				t.Errorf("Method.AddQueryParam() got different key value")
			}

		})
	}
}

func TestMethod_Get(t *testing.T) {
	tests := []struct {
		name    string
		m       *Method
		want    []byte
		want1   int
		wantErr bool
	}{
		{
			name:    "Method.Get",
			m:       MetaphorpsumMethod,
			want1:   200,
			wantErr: false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.m.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Method.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("Method.Get() got nil")
			}
			if got1 != tt.want1 {
				t.Errorf("Method.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
