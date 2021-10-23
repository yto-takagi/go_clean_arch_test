package auth

// struct
type SignUp struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
