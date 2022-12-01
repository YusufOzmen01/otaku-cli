package constants

import (
	"github.com/YusufOzmen01/otaku-cli/constants/styles"
)

const ApiUrl = "https://gogoanime.consumet.org"

type ErrMsg struct {
	Err error
}

type ResultMsg struct {
	Data []*styles.AnimeResult
}

type DetailMsg struct {
	Data *styles.AnimeDetails
}

type WaitMsg struct {
}

type StreamResultData struct {
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
