package triangle

import (
	"math"
)

const testVersion = 2

func KindFromSides(a, b, c float64) Kind {
	if !IsValidTriangle(a, b, c) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	} else if a == b || b == c || a == c {
		return Iso
	} else {
		return Sca
	}
}

func IsValidTriangle(a, b, c float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return false
	}
	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return false
	}
	if math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1) {
		return false
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	if (a+b) < c || (b+c) < a || (a+c) < b {
		return false
	}
	return true
}

type Kind string

const NaT = "NaT"
const Equ = "Equ"
const Iso = "Iso"
const Sca = "Sca"
