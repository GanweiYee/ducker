package request

type LoginForm struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	//Captcha   string `json:"captcha"`
	//CaptchaId string `json:"captchaId"`
}
