package bilibili

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/eric2788/common-utils/request"
)

func wRid(uid, wts int64) (string, error) {
	// update the salt key if required
	checkShouldUpdateSaltKey()
	if saltKey == nil || saltKey.Key == "" {
		return "", fmt.Errorf("salt key is empty")
	}
	c := saltKey.Key
	b := "mid=" + fmt.Sprint(uid) + "&platform=web&token=&web_location=1550101"
	a := b + "&wts=" + fmt.Sprint(wts) + c // mid + platform + token + web_location + 时间戳wts + 一个固定值
	hash := md5.Sum([]byte(a))
	return hex.EncodeToString(hash[:]), nil
}

func GetUserInfo(uid int64) (*UserInfo, error) {

	var resp UserInfo

	wts := time.Now().Unix()
	w_rid, err := wRid(uid, wts)
	if err != nil {
		return nil, err
	}

	_, err = requester.Get("https://api.bilibili.com/x/space/wbi/acc/info", &resp,
		request.Query(map[string]interface{}{
			"mid": uid,
			"platform": "web",
			"token": "",
			"web_location": "1550101",
			"wts": wts,
			"w_rid": w_rid,
		}),
	)

	if err != nil {
		return nil, err
	} else if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Message)
	}

	return &resp, nil
}