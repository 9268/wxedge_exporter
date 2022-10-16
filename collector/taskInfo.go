package collector

type Task struct {
	Name            string  `json:"name"`
	ID              string  `json:"id"`
	StateCode       int     `json:"state_code"`
	StateMessage    string  `json:"state_message"`
	IP              string  `json:"ip"`
	Mac             string  `json:"mac"`
	Speed           float64 `json:"speed"`
	Disk            int     `json:"disk"`
	Mem             float64 `json:"mem"`
	CPUUsage        float64 `json:"cpu_usage"`
	StartTime       int     `json:"start_time"`
	NeedRunningDays int     `json:"need_running_days"`
}
