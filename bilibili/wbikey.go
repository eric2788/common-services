package bilibili

import (
	"strings"
	"time"

	"github.com/eric2788/common-utils/request"
)

var mixinKeyEncTab = []byte{
	46, 47, 18, 2, 53, 8, 23, 32, 15, 50, 10, 31, 58, 3, 45, 35, 27, 43, 5, 49,
	33, 9, 42, 19, 29, 28, 14, 39, 12, 38, 41, 13, 37, 48, 7, 16, 24, 55, 40,
	61, 26, 17, 0, 1, 60, 51, 30, 4, 22, 25, 54, 21, 56, 59, 6, 63, 57, 62, 11,
	36, 20, 34, 44, 52,
}

type SaltKeyStorage struct {
	Key        string
	LastUpdate time.Time
}

var saltKey *SaltKeyStorage

func checkShouldUpdateSaltKey() {
	// key is null or expired for more than 24 hours
	if saltKey == nil || time.Since(saltKey.LastUpdate) > time.Hour*24 {
		UpdateSaltKey()
	}
}

// UpdateSaltKey enable update it manually
func UpdateSaltKey() {

	imgKey, subKey := getWbiKey()
	if imgKey == "" || subKey == "" {
		return
	}

	key := mixinKey(imgKey + subKey)

	saltKey = &SaltKeyStorage{
		Key:        key,
		LastUpdate: time.Now(),
	}
}

// getWbiKey get img_key and sub_key
func getWbiKey() (string, string) {
	var resp NavInfo
	_, err := apiRequester.Get("/web-interface/nav", &resp, request.DataDecoder(request.JsonDecoder)) // avoid error throw while resp.Code is not 0
	if err != nil {
		logger.Errorf("Error when getting wbi key: %v", err)
		return "", ""
	}
	if resp.Data == nil {
		logger.Errorf("Error when getting wbi key: %v", resp.Message)
		return "", ""
	}
	return extractKey(resp.Data.WbiImg.ImgUrl), extractKey(resp.Data.WbiImg.SubUrl)
}

func mixinKey(key string) string {
	var result string
	for _, v := range mixinKeyEncTab {
		result += string(key[v])
	}
	return result[:32]
}

func extractKey(url string) string {
	start := strings.LastIndex(url, "/") + 1
	end := len(url)
	chars := []rune(url)
	return strings.Split(string(chars[start:end]), ".")[0]
}

func init() {
	checkShouldUpdateSaltKey()
}
