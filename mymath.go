package mymath

import "math"

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Ceil(x float64) float64 {
	return math.Ceil(x)
}

func Floor(x float64) float64 {
	return math.Floor(x)
}

func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

func Max(x, y float64) float64 {
	return math.Max(x, y)
}

func Min(x, y float64) float64 {
	return math.Min(x, y)
}

func Abs(x float64) float64 {
	return math.Abs(x)
}

func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}

func Sin(x float64) float64 {
	return math.Sin(x)
}

func Acos(x float64) float64 {
	return math.Acos(x)
}

func Asin(x float64) float64 {
	return math.Asin(x)
}

func Acosh(x float64) float64 {
	return math.Acosh(x)
}

func Asinh(x float64) float64 {
	return math.Asinh(x)
}

func Atan(x float64) float64 {
	return math.Atan(x)
}

func Atanh(x float64) float64 {
	return math.Atanh(x)
}

func Cos(x float64) float64 {
	return math.Cos(x)
}

func Float32frombits(x uint32) float32 {
	return math.Float32frombits(x)
}

func Atan2(x, y float64) float64 {
	return math.Atan2(x, y)
}

func Cbrt(x float64) float64 {
	return math.Cbrt(x)
}

func Copysing(x, y float64) float64 {
	return math.Copysign(x, y)
}

func Cosh(x float64) float64 {
	return math.Cosh(x)
}

func Dim(x, y float64) float64 {
	return math.Dim(x, y)
}

func Erf(x, y float64) float64 {
	return math.Erf(x)
}

func Exp(x float64) float64 {
	return math.Exp(x)
}

func Erfc(x float64) float64 {
	return math.Erfc(x)
}

func Erfcinv(x float64) float64 {
	return math.Erfcinv(x)
}

func Erfinv(x float64) float64 {
	return math.Erfinv(x)
}

func Exp2(x float64) float64 {
	return math.Exp2(x)
}

func Float64bits(x float64) uint64 {
	return math.Float64bits(x)
}

func Float64frombits(x uint64) float64 {
	return math.Float64frombits(x)
}

func Expm1(x float64) float64 {
	return math.Expm1(x)
}

func FMA(x, z, y float64) float64 {
	return math.FMA(x, y, z)
}

func Frexp(x float64) (float64, int) {
	return math.Frexp(x)
}

func Gamma(x float64) float64 {
	return math.Gamma(x)
}

func Inf(x int) float64 {
	return math.Inf(x)
}

func Ilogb(x float64) int {
	return math.Ilogb(x)
}

func IsInf(x float64, y int) bool {
	return math.IsInf(x, y)
}

func IsNaN(x float64) bool {
	return math.IsNaN(x)
}

func J0(x float64) float64 {
	return math.J0(x)
}

func J1(x float64) float64 {
	return math.J1(x)
}

func Jn(n int, x float64) float64 {
	return math.Jn(n, x)
}

func Ldexp(x float64, exp int) float64 {
	return math.Ldexp(x, exp)
}

func Lgamma(x float64) (float64, int) {
	return math.Lgamma(x)
}

func Hypot(x, y float64) float64 {
	return math.Hypot(x, y)
}

func Log(x float64) float64 {
	return math.Log(x)
}

func Log1p(x float64) float64 {
	return math.Log1p(x)
}

func Log2(x float64) float64 {
	return math.Log2(x)
}

func Log10(x float64) float64 {
	return math.Log10(x)
}

func Logb(x float64) float64 {
	return math.Logb(x)
}

func Mod(x, y float64) float64 {
	return math.Mod(x, y)
}

func Modf(x float64) (float64, float64) {
	return math.Modf(x)
}

func NaN() float64 {
	return math.NaN()
}

func Nextafter(x, y float64) float64 {
	return math.Nextafter(x, y)
}

func Nextafter32(x, y float32) float32 {
	return math.Nextafter32(x, y)
}

func Pow10(x int) float64 {
	return math.Pow10(x)
}
