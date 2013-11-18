package GoSprout

// import (
// 	"fmt"
// 	"net/url"
// 	"strconv"
// 	"strings"
// )

// type playlists struct {
// 	uri    string
// 	values url.Values
// }

// func Playlists() playlists {
// 	return playlists{
// 		uri:    "playlists",
// 		values: url.Values{},
// 	}
// }

// func (t playlists) SetOrderBy(s string) {
// 	t.values.Set("order_by", s)
// }

// func (t playlists) OrderBy() string {
// 	return t.values.Get("order_by")
// }

// func (t playlists) SetOrderDir(s string) {
// 	s = strings.ToLower(s)
// 	if s == "asc" || s == "desc" {
// 		t.values.Set("order_dir", s)
// 	}
// }

// func (t playlists) OrderDir() string {
// 	return t.values.Get("order_dir")
// }

// func (t playlists) SetPrivacy(i int) {
// 	t.values.Set("order_by", strconv.Itoa(i))
// }

// func (t playlists) Privacy() int {
// 	p := t.values.Get("order_by")
// 	return strconv.Atoi(p)
// }

// func (t playlists) FindOne(id string) string {
// 	fmt.Printf("%s%s/%s\n", endpoint, t.uri, id)

// 	do()

// 	return t.values.Encode()
// }
