package model

type MContact struct {
	EmailId      string `json:"emailid"`
	Relationship string `json:"relationship"`
	Name         string `json:"name"`
	ContactType  string `json:"contacttype"`
	PhoneNumber  string `json:"phonenumber"`
}
