package price

import "testing"

func TestPrice(t *testing.T) {
	tests := []struct {
		season Season
		got    int
		want   int
	}{
		{Peak, 1000, 1200},
		{Normal, 1000, 1000},
		{Off, 1000, 1000},
	}

	for _, tt := range tests {
		got := tt.season.Price(tt.got)

		if got != tt.want {
			t.Errorf("Price(%v, %v) = %v; want %v", tt.season, tt.got, got, tt.want)
		}
	}
}
