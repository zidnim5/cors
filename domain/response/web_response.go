package domain

type WebResponse struct {
	Data    string `json:"data"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Success bool   `json:"success"`
}
