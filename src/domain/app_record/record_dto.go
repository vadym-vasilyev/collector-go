package app_record

type RecordsBatch struct {
	AppToken   string   `json:"app_token"`
	ClientType string   `json:"client_type" example:"iPhone10"`
	SessionId  string   `json:"session_id" example:"dyw234kjbb"`
	Records    []Record `json:"records"`
}

type Record struct {
	id           uint32       `json:"id" example:"1"`
	ScreenName   string       `json:"screen_name" example:"Main screen"`
	Timestamp    int64        `json:"timestamp" example:"1582759519"`
	ScreenLayout Element      `json:"screen_layout"`
	TouchEvents  []TouchEvent `json:"touch_events"`
	TransitionTo uint16       `json:"transition_to" example:"2"`
}

type Element struct {
	Id       uint32    `json:"id" example:"122323"`
	Name     string    `json:"name" example:"Self Destruct Button"`
	Type     string    `json:"type" example:"Button"`
	PosX     float32   `json:"pos_x" example:"10.0"`
	PosY     float32   `json:"pos_y"  example:"15.0"`
	Height   uint16    `json:"height"  example:"100"`
	Width    uint16    `json:"width"  example:"200"`
	Radius   uint16    `json:"radius"  example:"70"`
	Touched  bool      `json:"touched"  example:"true"`
	Children []Element `json:"children"`
}

type TouchEvent struct {
	PosX      float32 `json:"pos_x"  example:"10.0"`
	PosY      float32 `json:"pos_y"  example:"20.0"`
	TouchType uint8   `json:"touch_type"  example:"10"`
}
