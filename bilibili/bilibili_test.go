package bilibili

import (
	"encoding/json"
	"testing"
)




func TestGetRoomInfo(t *testing.T) {

	info, err := GetRoomInfo(545)
	if err != nil {
		t.Skip(err)
	}

	t.Logf("%+v", info.V1Resp)

	b, err := json.MarshalIndent(*info.Data, "", "  ")
	if err != nil {
		t.Skip(err)
	}
	t.Log(string(b))
}


func TestGetUserInfo(t *testing.T) {
	
	info, err := GetUserInfo(1838190318)
	if err != nil {
		t.Skip(err)
	}

	t.Logf("%+v", info.XResp)

	b, err := json.MarshalIndent(*info.Data, "", "  ")
	if err != nil {
		t.Skip(err)
	}
	t.Log(string(b))

}