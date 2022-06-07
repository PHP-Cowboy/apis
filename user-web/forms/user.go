package forms

type PasswordLoginForm struct {
	Mobile    string `json:"mobile" binding:"required,mobile"`
	Password  string `json:"password" binding:"required,min=6,max=20"`
	Captcha   string `json:"captcha" binding:"required,len=5"`
	CaptchaId string `json:"captcha_id" binding:"required"`
}
