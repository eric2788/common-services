package bilibili

type V1Resp struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
}

type XResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
}

// RoomInfoData copy from others
type RoomInfoData struct {
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

type NavInfoData struct {
	IsLogin bool `json:"isLogin"`
	WbiImg struct { 
		ImgUrl string `json:"img_url"`
		SubUrl string `json:"sub_url"`
	} 	`json:"wbi_img"`
}

type RoomInfo struct {
	V1Resp
	Data *RoomInfoData `json:"data"`
}

type UserInfo struct {
	XResp
	Data *UserInfoData `json:"data"`
}

type NavInfo struct {
	XResp
	Data *NavInfoData `json:"data"`
}

type UserInfoData struct {
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
		RoomStatus int	`json:"roomStatus"`
		LiveStatus int `json:"liveStatus"`
		Url string `json:"url"`
		Title string `json:"title"`
		Cover string `json:"cover"`
		RoomId int64 `json:"roomId"`
		RoundStatus int `json:"roundStatus"`
		BroadcastType int `json:"broadcast_type"`
	} `json:"live_room"`
}

type WebSocketInfo struct {
	V1Resp
	Data *WebSocketInfoData `json:"data"`

	LowLatencyHost string `json:"low_latency_host,omitempty"`
}

type HostServerInfo struct {
	ServerInfo
	WssPort int `json:"wss_port"`
	WsPort  int `json:"ws_port"`
}

type ServerInfo struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type WebSocketInfoData struct {
	RefreshRate int    `json:"refresh_rate"`
	MaxDelay    int    `json:"max_delay"`
	Port        int    `json:"port"`
	Host        string `json:"host"`

	HostServerList []HostServerInfo `json:"host_server_list"`

	ServerList []ServerInfo `json:"server_list"`
}