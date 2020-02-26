package app_record

type RecordsBatch struct {
	AppToken   string   `json:"app_token"`
	ClientType string   `json:"client_type"`
	SessionId  string   `json:"session_id"`
	Records    []Record `json:"records"`
}

type Record struct {
	id           uint16       `json:"id"`
	ScreenName   string       `json:"screen_name"`
	Timestamp    int64        `json:"timestamp"`
	ScreenLayout Element      `json:"screen_layout"`
	TouchEvents  []TouchEvent `json:"touch_events"`
	TransitionTo uint16       `json:"transition_to"`
}

type Element struct {
	Id      uint16  `json:"id"`
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	PosX    float32 `json:"pos_x"`
	PosY    float32 `json:"pos_y"`
	Height  uint8   `json:"height"`
	Width   uint8   `json:"width"`
	Radius  uint8   `json:"radius"`
	Touched bool    `json:"touched"`
}

type TouchEvent struct {
	PosX      float32 `json:"pos_x"`
	PosY      float32 `json:"pos_y"`
	TouchType uint8   `json:"touch_type"`
}
