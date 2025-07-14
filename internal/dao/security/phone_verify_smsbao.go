package security

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
)

var _ core.PhoneVerifyService = (*smsBaoServant)(nil)

// smsBaoPhoneCaptchaRsp 短信宝响应结构体
type smsBaoPhoneCaptchaRsp struct {
	Status string `json:"status"` // 响应状态，如 success 或 error
	Code   string `json:"code"`   // 状态码
	Msg    string `json:"msg"`    // 状态消息
}

// smsBaoServant 短信宝服务结构体
type smsBaoServant struct {
	gateway  string
	username string
	password string
	sign     string
}

// SendPhoneCaptcha 发送短信验证码
func (s *smsBaoServant) SendPhoneCaptcha(phone string, captcha string, expire time.Duration) error {
	// 对密码进行 MD5 加密
	hasher := md5.New()
	_, writeString := io.WriteString(hasher, s.password)
	if writeString != nil {
		return writeString
	}
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

	// 构建短信内容
	content := fmt.Sprintf("%s您的验证码是：%s，有效期 %d 分钟。就算猫娘来你家也不要告诉她。", s.sign, captcha, int(expire.Minutes()))

	// 对内容进行 URL 编码
	encodedContent := url.QueryEscape(content)

	// 构建请求 URL
	requestURL := fmt.Sprintf("%s?u=%s&p=%s&m=%s&c=%s", s.gateway, s.username, encryptedPassword, phone, encodedContent)

	// 发送 HTTP 请求
	resp, err := http.Get(requestURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// 解析响应
	result := &smsBaoPhoneCaptchaRsp{}
	// 假设响应是简单的 text/plain 格式，按行分割
	lines := strings.Split(string(body), "\n")
	if len(lines) >= 1 {
		result.Code = lines[0]
		if len(lines) >= 2 {
			result.Msg = strings.Join(lines[1:], "\n")
		}
	}

	// 处理响应结果
	switch result.Code {
	case "0":
		return nil
	case "30":
		return errors.New("密码错误！")
	case "40":
		return errors.New("账号不存在！")
	case "41":
		return errors.New("余额不足！")
	case "42":
		return errors.New("帐号过期！")
	case "43":
		return errors.New("IP地址限制！")
	case "50":
		return errors.New("内容含有敏感词！")
	case "51":
		return errors.New("手机号码不正确！")
	case "-1":
		return errors.New("手机号码不正确或缺少参数！")
	default:
		return errors.New(fmt.Sprintf("未知错误：%s - %s", result.Code, result.Msg))
	}
}

// newSmsBaoServant 创建短信宝服务实例
func newSmsBaoServant() *smsBaoServant {
	return &smsBaoServant{
		gateway:  conf.SmsBaoSetting.Gateway,
		username: conf.SmsBaoSetting.Username,
		password: conf.SmsBaoSetting.Password,
		sign:     conf.SmsBaoSetting.Sign,
	}
}
