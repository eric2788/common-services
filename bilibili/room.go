package bilibili

import (
	"fmt"

	"github.com/eric2788/common-utils/request"
)

func GetRoomInfo(roomId int64) (*RoomInfo, error) {

	var resp RoomInfo

	_, err := requester.Get("https://api.live.bilibili.com/room/v1/Room/get_info", &resp,
		request.Query(map[string]interface{}{
			"room_id": roomId,
		}),
		request.Headers(map[string]string{
			"Referer": "https://live.bilibili.com/",
			"Origin":  "https://live.bilibili.com/",
			"Host":    "api.live.bilibili.com",
		}))

	if err != nil {
		return nil, err
	} else if resp.Code != 0 {
		return nil, fmt.Errorf("bilibili: %s", resp.Message)
	}

	return &resp, nil
}
