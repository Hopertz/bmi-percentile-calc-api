package data

import (
	_ "embed"
	"encoding/json"
	"log"
)

type bmi struct {
	Agemos string
	Sex    string
	L      string
	M      string
	S      string
	P95    string
}

//go:embed bmi_data.json
var bmi_data []byte

func GetBmidata() []bmi {
	var bmiagerev []bmi
	err := json.Unmarshal(bmi_data, &bmiagerev)

	if err != nil {
		log.Println(err)
	}

	return bmiagerev
}
