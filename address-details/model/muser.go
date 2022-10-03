package model

// import "go.mongodb.org/mongo-driver/bson/primitive"

type MUser struct {
	// Id        primitive.ObjectID `json:"id,omitempty"`
	EmailId   string `json:"emailid"`
	Password  string `json:"password"`
	FirstName string `json:"firstname`
	LastName  string `json:"lastname"`
}
