package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"database/sql"
	"ebyte-locker/database"
	"ebyte-locker/models"
	"ebyte-locker/utils"
)



func HandleLaunch(w http.ResponseWriter, r *http.Request) {
	var data models.LaunchData

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	database.Mutex.Lock()
	defer database.Mutex.Unlock()

	err := utils.CheckDuplicate(database.DB, "launches", "locker_id", data.LockerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	_, err = database.DB.Exec(`
		INSERT INTO launches (locker_id, launch_date)
		VALUES (?, DATE('now'))
	`, data.LockerID)
	if err != nil {
		http.Error(w, "Failed to record launch", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec(`
		UPDATE lockers
		SET launches = launches + 1
		WHERE id = ?
	`, data.LockerID)
	if err != nil {
		http.Error(w, "Failed to update locker launches", http.StatusInternalServerError)
		return
	}

	var lastInfectionDate sql.NullString
	row := database.DB.QueryRow(`SELECT last_infection_date FROM lockers WHERE id = ?`, data.LockerID)
	if err := row.Scan(&lastInfectionDate); err != nil {
		http.Error(w, "Failed to fetch last infection date", http.StatusInternalServerError)
		return
	}

	currentDate := time.Now().Format("2006-01-02")
	if !lastInfectionDate.Valid || lastInfectionDate.String != currentDate {
		_, err := database.DB.Exec(`
			UPDATE lockers
			SET infections = infections + 1, last_infection_date = ?
			WHERE id = ?
		`, currentDate, data.LockerID)
		if err != nil {
			http.Error(w, "Failed to update infections", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Launch data recorded successfully"))
}


func HandleGenerateLocker(w http.ResponseWriter, r *http.Request) {
	lockerID := fmt.Sprintf("locker-%d", time.Now().UnixNano())

	go func() {
		err := GenerateLocker(lockerID)
		if err != nil {
			fmt.Printf("Error generating locker: %v\n", err)
		}
	}()

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(fmt.Sprintf("Locker generation started with ID: %s", lockerID)))
}

