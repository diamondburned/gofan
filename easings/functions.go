// Package easings implements all functions in easings.net.
package easings

import "math"

const (
	c1 = 1.70158
	c2 = c1 * 1.525
	c3 = c1 + 1
	c4 = (2 * math.Pi) / 3
	c5 = (2 * math.Pi) / 4.5
)

// Functions maps function names to the respective function.
var Functions = map[string]func(float64) float64{
	"Linear":           Linear,
	"EaseInQuad":       EaseInQuad,
	"EaseOutQuad":      EaseOutQuad,
	"EaseInOutQuad":    EaseInOutQuad,
	"EaseInCubic":      EaseInCubic,
	"EaseOutCubic":     EaseOutCubic,
	"EaseInOutCubic":   EaseInOutCubic,
	"EaseInQuart":      EaseInQuart,
	"EaseOutQuart":     EaseOutQuart,
	"EaseInOutQuart":   EaseInOutQuart,
	"EaseInQuint":      EaseInQuint,
	"EaseOutQuint":     EaseOutQuint,
	"EaseInOutQuint":   EaseInOutQuint,
	"EaseInSine":       EaseInSine,
	"EaseOutSine":      EaseOutSine,
	"EaseInOutSine":    EaseInOutSine,
	"EaseInExpo":       EaseInExpo,
	"EaseOutExpo":      EaseOutExpo,
	"EaseInOutExpo":    EaseInOutExpo,
	"EaseInCirc":       EaseInCirc,
	"EaseOutCirc":      EaseOutCirc,
	"EaseInOutCirc":    EaseInOutCirc,
	"EaseInBack":       EaseInBack,
	"EaseOutBack":      EaseOutBack,
	"EaseInOutBack":    EaseInOutBack,
	"EaseInElastic":    EaseInElastic,
	"EaseOutElastic":   EaseOutElastic,
	"EaseInOutElastic": EaseInOutElastic,
	"EaseInBounce":     EaseInBounce,
	"EaseOutBounce":    EaseOutBounce,
	"EaseInOutBounce":  EaseInOutBounce,
}

// CurveFunctions is a subset of Functions that contains only curvy functions.
var CurveFunctions = map[string]func(float64) float64{
	"Linear":         Linear,
	"EaseInQuad":     EaseInQuad,
	"EaseOutQuad":    EaseOutQuad,
	"EaseInOutQuad":  EaseInOutQuad,
	"EaseInCubic":    EaseInCubic,
	"EaseOutCubic":   EaseOutCubic,
	"EaseInOutCubic": EaseInOutCubic,
	"EaseInQuart":    EaseInQuart,
	"EaseOutQuart":   EaseOutQuart,
	"EaseInOutQuart": EaseInOutQuart,
	"EaseInQuint":    EaseInQuint,
	"EaseOutQuint":   EaseOutQuint,
	"EaseInOutQuint": EaseInOutQuint,
	"EaseInSine":     EaseInSine,
	"EaseOutSine":    EaseOutSine,
	"EaseInOutSine":  EaseInOutSine,
	"EaseInExpo":     EaseInExpo,
	"EaseOutExpo":    EaseOutExpo,
	"EaseInOutExpo":  EaseInOutExpo,
	"EaseInCirc":     EaseInCirc,
	"EaseOutCirc":    EaseOutCirc,
	"EaseInOutCirc":  EaseInOutCirc,
}

// BounceOut is probably used internally mostly.
func BounceOut(x float64) float64 {
	const n1 = 7.5625
	const d1 = 2.75

	if x < 1/d1 {
		return n1 * x * x
	} else if x < 2/d1 {
		x -= 1.5
		return n1*(x/d1)*x + 0.75
	} else if x < 2.5/d1 {
		x -= 2.25
		return n1*(x/d1)*x + 0.9375
	} else {
		x -= 2.625
		return n1*(x/d1)*x + 0.984375
	}
}

func Linear(x float64) float64 { return x }

func EaseInQuad(x float64) float64 {
	return x * x
}

func EaseOutQuad(x float64) float64 {
	return 1 - (1-x)*(1-x)
}

func EaseInOutQuad(x float64) float64 {
	if x < 0.5 {
		return 2 * x * x
	}
	return 1 - math.Pow(-2*x+2, 2)/2
}

