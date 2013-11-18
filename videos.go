package GoSprout

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type videos struct {
	uri    string
	values url.Values
	api    string
}

func Videos(api string) videos {
	return videos{
		uri:    "videos",
		values: url.Values{},
		api:    api,
	}
}

func (t videos) SetOrderBy(s string) {
	t.values.Set("order_by", s)
}

func (t videos) OrderBy() string {
	return t.values.Get("order_by")
}

func (t videos) SetOrderDir(s string) {
	s = strings.ToLower(s)
	if s == "asc" || s == "desc" {
		t.values.Set("order_dir", s)
	}
}

func (t videos) OrderDir() string {
	return t.values.Get("order_dir")
}

func (t videos) SetPerPage(i int) {
	t.values.Set("per_page", strconv.Itoa(i))
}

func (t videos) PerPage() int {
	i, _ := strconv.Atoi(t.values.Get("per_page"))
	return i
}

func (t videos) SetPage(i int) {
	t.values.Set("page", strconv.Itoa(i))
}

func (t videos) Page() int {
	i, _ := strconv.Atoi(t.values.Get("page"))
	return i
}

func (t videos) FindOne(id string) (VideoSingle, error) {
	uri := fmt.Sprintf("%s/%s", t.uri, id)
	res, err := do("GET", uri, t.api, t.values.Encode())
	if err != nil {
		return VideoSingle{}, err
	}
	var single VideoSingle
	err = json.Unmarshal(res, &single)
	if err != nil {
		return VideoSingle{}, err
	}
	return single, nil
}

func (t videos) FindAll() (VideoMany, error) {
	res, err := do("GET", t.uri, t.api, t.values.Encode())
	if err != nil {
		return VideoMany{}, err
	}
	var many VideoMany
	err = json.Unmarshal(res, &many)
	if err != nil {
		return VideoMany{}, err
	}
	return many, nil
}

type VideoMany struct {
	Total    int            `json:"total"`
	Videos   []*VideoSingle `json:"videos,omitempty"`
	NextPage string         `json:"next_page,omitempty"`
	Error    string         `json:"error,omitempty"`
}

type VideoSingle struct {
	Id                   string    `json:"id"`
	Width                int       `json:"width"`
	Height               int       `json:"height"`
	EmbedCode            string    `json:"embed_code"`
	SourceVideoFileSize  int       `json:"source_video_file_size"`
	SdVideoFileSize      int       `json:"sd_video_file_size"`
	HdVideoFileSize      int       `json:"hd_video_file_size"`
	SecurityToken        string    `json:"security_token"`
	Title                string    `json:"title"`
	Description          string    `json:"description"`
	Duration             float32   `json:"duration"`
	Privacy              int       `json:"privacy"`
	Password             string    `json:"password"`
	State                string    `json:"state"`
	Tags                 []string  `json:"tags"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	Plays                int       `json:"plays"`
	Progress             int       `json:"progress"`
	RequiresSignedEmbeds bool      `json:"requires_signed_embeds"`
	Assets               Assets    `json:"assets"`
	Error                string    `json:"error,omitempty"`
}

type Assets struct {
	Videos       VideosInfo `json:"videos"`
	Thumbnails   []string   `json:"thumbnails"`
	PosterFrames []string   `json:"poster_frames"`
}

type VideosInfo struct {
	SdVideoUrl     string `json:"sd_video_url"`
	HdVideoUrl     string `json:"hd_video_url"`
	SourceVideoUrl string `json:"source_video_url"`
}
