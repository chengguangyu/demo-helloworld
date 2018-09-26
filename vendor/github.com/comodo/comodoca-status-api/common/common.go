package common

type StatusResponse struct {
	ServiceName        string         `json:"name"`
	ServiceDescription string         `json:"description"`
	Status             string         `json:"status"` //Available, Unavailable, Pending
	SubComponents      []SubComponent `json:"subComponents"`
	VersionNumber      string         `json:"version"`
}

type SubComponent struct {
	ComponentName   string       `json:"name"`   //eg database, http config
	ComponentStatus string       `json:"status"` // stacktrace status
	ConfigItems     []ConfigItem `json:"configItems"`
}

type ConfigItem struct {
	ConfigName        string `json:"name"` //db host name, db user, http port, https certificate
	ConfigDescription string `json:"description"`
	ConfigStatus      string `json:"status"`
	ConfigMessage     string `json:"message"`
}
