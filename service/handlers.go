package service

import (
	"encoding/csv"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hatorikibble/diversity_calendar/model"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetHoliday(w http.ResponseWriter, r *http.Request) {

	// Read the 'accountId' path parameter from the mux map
	var date = mux.Vars(r)["date"]

	d := readSourcefile(date)
	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(d)
	writeJsonResponse(w, http.StatusOK, data)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	data, _ := json.Marshal(model.HealthCheckResponse{Status: "UP"})
	writeJsonResponse(w, http.StatusOK, data)
}

func writeJsonResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}

// check panics if an error is detected
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readSourcefile(date_string string) model.Holiday {
	f, err := os.Open("/home/peter/gocode/src/github.com/hatorikibble/diversity_calendar/Diversity_Kalender_2017.csv")
	check(err)
	defer f.Close()

	lineCount := 0

	r := csv.NewReader(f)
	r.Comma = ';'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			check(err)
		}
		lineCount += 1
		s := model.Holiday{Date: record[0], Name: strings.TrimSpace(record[1]), Religion: record[2]}
		if s.Date == date_string {
			return s

		}

	}
	return model.Holiday{}

}
