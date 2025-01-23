package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ebyte-locker/database"
	"ebyte-locker/models"
)

func HandleDashboardData(w http.ResponseWriter, r *http.Request) {
	database.Mutex.Lock()
	defer database.Mutex.Unlock()

	row := database.DB.QueryRow(`
		SELECT 
			COUNT(*) as total_lockers, 
			COALESCE(SUM(launches), 0) as total_launches,
			COALESCE(SUM(CASE WHEN last_infection_date = DATE('now') THEN infections ELSE 0 END), 0) as infections_today
		FROM lockers
	`)
	var totalLockers, totalLaunches, infectionsToday int
	err := row.Scan(&totalLockers, &totalLaunches, &infectionsToday)
	if err != nil {
		http.Error(w, "Failed to fetch dashboard data", http.StatusInternalServerError)
		return
	}

	launchToBuildRatio := "0%"
	if totalLockers > 0 {
		launchToBuildRatio = fmt.Sprintf("%.1f%%", (float64(totalLaunches)/float64(totalLockers))*100)
	}

	data := models.DashboardData{
		TotalLockers:       totalLockers,
		TotalLaunches:      totalLaunches,
		InfectionsToday:    infectionsToday,
		LaunchToBuildRatio: launchToBuildRatio,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
