package models

//User mpdel
type User struct {
	ID           int    `json:"id,omitempty"`
	FirstName    string `json:"firstname,omitempty"`
	LastName     string `json:"lastname,omitempty"`
	Country      string `json:"country,omitempty"`
	EmailAddress string `json:"emailaddress,omitempty"`
}
