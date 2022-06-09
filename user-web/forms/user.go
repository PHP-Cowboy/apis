package forms

type PasswordLoginForm struct {
	Mobile    string `json:"mobile" binding:"required,mobile"`
	Password  string `json:"password" binding:"required,min=6,max=20"`
	Captcha   string `json:"captcha" binding:"required,len=5"`
	CaptchaId string `json:"captcha_id" binding:"required"`
}

type RegisterForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式有规范可寻， 自定义validator
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Code     string `form:"code" json:"code" binding:"required,min=6,max=6"`
}

type UpdateUserForm struct {
	Name     string `form:"name" json:"name" binding:"required,min=3,max=10"`
	Gender   uint32 `form:"gender" json:"gender" binding:"required,oneof=1 2"`
	Birthday string `form:"birthday" json:"birthday" binding:"required,datetime=2006-01-02"`
}
