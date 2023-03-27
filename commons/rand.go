package commons

const (
	r23 = 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5
	t23 = 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0
	r46 = r23 * r23
	t46 = t23 * t23
)

// TODO: verificar o pq essa funcao fica mais lenta usando a formula
func Vranlc(n int, xSeed *float64, a float64, y []float64) {
	ux := uint64(*xSeed)
	ua := uint64(a)
	const ut46 = uint64(t46)
	const it46 = 1 / t46

	for i := 0; i < n; i++ {
		ux = ux * ua % ut46
		y[i] = it46 * float64(ux)
	}
	*xSeed = float64(ux)
}

// TODO: verificar se para outras classes usar a formula pode causar problema
func Randlc(x *float64, a float64) float64 {
	ux := uint64(*x) * uint64(a) % uint64(t46)
	*x = float64(ux)
	const it46 = 1 / t46
	ret := it46 * (*x)
	return ret
}

/* formula original do jeito mais simples
   causa problemas nas operacoes com float
*x = math.Mod(a*(*x), t46)
return math.Pow(2, -46) * (*x)
*/
