package lfserve

type Mappings struct {
	Pairs map[int]int
	Free  map[int]struct{}
}

var m *Mappings = nil

func NewMap() {
	m = &Mappings{Pairs: make(map[int]int), Free: make(map[int]struct{})}
}

func GetMap() *Mappings {
	return m
}
