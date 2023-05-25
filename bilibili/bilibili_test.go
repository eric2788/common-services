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

func TestGetRoomInfoWrongly(t *testing.T) {
	
	info, err := GetRoomInfo(99999)
	if err != nil {
		if e, ok := err.(*APIError); ok {
			t.Skip(e)
		} else {
			t.Fatal("unexpected err: ", err)
		}
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


func TestGetUserInfoWrongly(t *testing.T) {
	
	info, err := GetUserInfo(1145141919810)
	if err != nil {
		if e, ok := err.(*APIError); ok {
			t.Skip(e)
		} else {
			t.Fatal("unexpected err: ", err)
		}
	}

	t.Logf("%+v", info.XResp)

	b, err := json.MarshalIndent(*info.Data, "", "  ")
	if err != nil {
		t.Skip(err)
	}
	t.Log(string(b))
}


func TestGetVideoInfo(t *testing.T) {

	info, err := GetVideoInfo("BV1Y7411h7gR")
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

func TestGetVideoInfoWrongly(t *testing.T) {
	
	info, err := GetVideoInfo("BV1Y7411h9gR")
	if err != nil {
		if e, ok := err.(*APIError); ok {
			t.Skip(e)
		} else {
			t.Fatal("unexpected err: ", err)
		}
	}

	t.Logf("%+v", info.XResp)

	b, err := json.MarshalIndent(*info.Data, "", "  ")
	if err != nil {
		t.Skip(err)
	}
	t.Log(string(b))
}

func TestError(t *testing.T) {
	
	var r XResp
	err := mar(t, &r)
	if err != nil {
		t.Skip(err)
	}
	t.Logf("%+v", r)

	err = mar(t, nil)
	if err != nil {
		t.Skip(err)
	}

}

func mar(t *testing.T, res interface{}) error {
	t.Logf("%+v (is nil: %v)", res, res == nil)
	return json.Unmarshal([]byte(`{"code": 1, "message": "test", "ttl": 1}`), res)
}