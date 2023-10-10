package skyutl

type ServiceAddr struct {
	CoreService string `json:"coreService"`
	OrbitHrm    string `json:"orbitHrm"`
	OrbitTask   string `json:"orbitTask"`
	OrbitReport string `json:"orbitReport"`
	OrbitAuto   string `json:"orbitAuto"`
	Skyins      string `json:"skyins"`
	Skycmn      string `json:"skycmn"`
	Skyinv      string `json:"skyinv"`
	Skyatc      string `json:"skyatc"`
	Skyrpt      string `json:"skyrpt"`
}
