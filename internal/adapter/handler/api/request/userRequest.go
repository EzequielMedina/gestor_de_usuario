package request

type UserRequest struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
