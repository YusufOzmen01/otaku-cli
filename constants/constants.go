package constants

type ErrMsg struct {
	Err error
}

type WaitMsg struct {
}

type StreamingUrlsMsg struct {
	Data *StreamData
}

type StreamData struct {
	Id       string `json:"id"`
	Url      string `json:"url"`
	Referer  string `json:"referer"`
	Priority int    `json:"priority"`
	Browser  bool   `json:"browser"`
	Website  string `json:"website"`
}
