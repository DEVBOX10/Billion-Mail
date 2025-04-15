package maillog_stat

import (
	"fmt"
	"image"
	"image/png"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// CampaignEventHandler 处理活动事件
func CampaignEventHandler(w http.ResponseWriter, r *http.Request, encStr string) {
	ctx := context.Background()
	data, err := Decrypt(encStr)
	if err != nil {
		fmt.Fprint(w, "invalid data")
		return
	}

	// 验证数据
	if err := ValidateData(data); err != nil {
		fmt.Fprint(w, "invalid data -2")
		return
	}

	curTime := time.Now().Unix()
	postfixMessageID := SearchPostfixMessageIDByMessageID(data["message_id"].(string))

	if postfixMessageID == "" {
		postfixMessageID = ""
	}

	todayDate := time.Now().Format("20060102")

	if data["type"] == "open" {
		query := MaillogDBQuery("opened", todayDate)
		_, err := query.Insert(g.Map{
			"campaign_id":        data["campaign_id"],
			"log_time":           curTime,
			"recipient":          data["recipient"],
			"message_id":         data["message_id"],
			"postfix_message_id": postfixMessageID,
		})

		if err != nil {
			g.Log().Error(ctx, err)
		}

		// 创建1x1像素图片
		img := image.NewRGBA(image.Rect(0, 0, 1, 1))
		w.Header().Set("Content-Type", "image/png")
		png.Encode(w, img)
		return

	} else if data["type"] == "click" {
		query := MaillogDBQuery("clicked", todayDate)
		_, err := query.Insert(g.Map{
			"campaign_id":        data["campaign_id"],
			"log_time":           time.Now().Unix(),
			"recipient":          data["recipient"],
			"message_id":         data["message_id"],
			"postfix_message_id": postfixMessageID,
			"url":                data["url"],
		})

		if err != nil {
			g.Log().Error(ctx, err)
		}

		// 重定向到目标URL
		http.Redirect(w, r, data["url"].(string), http.StatusFound)
		return
	}

	fmt.Fprint(w, "success")
}

// MailTracker 邮件跟踪结构
type MailTracker struct {
	originalMailHTML string
	modified         bool
	mailHTML         string
	campaignID       int
	messageID        string
	recipient        string
	baseURL          string
	hrefPattern      *regexp.Regexp
}

// NewMailTracker 创建新的邮件跟踪器
func NewMailTracker(mailHTML string, campaignID int, messageID, recipient, baseURL string) *MailTracker {
	return &MailTracker{
		originalMailHTML: mailHTML,
		modified:         false,
		mailHTML:         mailHTML,
		campaignID:       campaignID,
		messageID:        messageID,
		recipient:        recipient,
		baseURL:          fmt.Sprintf("%s/v2/pmta", strings.TrimRight(baseURL, "/")),
		hrefPattern:      regexp.MustCompile(`href\s*=\s*"([^"]+)"`),
	}
}

// TrackLinks 处理邮件中的链接进行跟踪
func (t *MailTracker) TrackLinks() {
	t.mailHTML = t.hrefPattern.ReplaceAllStringFunc(t.mailHTML, func(s string) string {
		matches := t.hrefPattern.FindStringSubmatch(s)
		if len(matches) < 2 {
			return s
		}

		// 检查链接是否为有效URL
		u, err := url.Parse(matches[1])
		if err != nil || u.Scheme == "" || u.Host == "" {
			return s
		}

		return fmt.Sprintf(`href="%s"`, t.GetTrackingURL(matches[1]))
	})

	t.modified = true
}

// GetTrackingURL 获取跟踪URL
func (t *MailTracker) GetTrackingURL(url string) string {
	data := map[string]interface{}{
		"type":        "click",
		"campaign_id": t.campaignID,
		"recipient":   t.recipient,
		"message_id":  t.messageID,
		"url":         url,
	}

	return fmt.Sprintf("%s/%s", t.baseURL, Encrypt(data))
}

// AppendTrackingPixel 附加跟踪像素到邮件
func (t *MailTracker) AppendTrackingPixel() {
	trackingPixel := fmt.Sprintf(`<img src="%s" style="display:none" />`, t.GetTrackingPixel())

	if strings.Contains(t.mailHTML, "</body>") {
		t.mailHTML = strings.Replace(t.mailHTML, "</body>", trackingPixel+"</body>", 1)
		t.modified = true
		return
	}

	if strings.Contains(t.mailHTML, "</html>") {
		t.mailHTML = strings.Replace(t.mailHTML, "</html>", trackingPixel+"</html>", 1)
		t.modified = true
		return
	}

	t.mailHTML += trackingPixel
	t.modified = true
}

// GetTrackingPixel 获取跟踪像素URL
func (t *MailTracker) GetTrackingPixel() string {
	data := map[string]interface{}{
		"type":        "open",
		"campaign_id": t.campaignID,
		"recipient":   t.recipient,
		"message_id":  t.messageID,
	}

	return fmt.Sprintf("%s/%s", t.baseURL, Encrypt(data))
}

// IsModified 检查是否已修改
func (t *MailTracker) IsModified() bool {
	return t.modified
}

// GetOriginalHTML 获取原始HTML内容
func (t *MailTracker) GetOriginalHTML() string {
	return t.originalMailHTML
}

// GetHTML 获取当前HTML内容
func (t *MailTracker) GetHTML() string {
	return t.mailHTML
}
