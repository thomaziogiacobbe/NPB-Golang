package commons

import "math"

const (
	r23 = 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5
	t23 = 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0
	r46 = r23 * r23
	t46 = t23 * t23
)

// TODO: verificar o pq essa funcao fica mais lenta usando a formula
func Vranlc(n int, xSeed *float64, a float64, y []float64) {

	var x float64

	x = *xSeed
	for i := 0; i < n; i++ {
		ux := uint64(x)
		ua := uint64(a)
		mul := ux * ua
		x = float64(mul % uint64(t46))
		y[i] = math.Pow(2, -46) * (x)
	}
	*xSeed = x
}

// TODO: verificar se para outras classes usar a formula pode causar problema
func Randlc(x *float64, a float64) float64 {
	ux := uint64(*x)
	ua := uint64(a)
	mul := ux * ua
	*x = float64(mul % uint64(t46))
	ret := math.Pow(2, -46) * (*x)
	return ret
}

/* formula original do jeito mais simples
   causa problemas nas operacoes com float
*x = math.Mod(a*(*x), t46)
return math.Pow(2, -46) * (*x)
*/
