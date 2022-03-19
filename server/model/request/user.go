package request

type Pwd struct {
	Account  string `json:"account"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
