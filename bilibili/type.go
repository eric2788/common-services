package bilibili

import "fmt"

type (
	V1Resp struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		Message string `json:"message"`
	}

	XResp struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		TTL     int    `json:"ttl"`
	}

	RoomInfoData struct {
		Uid              int64    `json:"uid"`
		RoomId           int64    `json:"room_id"`
		ShortId          int      `json:"short_id"`
		Attention        int64    `json:"attention"`
		Online           int      `json:"online"`
		IsPortrait       bool     `json:"is_portrait"`
		Description      string   `json:"description"`
		LiveStatus       int8     `json:"live_status"`
		AreaId           int      `json:"area_id"`
		ParentAreaId     int      `json:"parent_area_id"`
		ParentAreaName   string   `json:"parent_area_name"`
		OldAreaId        int      `json:"old_area_id"`
		Background       string   `json:"background"`
		Title            string   `json:"title"`
		UserCover        string   `json:"user_cover"`
		Keyframe         string   `json:"keyframe"`
		IsStrictRoom     bool     `json:"is_strict_room"`
		LiveTime         string   `json:"live_time"`
		Tags             string   `json:"tags"`
		IsAnchor         int      `json:"is_anchor"`
		RoomSilentType   string   `json:"room_silent_type"`
		RoomSilentLevel  int      `json:"room_silent_level"`
		RoomSilentSecond int      `json:"room_silent_second"`
		AreaName         string   `json:"area_name"`
		Pendants         string   `json:"pendants"`
		AreaPendants     string   `json:"area_pendants"`
		HotWords         []string `json:"hot_words"`
		HotWordsStatus   int      `json:"hot_words_status"`
		Verify           string   `json:"verify"`
		NewPendants      struct {
			Frame struct {
				Name       string `json:"name"`
				Value      string `json:"value"`
				Position   int    `json:"position"`
				Desc       string `json:"desc"`
				Area       int    `json:"area"`
				AreaOld    int    `json:"area_old"`
				BgColor    string `json:"bg_color"`
				BgPic      string `json:"bg_pic"`
				UseOldArea bool   `json:"use_old_area"`
			} `json:"frame"`
			Badge       interface{} `json:"badge"`
			MobileFrame struct {
				Name       string `json:"name"`
				Value      string `json:"value"`
				Position   int    `json:"position"`
				Desc       string `json:"desc"`
				Area       int    `json:"area"`
				AreaOld    int    `json:"area_old"`
				BgColor    string `json:"bg_color"`
				BgPic      string `json:"bg_pic"`
				UseOldArea bool   `json:"use_old_area"`
			} `json:"mobile_frame"`
			MobileBadge interface{} `json:"mobile_badge"`
		} `json:"new_pendants"`
		UpSession            string `json:"up_session"`
		PkStatus             int    `json:"pk_status"`
		PkId                 int    `json:"pk_id"`
		BattleId             int    `json:"battle_id"`
		AllowChangeAreaTime  int    `json:"allow_change_area_time"`
		AllowUploadCoverTime int    `json:"allow_upload_cover_time"`
		StudioInfo           struct {
			Status     int           `json:"status"`
			MasterList []interface{} `json:"master_list"`
		} `json:"studio_info"`
	}

	NavInfoData struct {
		IsLogin bool `json:"isLogin"`
		WbiImg  struct {
			ImgUrl string `json:"img_url"`
			SubUrl string `json:"sub_url"`
		} `json:"wbi_img"`
	}

	UserInfoData struct {
		Mid      int64  `json:"mid"`
		Name     string `json:"name"`
		Sex      string `json:"sex"`
		Face     string `json:"face"`
		Sign     string `json:"sign"`
		Rank     int    `json:"rank"`
		Level    int8   `json:"level"`
		Official *struct {
			Role  int    `json:"role"`
			Desc  string `json:"desc"`
			Title string `json:"title"`
			Type  int    `json:"type"`
		} `json:"official"`
		LiveRoom struct {
			RoomStatus    int    `json:"roomStatus"`
			LiveStatus    int    `json:"liveStatus"`
			Url           string `json:"url"`
			Title         string `json:"title"`
			Cover         string `json:"cover"`
			RoomId        int64  `json:"roomId"`
			RoundStatus   int    `json:"roundStatus"`
			BroadcastType int    `json:"broadcast_type"`
		} `json:"live_room"`
	}

	VideoInfoData struct {
		View struct {
			Bvid        string `json:"bvid"`
			Aid         int64  `json:"aid"`
			TName       string `json:"tname"`
			Title       string `json:"title"`
			Pic         string `json:"pic"`
			PublishDate int64  `json:"pubdate"`
			Ctime       int64  `json:"ctime"`
			Desc        string `json:"desc"`
			Duration    int64  `json:"416"`
			Owner       struct {
				Mid  int64  `json:"mid"`
				Name string `json:"name"`
				Face string `json:"face"`
			} `json:"owner"`
			Stats struct {
				View      int64 `json:"view"`
				Danmaku   int64 `json:"danmaku"`
				Reply     int64 `json:"reply"`
				Favourite int64 `json:"favorite"`
				Coin      int64 `json:"coin"`
				Share     int64 `json:"share"`
				Like      int64 `json:"like"`
				DisLike   int64 `json:"dislike"`
			} `json:"stat"`
			Cid   int64 `json:"cid"`
			Pages []struct {
				Part       string `json:"part"`
				FirstFrame string `json:"first_frame"`
				Cid        int64  `json:"cid"`
				Page       int    `json:"page"`
				Vid        string `json:"vid"`
				Weblink    string `json:"weblink"`
			} `json:"pages"`
		} `json:"View"`

		Tags []struct {
			TagId        int    `json:"tag_id"`
			TagName      string `json:"tag_name"`
			Cover        string `json:"cover"`
			HeadCover    string `json:"head_cover"`
			Content      string `json:"content"`
			ShortContent string `json:"short_content"`
			Type         int    `json:"type"`
			Color        string `json:"color"`
		} `json:"Tags"`

		Reply struct {
			Page struct {
				Account int `json:"account"`
				Count   int `json:"count"`
				Num     int `json:"num"`
				Size    int `json:"size"`
			} `json:"page"`

			Replies []struct {
				Type   int   `json:"type"`
				Mid    int64 `json:"mid"`
				RCount int   `json:"rcount"`
				Count  int   `json:"count"`
				Floor  int   `json:"floor"`
				CTime  int64 `json:"ctime"`
				Like   int   `json:"like"`

				Content struct {
					Message string `json:"message"`
					Plat    int    `json:"plat"`
					Devide  string `json:"device"`
				} `json:"content"`
			} `json:"replies"`
		} `json:"Reply"`
	}

	RoomInfo struct {
		V1Resp
		Data *RoomInfoData `json:"data"`
	}

	UserInfo struct {
		XResp
		Data *UserInfoData `json:"data"`
	}

	NavInfo struct {
		XResp
		Data *NavInfoData `json:"data"`
	}

	VideoInfo struct {
		XResp
		Data *VideoInfoData `json:"data"`
	}

	APIError struct {
		Code    int   
		Message string
		Data *interface{} 
	}
)


func (e *APIError) Error() string {
	return fmt.Sprintf("Bilibili API Error: [%d] %s", e.Code, e.Message)
}
