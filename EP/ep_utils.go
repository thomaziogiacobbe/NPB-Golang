package main

import (
	"fmt"
	"math"
	"os"
)

func getNPBClass(class string) {
	switch class {
	case "S":
		M = 24
	case "W":
		M = 25
	case "A":
		M = 28
	case "B":
		M = 30
	case "C":
		M = 32
	case "D":
		M = 36
	case "E":
		M = 40
	default:
		fmt.Println("Incorrect class argument")
		os.Exit(1)
	}
}

func verify(sx float64, sy float64) bool {
	const EPSILON = 1.0e-8
	var (
		verified        bool
		sx_verify_value float64
		sy_verify_value float64
		sx_err          float64
		sy_err          float64
	)
	verified = true
	switch M {
	case 24:
		sx_verify_value = -3.247834652034740e+3
		sy_verify_value = -6.958407078382297e+3
	case 25:
		sx_verify_value = -2.863319731645753e+3
		sy_verify_value = -6.320053679109499e+3
	case 28:
		sx_verify_value = -4.295875165629892e+3
		sy_verify_value = -1.580732573678431e+4
	case 30:
		sx_verify_value = 4.033815542441498e+4
		sy_verify_value = -2.660669192809235e+4
	case 32:
		sx_verify_value = 4.764367927995374e+4
		sy_verify_value = -8.084072988043731e+4
	case 36:
		sx_verify_value = 1.982481200946593e+5
		sy_verify_value = -1.020596636361769e+5
	case 40:
		sx_verify_value = -5.319717441530e+05
		sy_verify_value = -3.688834557731e+05
	default:
		verified = false
	}

	if verified {
		sx_err = math.Abs((sx - sx_verify_value) / sx_verify_value)
		sy_err = math.Abs((sy - sy_verify_value) / sy_verify_value)
		verified = (sx_err <= EPSILON) && (sy_err <= EPSILON)
	}
	return verified
}
