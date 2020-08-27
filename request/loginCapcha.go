package request

type RegisterAndLoginStruct struct {
	LoginName string `json:"loginname"`
	Pwd       string `json:"pwd"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}
