package request

import (
	"reflect"
	"testing"
)

func TestNewDaily(t *testing.T) {
	tests := []struct {
		name string
		want *Daily
	}{
		{
			name: "NewDaily",
			want: &Daily{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDaily(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDaily() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDaily_GetSentence(t *testing.T) {
	type args struct {
		m *Method
	}
	tests := []struct {
		name    string
		d       *Daily
		args    args
		wantErr bool
	}{
		{
			name: "Daily.GetSentence",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.GetSentence(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("Daily.GetSentence() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
