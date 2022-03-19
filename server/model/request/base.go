package request

type Login struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	//Captcha   string `json:"captcha"`
	//CaptchaId string `json:"captchaId"`
}

type Signup struct {
	Account  string `json:"account"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
