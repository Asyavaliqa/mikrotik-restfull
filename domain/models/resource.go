package models

type Resource struct {
	Platform          string `json:"platform"`
	Board             string `json:"board-name"`
	Uptime            string `json:"uptime"`
	Version           string `json:"version"`
	BuildTime         string `json:"build-time"`
	FreeMemory        string `json:"free-memory"`
	TotalMemory       string `json:"total-memory"`
	Cpu               string `json:"cpu"`
	CpuCount          string `json:"cpu-count"`
	CpuFrequency      string `json:"cpu-frequency"`
	CpuLoad           string `json:"cpu-load"`
	FreeHddSpace      string `json:"free-hdd-space"`
	TotalHddSpace     string `json:"total-hdd-space"`
	BadBlocks         string `json:"bad-blocks"`
	ArchitectureNames string `json:"architecture-name"`
}
