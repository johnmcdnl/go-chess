package chess

import "testing"

func TestFEN_IsValid(t *testing.T) {
	tests := []struct {
		name string
		f    *FEN
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.IsValid(); got != tt.want {
				t.Errorf("FEN.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
