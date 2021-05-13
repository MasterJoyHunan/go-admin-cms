package service

import (
	"blog/pkg/setting"
	"github.com/mojocn/base64Captcha"
	"github.com/satori/go.uuid"
	"image/color"
)

type CaptchaConfig struct {
	Id           string
	CaptchaType  string
	VerifyValue  string
	DriverString *base64Captcha.DriverString
}

// 创建新验证码
func NewCaptcha(id string) (image base64Captcha.Item, err error) {
	e := CaptchaConfig{}
	e.Id = uuid.NewV4().String()
	e.DriverString = base64Captcha.NewDriverString(
		setting.CaptchaConf.Height,
		setting.CaptchaConf.Width,
		setting.CaptchaConf.Noise,
		setting.CaptchaConf.Line,
		setting.CaptchaConf.Length,
		setting.CaptchaConf.Str,
		&color.RGBA{240, 240, 246, 246},
		setting.CaptchaConf.Font,
	)
	driver := e.DriverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	_, content, answer := captcha.Driver.GenerateIdQuestionAnswer()
	item, err := captcha.Driver.DrawCaptcha(content)
	if err != nil {
		return nil, err
	}
	captcha.Store.Set(id, answer)
	return item, nil
}
