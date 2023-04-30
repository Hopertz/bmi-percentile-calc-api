package main

import (
	"math"
	"strconv"

	"github.com/Hopertz/bmi-percentile-calc-api/internal/data"
)

type calcBmiObj struct {
	Bmi        float64
	Z_perc     int
	OverP95    *int
	M          float64
	Bmi_status string
}

var bmiagerev = data.GetBmidata()

func calcBMIandPerc_Metr(kgs, meters float64, sex string, agem float64) calcBmiObj {

	var calcBmiObj calcBmiObj
	calcBmiObj.Bmi = -1
	calcBmiObj.Z_perc = -1
	calcBmiObj.OverP95 = nil
	calcBmiObj.Bmi_status = ""

	bmi := kgs / (meters * meters)
	calcBmiObj.Bmi = math.Round(bmi*10) / 10

	intsex := 0
	if sex == "M" {
		intsex = 1
	} else {
		intsex = 2
	}
	for _, data := range bmiagerev {

		age, _ := strconv.ParseFloat(data.Agemos, 64)
		sex, _ := strconv.ParseInt(data.Sex, 10, 64)
		if int(sex) == intsex && ((agem + 0.5) == age) {

			M, _ := strconv.ParseFloat(data.M, 64)
			S, _ := strconv.ParseFloat(data.S, 64)
			L, _ := strconv.ParseFloat(data.L, 64)
			calcBmiObj.M = M

			bmi_z := (math.Pow(bmi/M, L) - 1) / (S * L)

			z_perc_not_rounded := GetZPercent(bmi_z) * 100

			z_perc := int(math.Round(z_perc_not_rounded))
			calcBmiObj.Z_perc = z_perc
			calcBmiObj.Bmi_status = giveBmiStatus(z_perc)

			if z_perc_not_rounded > 97 {
				p95, _ := strconv.ParseFloat(data.P95, 64)
				overP95 := int(math.Round(100 * bmi / p95))
				calcBmiObj.OverP95 = &overP95
			}

			break
		}
	}

	return calcBmiObj
}

func GetZPercent(z float64) float64 {
	if z < -6.5 {
		return 0.0
	}
	if z > 6.5 {
		return 1.0
	}

	factK := 1.0
	sum := 0.0
	term := 1.0
	k := 0.0
	loopStop := math.Exp(-23)

	for math.Abs(term) > loopStop {
		term = 0.3989422804 * math.Pow(-1, k) * math.Pow(z, k) / (2*k + 1) / math.Pow(2, k) * math.Pow(z, k+1) / factK
		sum += term
		k++
		factK *= k
	}

	sum += 0.5

	return sum
}

func giveBmiStatus(percentile int) string {

	var classification string
	switch {

	case percentile <= 5:
		classification = "underweight"

	case percentile <= 85:
		classification = "normal"

	case percentile <= 95:
		classification = "overweight"

	default:
		classification = "obesity"
	}

	return classification
}
