package GoSprout

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestVideosFindOne(t *testing.T) {
	Convey("All attempts to Videos.FindOne should pass.", t, func() {
		videos := Videos(api)

		Convey("When a valid video id is used:", func() {
			res, err := videos.FindOne(five)

			Convey("There should be no errors.", func() {
				So(err, ShouldBeNil)
			})

			Convey("Video should be called '"+six+"'.", func() {
				So(res.Title, ShouldEqual, six)
			})

		})

		Convey("When an invalid tag is used:", func() {
			res, err := videos.FindOne("abc123")

			Convey("There should be no http errors.", func() {
				So(err, ShouldBeNil)
			})

			Convey("Error should be set in the json object.", func() {
				So(res.Error, ShouldHaveSameTypeAs, "abc")
				So(res.Error, ShouldEqual, "Video Not Found")
			})
		})

	})

}
