package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"unit-converter/internal/conversion"
	"unit-converter/internal/model"
)

func respondWithError(w http.ResponseWriter, code int, message string){
	respondWithJSON(w, code, model.ErrorResponse{Error: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		// Fallback error if JSON marshalling fails
		http.Error(w, `{"error": "Internal server error marshalling JSON"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ConversionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var value float64
	var fromUnit string
	var toUnit string
	var err error

	if r.Method == http.MethodPost {
		var payload model.RequestPayload

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid JSON payload in request body: "+ err.Error())
			return
		}

		value = payload.Value
		fromUnit = strings.ToLower(payload.UnitFrom)
		toUnit = strings.ToLower(payload.UnitTo)

		if fromUnit == "" || toUnit == "" {
			respondWithError(w, http.StatusBadRequest, "Missing required fields in JSON payload: 'value', 'unitFrom', or 'unitTo'")
			return
		}
	} else if r.Method == http.MethodGet {
		params := r.URL.Query()

		valueStr := params.Get("value")
		fromUnit = strings.ToLower(params.Get("from"))
		toUnit = strings.ToLower(params.Get("to"))

		if valueStr == "" || fromUnit == "" || toUnit == "" {
			respondWithError(w, http.StatusBadRequest, "Missing required query parameters: 'value', 'from', or 'to'")
			return
		}

		value, err = strconv.ParseFloat(valueStr, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid value provided: '%s' is not a number", valueStr))
			return
		}
	} else {
		respondWithError(w, http.StatusMethodNotAllowed, "Only GET and POST mehtods are supported")
		return
	}

	var targetValue float64
	var conversionType string
	

	switch {
	case strings.Contains("celsiusfahrenheitkelvin", fromUnit) || strings.Contains("celsiusfahrenheitkelvin", toUnit):
		// Temperature
		targetValue, err = conversion.HandleTemperatureConversion(value, fromUnit, toUnit)
		conversionType = "Temperature"
	case strings.Contains("meterfootinchkilometercentimetermillimeteryardmile", fromUnit) || strings.Contains("meterfootinchkilometercentimetermillimeteryardmile", toUnit):
		// Length
		targetValue, err = conversion.HandleLengthConversion(value, fromUnit, toUnit)
		conversionType = "Length"
	case strings.Contains("kilogrampoundgrammilligramounce", fromUnit) || strings.Contains("kilogrampoundgrammilligramounce", toUnit):
		// Weight
		targetValue, err = conversion.HandleWeightConversion(value, fromUnit, toUnit)
		conversionType = "Weight"
	default:
		// Unknown or mismatched units
		err = fmt.Errorf("could not determine conversion type for units '%s' and '%s'", fromUnit, toUnit)
	}

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := model.APIResponse{
		OriginalValue:  value,
		FromUnit: 	   fromUnit,
		TargetValue:    targetValue,
		ToUnit: 	   toUnit,
		ConversionType: conversionType,
	}

	respondWithJSON(w, http.StatusOK, response)
}