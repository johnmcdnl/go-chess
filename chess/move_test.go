package chess

import (
	"reflect"
	"testing"
)

func TestNewMove(t *testing.T) {
	type args struct {
		origin      *Square
		destination *Square
	}
	tests := []struct {
		name    string
		args    args
		want    *Move
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMove(tt.args.origin, tt.args.destination)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMove_isValid(t *testing.T) {
	type fields struct {
		origin        *Square
		destination   *Square
		IsCaptureMove bool
		name          string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Move{
				origin:        tt.fields.origin,
				destination:   tt.fields.destination,
				IsCaptureMove: tt.fields.IsCaptureMove,
				name:          tt.fields.name,
			}
			if err := m.isValid(); (err != nil) != tt.wantErr {
				t.Errorf("Move.isValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMove_generateName(t *testing.T) {
	type fields struct {
		origin        *Square
		destination   *Square
		IsCaptureMove bool
		name          string
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Move{
				origin:        tt.fields.origin,
				destination:   tt.fields.destination,
				IsCaptureMove: tt.fields.IsCaptureMove,
				name:          tt.fields.name,
			}
			m.generateName()
		})
	}
}

func TestMove_Name(t *testing.T) {
	type fields struct {
		origin        *Square
		destination   *Square
		IsCaptureMove bool
		name          string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Move{
				origin:        tt.fields.origin,
				destination:   tt.fields.destination,
				IsCaptureMove: tt.fields.IsCaptureMove,
				name:          tt.fields.name,
			}
			if got := m.Name(); got != tt.want {
				t.Errorf("Move.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
