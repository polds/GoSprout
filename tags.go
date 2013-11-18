package GoSprout

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type tags struct {
	uri    string
	values url.Values
	api    string
}

func Tags(api string) tags {
	return tags{
		uri:    "tags",
		values: url.Values{},
		api:    api,
	}
}

func (t tags) SetOrderBy(s string) {
	t.values.Set("order_by", s)
}

func (t tags) OrderBy() string {
	return t.values.Get("order_by")
}

func (t tags) SetOrderDir(s string) {
	s = strings.ToLower(s)
	if s == "asc" || s == "desc" {
		t.values.Set("order_dir", s)
	}
}

func (t tags) OrderDir() string {
	return t.values.Get("order_dir")
}

func (t tags) SetPerPage(i int) {
	t.values.Set("per_page", strconv.Itoa(i))
}

func (t tags) PerPage() int {
	i, _ := strconv.Atoi(t.values.Get("per_page"))
	return i
}

func (t tags) SetPage(i int) {
	t.values.Set("page", strconv.Itoa(i))
}

func (t tags) Page() int {
	i, _ := strconv.Atoi(t.values.Get("page"))
	return i
}

func (t tags) FindOne(id string) (TagSingle, error) {
	uri := fmt.Sprintf("%s/%s", t.uri, id)
	res, err := do("GET", uri, t.api, t.values.Encode())
	if err != nil {
		return TagSingle{}, err
	}
	var single TagSingle
	err = json.Unmarshal(res, &single)
	if err != nil {
		return TagSingle{}, err
	}
	return single, nil
}

func (t tags) FindAll() (TagMany, error) {
	res, err := do("GET", t.uri, t.api, t.values.Encode())
	if err != nil {
		return TagMany{}, err
	}
	var many TagMany
	err = json.Unmarshal(res, &many)
	if err != nil {
		return TagMany{}, err
	}
	return many, nil
}

func (t tags) Create(name string) {

}

func (t tags) Edit() {

}

type TagMany struct {
	Total    int          `json:"total"`
	Tags     []*TagSingle `json:"tags"`
	NextPage string       `json:"next_page,omitempty"`
	Error    string       `json:"error,omitempty"`
}

type TagSingle struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Videos    []string  `json:"videos"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Error     string    `json:"error,omitempty"`
}
