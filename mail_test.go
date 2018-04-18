package utils

import "testing"

func TestSendMail(t *testing.T) {

	SendMail("13683515835@163.com",
		"", //使用邮箱设置的发送密码
		"smtp.163.com:25",
		"系统问题反馈",
		"内容哈",
		[]string{"chenzhen@cmcm.com"})

}
