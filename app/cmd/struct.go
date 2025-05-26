package cmd

type DogResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
type DogListResponse struct {
	Message []string `json:"message"`
	Status  string   `json:"status"`
}
