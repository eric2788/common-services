package bilibili

import (
	"fmt"

	"github.com/eric2788/common-utils/request"
)

func GetRoomInfo(roomId int64) (*RoomInfo, error) {

	var resp RoomInfo

	_, err := liveRequester.Get("/room/v1/Room/get_info", &resp,
		request.Query(map[string]interface{}{
			"room_id": roomId,
		}),
	)

	if err != nil {
		return nil, err
	} else if resp.Code != 0 {
		return nil, fmt.Errorf("bilibili: %s", resp.Message)
	}

	return &resp, nil
}
