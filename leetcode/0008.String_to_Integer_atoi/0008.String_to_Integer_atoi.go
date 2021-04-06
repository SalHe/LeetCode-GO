package leetcode

import (
	"math"
)

func myAtoi(s string) int {
	number := 0
	hasStarted := false
	hasSigned := false
	isNegative := false

	for _, c := range s {
		if c >= '0' && c <= '9' {
			num := int(c - '0')

			// 溢出检测
			if !isNegative && math.MaxInt32-number*10 < num {
				number = math.MaxInt32
			} else if isNegative && math.MaxInt32+1-number*10 < num {
				number = math.MaxInt32 + 1
			} else {
				number = number*10 + num
			}

			hasStarted = true
		} else if c == '-' || c == '+' {
			if hasStarted {
				break
			}
			if hasSigned {
				return 0
			}
			isNegative = c == '-'
			hasSigned = true
			hasStarted = true
		} else if c == ' ' {
			if hasStarted {
				break
			}
			continue
		} else {
			break
		}
	}

	if isNegative {
		return -number
	}
	return number
}

type atoiAutomatonState int

const (
	start atoiAutomatonState = iota
	signed
	in_number
	end
)

type atoiAutomaton struct {
	sign              int
	number            int64
	currentState      atoiAutomatonState
	stateConvertTable map[atoiAutomatonState][]atoiAutomatonState
}

func NewAtoiAutomaton() *atoiAutomaton {
	auto := &atoiAutomaton{
		sign:         1,
		number:       0,
		currentState: start,
		stateConvertTable: map[atoiAutomatonState][]atoiAutomatonState{
			start: {
				start, signed, in_number, end,
			},
			signed: {
				end, end, in_number, end,
			},
			in_number: {
				end, end, in_number, end,
			},
			end: {
				end, end, end, end,
			},
		},
	}
	return auto
}

func (a *atoiAutomaton) Input(ch rune) {
	col := 0
	if ch == ' ' {
		col = 0
	} else if ch == '+' || ch == '-' {
		col = 1
	} else if ch >= '0' && ch <= '9' {
		col = 2
	} else {
		col = 3
	}
	a.currentState = a.stateConvertTable[a.currentState][col]

	if a.currentState == in_number {
		a.number = a.number*10 + int64(ch-'0')
		if a.sign == 1 {
			a.number = min(math.MaxInt32, a.number)
		} else {
			a.number = min(math.MaxInt32+1, a.number)
		}
	} else if a.currentState == signed {
		if ch == '+' {
			a.sign = 1
		} else {
			a.sign = -1
		}
	}
}

func (a *atoiAutomaton) GetAns() int {
	return int(a.number) * a.sign
}

func min(a int64, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func myAtoi_Automaton(s string) int {
	auto := NewAtoiAutomaton()
	for _, ch := range s {
		auto.Input(ch)
	}
	return auto.GetAns()
}
