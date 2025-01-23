package handlers

import (
	"encoding/json"
	"net/http"

	"ebyte-locker/database"
	"ebyte-locker/models"
)

func HandleGraphData(w http.ResponseWriter, r *http.Request) {
	database.Mutex.Lock()
	defer database.Mutex.Unlock()

	rows, err := database.DB.Query(`
		WITH RECURSIVE last_7_days(day) AS (
			SELECT DATE('now', '-6 days')
			UNION ALL
			SELECT DATE(day, '+1 day')
			FROM last_7_days
			WHERE day < DATE('now')
		)
		SELECT 
			day,
			COALESCE(SUM(CASE WHEN launch_date = day THEN 1 ELSE 0 END), 0) AS total_launches
		FROM last_7_days
		LEFT JOIN launches ON launches.launch_date = day
		GROUP BY day
		ORDER BY day ASC
	`)
	if err != nil {
		http.Error(w, "Failed to fetch graph data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var labels []string
	var data []int
	for rows.Next() {
		var date string
		var totalLaunches int
		if err := rows.Scan(&date, &totalLaunches); err != nil {
			http.Error(w, "Failed to process graph data", http.StatusInternalServerError)
			return
		}
		labels = append(labels, date)
		data = append(data, totalLaunches)
	}

	graphData := models.GraphData{
		Labels: labels,
		Data:   data,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(graphData)
}

func HandleMiniGraphsData(w http.ResponseWriter, r *http.Request) {
	database.Mutex.Lock()
	defer database.Mutex.Unlock()

	rows, err := database.DB.Query(`
		WITH RECURSIVE last_7_days(day) AS (
			SELECT DATE('now', '-6 days')
			UNION ALL
			SELECT DATE(day, '+1 day')
			FROM last_7_days
			WHERE day < DATE('now')
		)
		SELECT 
			day,
			COALESCE(SUM(CASE WHEN lockers.launches > 0 THEN 1 ELSE 0 END), 0) AS total_lockers,
			COALESCE(SUM(CASE WHEN launches.launch_date = day THEN 1 ELSE 0 END), 0) AS total_launches,
			COALESCE(SUM(CASE WHEN lockers.last_infection_date = day THEN lockers.infections ELSE 0 END), 0) AS total_infections
		FROM last_7_days
		LEFT JOIN lockers ON DATE(lockers.last_infection_date) = day
		LEFT JOIN launches ON launches.launch_date = day
		GROUP BY day
		ORDER BY day ASC
	`)
	if err != nil {
		http.Error(w, "Failed to fetch mini graph data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var labels []string
	var lockerData, launchData, infectionData []int
	for rows.Next() {
		var date string
		var totalLockers, totalLaunches, totalInfections int
		if err := rows.Scan(&date, &totalLockers, &totalLaunches, &totalInfections); err != nil {
			http.Error(w, "Failed to process mini graph data", http.StatusInternalServerError)
			return
		}
		labels = append(labels, date)
		lockerData = append(lockerData, totalLockers)
		launchData = append(launchData, totalLaunches)
		infectionData = append(infectionData, totalInfections)
	}

	miniGraphData := models.MiniGraphData{
		Lockers: models.GraphData{
			Labels: labels,
			Data:   lockerData,
		},
		Launches: models.GraphData{
			Labels: labels,
			Data:   launchData,
		},
		Infections: models.GraphData{
			Labels: labels,
			Data:   infectionData,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(miniGraphData)
}
