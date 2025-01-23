package main

import (
	"log"
	"net/http"

	"ebyte-locker/database"
	"ebyte-locker/handlers"
)

func main() {
	database.InitDB()
	defer database.DB.Close()

	http.Handle("/", http.FileServer(http.Dir("./frontend")))

	http.HandleFunc("/dashboard-data", handlers.HandleDashboardData)
	http.HandleFunc("/launch", handlers.HandleLaunch)
	http.HandleFunc("/graph-data", handlers.HandleGraphData)
	http.HandleFunc("/mini-graphs-data", handlers.HandleMiniGraphsData)
	http.HandleFunc("/generate-locker", handlers.HandleGenerateLocker)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