func EaseInCubic(x float64) float64 {
	return x * x * x
}

func EaseOutCubic(x float64) float64 {
	return 1 - math.Pow(1-x, 3)
}

func EaseInOutCubic(x float64) float64 {
	if x < 0.5 {
		return 4 * x * x * x
	}
	return 1 - math.Pow(-2*x+2, 3)/2
}

func EaseInQuart(x float64) float64 {
	return x * x * x * x
}

func EaseOutQuart(x float64) float64 {
	return 1 - math.Pow(1-x, 4)
}

func EaseInOutQuart(x float64) float64 {
	if x < 0.5 {
		return 8 * x * x * x * x
	}
	return 1 - math.Pow(-2*x+2, 4)/2
}

func EaseInQuint(x float64) float64 {
	return x * x * x * x * x
}

func EaseOutQuint(x float64) float64 {
	return 1 - math.Pow(1-x, 5)
}

func EaseInOutQuint(x float64) float64 {
	if x < 0.5 {
		return 16 * x * x * x * x * x
	}
	return 1 - math.Pow(-2*x+2, 5)/2
}

func EaseInSine(x float64) float64 {
	return 1 - math.Cos((x*math.Pi)/2)
}

func EaseOutSine(x float64) float64 {
	return math.Sin((x * math.Pi) / 2)
}

func EaseInOutSine(x float64) float64 {
	return -(math.Cos(math.Pi*x) - 1) / 2
}

func EaseInExpo(x float64) float64 {
	if x == 0 {
		return 0
	}
	return math.Pow(2, 10*x-10)
}

func EaseOutExpo(x float64) float64 {
	if x == 1 {
		return 1
	}
	return 1 - math.Pow(2, -10*x)
}

func EaseInOutExpo(x float64) float64 {
	switch {
	case x == 0:
		return 0
	case x == 1:
		return 1
	case x < 0.5:
		return math.Pow(2, 20*x-10) / 2
	default:
		return (2 - math.Pow(2, -20*x+10)) / 2
	}
}

func EaseInCirc(x float64) float64 {
	return 1 - math.Sqrt(1-math.Pow(x, 2))
}

func EaseOutCirc(x float64) float64 {
	return math.Sqrt(1 - math.Pow(x-1, 2))
}

func EaseInOutCirc(x float64) float64 {
	if x < 0.5 {
		return (1 - math.Sqrt(1-math.Pow(2*x, 2))) / 2
	}
	return (math.Sqrt(1-math.Pow(-2*x+2, 2)) + 1) / 2
}

func EaseInBack(x float64) float64 {
	return c3*x*x*x - c1*x*x
}

func EaseOutBack(x float64) float64 {
	return 1 + c3*math.Pow(x-1, 3) + c1*math.Pow(x-1, 2)
}

func EaseInOutBack(x float64) float64 {
	if x < 0.5 {
		return (math.Pow(2*x, 2) * ((c2+1)*2*x - c2)) / 2
	}
	return (math.Pow(2*x-2, 2)*((c2+1)*(x*2-2)+c2) + 2) / 2
}

func EaseInElastic(x float64) float64 {
	switch x {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return -math.Pow(2, 10*x-10) * math.Sin((x*10-10.75)*c4)
	}
}

func EaseOutElastic(x float64) float64 {
	switch x {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return math.Pow(2, -10*x)*math.Sin((x*10-0.75)*c4) + 1
	}
}

func EaseInOutElastic(x float64) float64 {
	switch {
	case x == 0:
		return 0
	case x == 1:
		return 1
	case x < 0.5:
		return -(math.Pow(2, 20*x-10) * math.Sin((20*x-11.125)*c5)) / 2
	default:
		return (math.Pow(2, -20*x+10)*math.Sin((20*x-11.125)*c5))/2 + 1
	}
}

func EaseInBounce(x float64) float64 {
	return 1 - BounceOut(1-x)
}

func EaseOutBounce(x float64) float64 { return BounceOut(x) }

func EaseInOutBounce(x float64) float64 {
	if x < 0.5 {
		return (1 - BounceOut(1-2*x)) / 2
	}
	return (1 + BounceOut(2*x-1)) / 2
}
