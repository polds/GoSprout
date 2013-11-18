package GoSprout

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const ()

func TestTagsFindAll(t *testing.T) {
	Convey("All attempts to Tags.FindAll should pass.", t, func() {
		Convey("When no inputs are used:", func() {
			tags := Tags(api)
			res, err := tags.FindAll()

			Convey("There should be no errors.", func() {
				So(err, ShouldBeNil)
			})

			Convey("There should be more than one result.", func() {
				So(res.Total, ShouldBeGreaterThan, 0)
			})

			Convey("'"+one+"' should be the first tag shown.", func() {
				So(res.Tags[0].Name, ShouldEqual, one)
			})
		})

		Convey("When 'order_by' is set:", func() {
			tags := Tags(api)

			Convey("When 'order_by' is 'created_at':", func() {
				tags.SetOrderBy("created_at")
				res, err := tags.FindAll()

				Convey("There should be no errors.", func() {
					So(err, ShouldBeNil)
				})

				Convey("'"+one+"' should be the first tag shown.", func() {
					So(res.Tags[0].Name, ShouldEqual, "Vault")
				})
			})

			Convey("When 'order_by' is 'updated_at':", func() {
				tags.SetOrderBy("updated_at")
				_, err := tags.FindAll()

				Convey("There should be no errors.", func() {
					So(err, ShouldBeNil)
				})
			})

			Convey("When 'order_by' is 'name':", func() {
				tags.SetOrderBy("name")
				res, err := tags.FindAll()

				Convey("There should be no errors.", func() {
					So(err, ShouldBeNil)
				})

				Convey("'"+two+"' should be the first tag shown.", func() {
					So(res.Tags[0].Name, ShouldEqual, two)
				})
			})

		})

		Convey("When 'order_dir' is set:", func() {
			tags := Tags(api)
			tags.SetOrderBy("name")

			Convey("When 'order_dir' is 'desc'", func() {
				tags.SetOrderDir("desc")
				res, err := tags.FindAll()

				Convey("There should be no errors.", func() {
					So(err, ShouldBeNil)
				})

				Convey("'"+three+"' should be the first tag shown.", func() {
					So(res.Tags[0].Name, ShouldEqual, three)
					So(res.Tags[0].Name, ShouldNotEqual, one)
				})
			})

			Convey("When 'order_dir' is 'asc'", func() {
				tags.SetOrderDir("asc")
				res, err := tags.FindAll()

				Convey("There should be no errors.", func() {
					So(err, ShouldBeNil)
				})

				Convey("'"+two+"' should be the first tag shown.", func() {
					So(res.Tags[0].Name, ShouldEqual, two)
					So(res.Tags[0].Name, ShouldNotEqual, one)
				})
			})
		})
	})
}

func TestTagsFindOne(t *testing.T) {
	Convey("All attempts to Tags.FindOne should pass.", t, func() {
		tags := Tags(api)

		Convey("When a valid tag is used:", func() {
			res, err := tags.FindOne(four)

			Convey("There should be no errors.", func() {
				So(err, ShouldBeNil)
			})

			Convey("There should be one or more videos.", func() {
				So(len(res.Videos), ShouldBeGreaterThan, 0)
				So(len(res.Videos), ShouldNotEqual, 0)
			})
		})

		Convey("When an invalid tag is used:", func() {
			res, err := tags.FindOne("abc123")

			Convey("There should be no http errors.", func() {
				So(err, ShouldBeNil)
			})

			Convey("Error should be set in the json object.", func() {
				So(res.Error, ShouldHaveSameTypeAs, "abc")
				So(res.Error, ShouldEqual, "Access Denied")
			})
		})

	})
}
