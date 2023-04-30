package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {

	type input struct {
		Height float64 `json:"height"`
		Weight float64 `json:"weight"`
		Age    int     `json:"age"`
		Sex    string  `json:"sex"`
	}

	if r.Method == "POST" {
		var body input

		dec := json.NewDecoder(r.Body)

		err := dec.Decode(&body)

		if err != nil {
			w.Write([]byte("error in decoding the input"))
			return
		}
		bmi_obj := calcBMIandPerc_Metr(body.Weight, body.Height, body.Sex, float64(body.Age*12))
		js, err := json.Marshal(bmi_obj)
		if err != nil {
			log.Println(err)
			return
		}
		w.Write([]byte(js))
		return
	}
	w.Write([]byte("The api uses POST method only"))
}
