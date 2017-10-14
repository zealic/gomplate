package funcs

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var (
	mathNS     *MathFuncs
	mathNSInit sync.Once
)

// MathNS - the math namespace
func MathNS() *MathFuncs {
	mathNSInit.Do(func() { mathNS = &MathFuncs{} })
	return mathNS
}

// AddMathFuncs -
func AddMathFuncs(f map[string]interface{}) {
	f["math"] = MathNS
}

// MathFuncs -
type MathFuncs struct{}

// Add -
func (f *MathFuncs) Add(x, y interface{}) (interface{}, error) {
	a, err := parseFloat(x)
	if err != nil {
		return 0, err
	}
	b, err := parseFloat(y)
	if err != nil {
		return 0, err
	}
	return a + b, err
}

func parseFloat(in interface{}) (n float64, err error) {
	if f, ok := in.(float64); ok {
		return f, nil
	}
	if i, ok := in.(int64); ok {
		return float64(i), nil
	}
	if u, ok := in.(uint64); ok {
		return float64(u), nil
	}
	if s, ok := in.(string); ok {
		ss := strings.Split(s, ".")
		if len(ss) > 2 {
			return 0, fmt.Errorf("can not parse '%s' as a number - too many decimal points", s)
		}
		n, err = strconv.ParseFloat(s, 64)
		return n, err
	}
	if s, ok := in.(fmt.Stringer); ok {
		return parseFloat(s.String())
	}
	if in == nil {
		return 0, nil
	}
	return 0, nil
}
