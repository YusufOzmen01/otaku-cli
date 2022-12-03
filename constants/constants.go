package constants

type ErrMsg struct {
	Err error
}

type WaitMsg struct {
}

type StreamingUrlsMsg struct {
	Data *StreamData
}

type Stream struct {
	File  string `json:"file"`
	Label string `json:"label"`
	Type  string `json:"type"`
}

type StreamData struct {
	Referer string `json:"Referer"`
	Sources []struct {
		File  string `json:"file"`
		Label string `json:"label"`
		Type  string `json:"type"`
	} `json:"sources"`
	SourcesBk []struct {
		File  string `json:"file"`
		Label string `json:"label"`
		Type  string `json:"type"`
	} `json:"sources_bk"`
}
