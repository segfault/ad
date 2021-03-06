// mathematical function templates
//
// they come in pairs: function per se, and its derivative.
// the derivative function has "d_" prepended to function name.
// if function has more than one argument, then the first derivative
// argument is the index we're taking derivative with respect to
//
package templates

import "math"

func tan(a float64) float64 {
	return math.Tan(a)
}
func d_tan(a float64) float64 {
	return math.Pow(1/math.Cos(a), 2)
}

func abs(a float64) float64 {
	return math.Abs(a)
}
func d_abs(a float64) float64 {
	switch {
	case a > 0:
		return +1
	case a < 0:
		return -1
	default:
		panic("illegal derivative when abs(0)=0")
	}
}

func atan(a float64) float64 {
	return math.Atan(a)
}
func d_atan(a float64) float64 {
	return 1 / (1 + a*a)
}

func tanh(a float64) float64 {
	return math.Tanh(a)
}
func d_tanh(a float64) float64 {
	return 1 - math.Pow(math.Tanh(a), 2)
}

func sin(a float64) float64 {
	return math.Sin(a)
}
func d_sin(a float64) float64 {
	return math.Cos(a)
}

func asin(a float64) float64 {
	return math.Asin(a)
}
func d_asin(a float64) float64 {
	return 1 / math.Sqrt(1-a*a)
}

func sinh(a float64) float64 {
	return math.Sinh(a)
}
func d_sinh(a float64) float64 {
	return math.Cosh(a)
}

func cos(a float64) float64 {
	return math.Cos(a)
}
func d_cos(a float64) float64 {
	return -math.Sin(a)
}

func acos(a float64) float64 {
	return math.Acos(a)
}
func d_acos(a float64) float64 {
	return -1 / math.Sqrt(1-a*a)
}

func cosh(a float64) float64 {
	return math.Cosh(a)
}
func d_cosh(a float64) float64 {
	return math.Sinh(a)
}

func add(a, b float64) float64 {
	return a + b
}
func d_add(i int, a, b float64) float64 {
	if i < 0 || i > 1 {
		panic("illegal index")
	}
	return 1
}

func multiply(a, b float64) float64 {
	return a * b
}
func d_multiply(i int, a, b float64) float64 {
	switch i {
	case 0:
		return b
	case 1:
		return a
	default:
		panic("illegal index")
	}
}

func subtract(a, b float64) float64 {
	return a - b
}
func d_subtract(i int, a, b float64) float64 {
	switch i {
	case 0:
		return 1
	case 1:
		return -1
	default:
		panic("illegal index")
	}
}

func divide(a, b float64) float64 {
	return a / b
}
func d_divide(i int, a, b float64) float64 {
	switch i {
	case 0:
		return 1 / b
	case 1:
		return -a / (b * b)
	default:
		panic("illegal index")
	}
}

func sqrt(a float64) float64 {
	return math.Sqrt(a)
}
func d_sqrt(a float64) float64 {
	return 0.5 * math.Pow(a, -0.5)
}

func exp(a float64) float64 {
	return math.Exp(a)
}
func d_exp(a float64) float64 {
	return math.Exp(a)
}

// 2^a --- sort of like exponentiation base 2
func exp2(a float64) float64 {
	return math.Exp2(a)
}
func d_exp2(a float64) float64 {
	return math.Log(2) * math.Pow(2, a)
}

func log(a float64) float64 {
	return math.Log(a)
}
func d_log(a float64) float64 {
	return 1 / a
}

func log2(a float64) float64 {
	return math.Log2(a)
}
func d_log2(a float64) float64 {
	return 1 / (a * math.Log(2))
}

func pow(a, b float64) float64 {
	return math.Pow(a, b)
}
func d_pow(i int, a, b float64) float64 {
	switch i {
	case 0:
		return b * math.Pow(a, b-1)
	case 1:
		return math.Pow(a, b) * math.Log(a)
	default:
		panic("illegal index")
	}
}
