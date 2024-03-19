package user

type UserDto struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	Photo         string `json:"photo"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	CreditBalance int32  `json:"credit_balance"`
}
