package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Data struct {
	Id        string  `json:"id"`
	Mac       string  `json:"mac"`
	Timestamp string  `json:"timestamp"`
	Long      float64 `json:"long"`
	Lat       float64 `json:"lat"`
	//Se almacenan los datos medidos como un map
	Params map[string]float64 `json:"params"`
}

var data []*Data = []*Data{}

func handleData(w http.ResponseWriter, r *http.Request) {

	m := Data{}
	var n = &m
	err := json.NewDecoder(r.Body).Decode(&n)
	var simData []Data
	c := make(chan []Data)

	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error al leer el cuerpo del mensaje")
		return
	}
	data = append(data, n)
	log.Println(data)
	fmt.Fprintf(w, "Module was added")
	simData = append(simData, m)

	log.Println(simData)

	go receiveData(c)
	c <- simData

}
func getModulos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(data[:])

}

//funciones para poder usar los datos reales segunda parte
func isIn(s string, arr []string) bool {
	for _, el := range arr {
		if el == s {
			return true
		}
	}
	return false
}

//funcion para poder recibir la data
func sendMetric(data []Data) {
	for _, v := range data {
		for k, val := range v.Params {
			for _, m := range metrics {
				if m.name == k {
					m.p.WithLabelValues(v.Id).Set(val)
				}
			}

		}

	}
}
func receiveData(c chan []Data) {
	var mods []string // array que almacena los modulos ya registrados
	for dataArray := range c {
		if !isIn(dataArray[0].Id, mods) {
			mods = append(mods, dataArray[0].Id) //agrega si el id del modulo no se encuentra registrado
			NewMetric(dataArray)
			sendMetric(dataArray)
		} else {
			sendMetric(dataArray)
		}
	}
}
func main() {
	r := mux.NewRouter()
	r.Handle("/metrics", promhttp.Handler()).Methods("GET")
	r.HandleFunc("/dataget", getModulos).Methods("GET")
	r.HandleFunc("/data", handleData).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
