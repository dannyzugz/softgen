package models

type ProjectData struct {
	ProjectName string `json:"project"`
	ModName     string `json:"mod"`
	RouterName  string `json:"router"`
	DbName      string `json:"db"`
	Ui          bool   `json:"ui"`
}
