package models

// UserRequest is used for POST and PUT validations
type UserRequest struct {
	Name string `json:"name" validate:"required"`
	Dob  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

// UserResponse includes the dynamically calculated age
type UserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Dob  string `json:"dob"`
	Age  int    `json:"age,omitempty"` // Omitempty hides it if age is 0 (for create/update responses)
}
