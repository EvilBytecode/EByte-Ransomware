package models

type DashboardData struct {
	TotalLockers       int    `json:"total_lockers"`
	TotalLaunches      int    `json:"total_launches"`
	InfectionsToday    int    `json:"infections_today"`
	LaunchToBuildRatio string `json:"launch_to_build_ratio"`
}

type MiniGraphData struct {
	Lockers    GraphData `json:"lockers"`
	Launches   GraphData `json:"launches"`
	Infections GraphData `json:"infections"`
}

type GraphData struct {
	Labels []string `json:"labels"`
	Data   []int    `json:"data"`
}

type LaunchData struct {
	LockerID string `json:"locker_id"`
}
