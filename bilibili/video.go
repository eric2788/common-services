package bilibili

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/eric2788/common-utils/request"
)

var (
	shortURLLink = regexp.MustCompile(`https?:\/\/b23\.tv\/(\w+)\/?`)
)


func GetVideoInfo(bvid string) (*VideoInfo, error) {
	var resp VideoInfo
	_, err := apiRequester.Get("/web-interface/view/detail", &resp,
		request.Query(map[string]interface{}{
			"bvid": bvid,
		}),
		request.Headers(map[string]string{
			"Referer": "https://www.bilibili.com/video/" + bvid,
			"Origin": "https://www.bilibili.com",
		}),
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func ParseShortLink(url string) (string, error) {
	if !shortURLLink.MatchString(url) {
		return "", fmt.Errorf("不是 bilibili 短链接: %s", url)
	}

	matches := shortURLLink.FindStringSubmatch(url)
	if len(matches) < 2 {
		return "", fmt.Errorf("不是 bilibili 短链接: %s", url)
	}

	if strings.HasPrefix(matches[1], "BV") {
		return "", fmt.Errorf("BV號格式無效: %s", url)
	}

	link := matches[0]

	s, err := getRedirectLink(link)
	if err != nil {
		logger.Errorf("解析 bilibili 短链接 %s 时出现错误: %v", link, err)
	} else {
		url = strings.ReplaceAll(url, link, s)
	}

	return url, nil
}

func getRedirectLink(shortURL string) (string, error) {
	resp, err := request.New().Raw().Get(shortURL, request.Headers(map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64",
	}))
	if err != nil {
		return "", err
	}
	defer resp.Resp.Body.Close()
	return resp.GetRequestUrl(), nil
}