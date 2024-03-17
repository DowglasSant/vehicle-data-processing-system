package api

import (
	"encoding/json"
	"net/http"
	"receiver-service/model"
	"receiver-service/service"
)

func HandleVehicleReceived(w http.ResponseWriter, r *http.Request) {
	var vehicle model.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := service.ProcessVehicleData(vehicle)
	if err != nil {
		http.Error(w, "Failed to process vehicle data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Vehicle data processed successfully")
}
