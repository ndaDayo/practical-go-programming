package method

type Struct struct {
	v int
}

func (s Struct) SetValue(v int) {
	s.v = v
}

func (s *Struct) SetValuePointer(v int) {
	s.v = v
}
