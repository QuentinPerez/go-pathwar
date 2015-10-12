package api

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetUsers(t *testing.T) {
	Convey("Testing GetUsers", t, func() {
		Convey("without where clause", func() {
			client := NewAPIPathwar(os.Getenv("PATHWAR_TOKEN"), os.Getenv("PATHWAR_DEBUG"))
			users, err := client.GetUsers(nil)
			So(err, ShouldBeNil)
			So(len(users.Items) > 0, ShouldBeTrue)
			So(users.Items[0].Id, ShouldNotBeNil)
		})
		Convey("with where clause", func() {
			client := NewAPIPathwar(os.Getenv("PATHWAR_TOKEN"), os.Getenv("PATHWAR_DEBUG"))
			users, err := client.GetUsers(map[string]string{"login": "moul"})
			So(err, ShouldBeNil)
			So(len(users.Items), ShouldEqual, 1)
			So(users.Items[0].Login, ShouldEqual, "moul")
		})
	})
}

func TestGetRawOrganizationUsers(t *testing.T) {
	Convey("Testing GetRawOrganizationUsers", t, func() {
		Convey("without where clause", func() {
			client := NewAPIPathwar(os.Getenv("PATHWAR_TOKEN"), os.Getenv("PATHWAR_DEBUG"))
			rawOrganizationUsers, err := client.GetRawOrganizationUsers(nil)
			So(err, ShouldBeNil)
			So(len(rawOrganizationUsers.Items), ShouldNotEqual, 0)
			So(rawOrganizationUsers.Items[0].Id, ShouldNotBeNil)
		})
		Convey("with where clause", func() {
			client := NewAPIPathwar(os.Getenv("PATHWAR_TOKEN"), os.Getenv("PATHWAR_DEBUG"))
			rawOrganizationUsers, err := client.GetRawOrganizationUsers(map[string]string{"role": "pwner"})
			So(err, ShouldBeNil)
			So(len(rawOrganizationUsers.Items), ShouldNotEqual, 0)
			So(rawOrganizationUsers.Items[0].Role, ShouldEqual, "pwner")
		})
	})
}
