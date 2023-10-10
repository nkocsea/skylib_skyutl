package remote_service

type ServiceAddr struct {
	CoreService string `json:"coreService"`
	File        string `json:"file"`
	Report      string `json:"report"`
	Skyins      string `json:"skyins"`
	Skycmn      string `json:"skycmn"`
	Skyinv      string `json:"skyinv"`
	Skyatc      string `json:"skyatc"`
	Skyreg      string `json:"skyreg"`
	Skyimg      string `json:"skyimg"`
	Skyemr      string `json:"skyemr"`
	Skylab      string `json:"skylab"`
	Skyacc      string `json:"skyacc"`
	Skyrpt      string `json:"skyrpt"`
}
