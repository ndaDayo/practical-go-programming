package method

import "testing"

func TestSetValue(t *testing.T) {
	s := Struct{v: 0}

	s.SetValue(100000)

	if s.v != 0 {
		t.Errorf("SetValue failed, got: %d, want: 0", s.v)
	}
}

func TestSetValuePointer(t *testing.T) {
	s := Struct{v: 0}

	s.SetValuePointer(10)

	if s.v != 10 {
		t.Errorf("SetValuePointer failed, got: %d, want: 0", s.v)
	}
}
