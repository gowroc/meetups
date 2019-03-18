package metrics

import "fmt"

var metrics = map[string]*metric{}

type metric struct {
	c int
}

func WithLabel(label string) *metric {
	v, ok := metrics[label]
	if !ok {
		v := &metric{c: 0}
		metrics[label] = v
		return v
	}
	return v
}

func (m *metric) Inc() {
	m.c++
}

func Print() {
	for k, v := range metrics {
		fmt.Printf("Label: %s, count %d\n", k, v.c)
	}
}
