package main

const r23 = (0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5 * 0.5)
const t23 = (2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0 * 2.0)
const r46 = (r23 * r23)
const t46 = (t23 * t23)

//TODO: x_seed precisa ser uma passagem por referencia
func vranlc(n int, x_seed *float64, a float64, y []float64) {

	var x, t1, t2, t3, t4, a1, a2, x1, x2, z float64

	t1 = r23 * a
	a1 = t1
	a2 = a - t23*a1
	x = *x_seed

	for i := 0; i < n; i++ {
		t1 = r23 * x
		x1 = t1
		x2 = x - t23*x1
		t1 = a1*x2 + a2*x1
		t2 = r23 * t1
		z = t1 - t23*t2
		t3 = t23*z + a2*x2
		t4 = r46 * t3
		x = t3 - t46*t4
		y[i] = r46 * x
	}
	*x_seed = x
}

//TODO: x precisa ser uma passagem por referencia
func randlc(x *float64, a float64) float64 {
	var t1, t2, t3, t4, a1, a2, x1, x2, z float64

	t1 = r23 * a
	a1 = t1
	a2 = a - t23*a1

	t1 = r23 * (*x)
	x1 = t1
	x2 = (*x) - t23*x1
	t1 = a1*x2 + a2*x1
	t2 = r23 * t1
	z = t1 - t23*t2
	t3 = t23*z + a2*x2
	t4 = r46 * t3
	*x = t3 - t46*t4

	return r46 * (*x)
}
