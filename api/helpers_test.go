package radosAPI
/*
import (
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewAPI(t *testing.T) {
	Convey("Testing New API", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		So(api, ShouldNotBeNil)
	})
}

func TestNewAPIWithPrefix(t *testing.T) {
	Convey("Testing New API with prefix", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"), "admin")

		So(api, ShouldNotBeNil)
	})
}

func TestCreateUser(t *testing.T) {
	Convey("Testing Create user", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"), "admin")

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")
	})
}

func TestCreateUserWithoutUID(t *testing.T) {
	Convey("Testing Create user without UID", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"), "admin")

		user, err := api.CreateUser(UserConfig{
			DisplayName: "Unit Test",
		})
		So(err, ShouldNotBeNil)
		So(user, ShouldBeNil)
	})
}

func TestCreateUserWithoutName(t *testing.T) {
	Convey("Testing Create user without name", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"), "admin")

		user, err := api.CreateUser(UserConfig{
			UID: "UnitTest",
		})
		So(err, ShouldNotBeNil)
		So(user, ShouldBeNil)
	})
}

func TestGetUser(t *testing.T) {
	Convey("Testing Get user", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		user, err := api.GetUser("UnitTest")
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
	})
}

func TestUpdateUser(t *testing.T) {
	Convey("Testing Update user", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		user, err := api.UpdateUser(UserConfig{
			UID:   "UnitTest",
			Email: "UnitTest@test.com",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.Email, ShouldEqual, "UnitTest@test.com")
	})
}

func TestUpdateUserWithoutUID(t *testing.T) {
	Convey("Testing Update user without UID", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		user, err := api.UpdateUser(UserConfig{
			Email: "UnitTest@test.com",
		})
		So(err, ShouldNotBeNil)
		So(user, ShouldBeNil)
	})
}

func TestRemoveUser(t *testing.T) {
	Convey("Testing Remove user", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		err := api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
		user, err := api.GetUser("UnitTest")
		So(err, ShouldNotBeNil)
		So(user, ShouldBeNil)
	})
}

func TestRemoveUserWithoutUID(t *testing.T) {
	Convey("Testing Remove user without UID", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		err := api.RemoveUser(UserConfig{
			PurgeData: true,
		})
		So(err, ShouldNotBeNil)
	})
}

func TestGetUsageEmpty(t *testing.T) {
	Convey("Testing Get Usage with empty struct", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		usage, err := api.GetUsage(UsageConfig{})
		So(err, ShouldBeNil)
		So(usage, ShouldNotBeNil)
		So(usage.Entries, ShouldBeNil)
		So(usage.Summary, ShouldBeNil)
	})
}

func TestGetUsageSummary(t *testing.T) {
	Convey("Testing Get Usage summary field", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		usage, err := api.GetUsage(UsageConfig{
			ShowSummary: true,
		})
		So(err, ShouldBeNil)
		So(usage, ShouldNotBeNil)
		So(usage.Entries, ShouldBeNil)
		So(usage.Summary, ShouldNotBeNil)
	})
}

func TestGetUsageEntries(t *testing.T) {
	Convey("Testing Get Usage entries field", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		usage, err := api.GetUsage(UsageConfig{
			ShowEntries: true,
		})
		So(err, ShouldBeNil)
		So(usage, ShouldNotBeNil)
		So(usage.Entries, ShouldNotBeNil)
		So(usage.Summary, ShouldBeNil)
	})
}

func TestGetUsageEntriesSummary(t *testing.T) {
	Convey("Testing Get Usage entries/summary field", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		usage, err := api.GetUsage(UsageConfig{
			ShowEntries: true,
			ShowSummary: true,
		})
		So(err, ShouldBeNil)
		So(usage, ShouldNotBeNil)
		So(usage.Entries, ShouldNotBeNil)
		So(usage.Summary, ShouldNotBeNil)
	})
}

func TestGetUsageEntriesSummaryWithUID(t *testing.T) {
	Convey("Testing Get Usage entries/summary field with specified uid", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		usage, err := api.GetUsage(UsageConfig{
			UID:         "UnitTest",
			ShowEntries: true,
			ShowSummary: true,
		})
		So(err, ShouldBeNil)
		So(usage, ShouldNotBeNil)
		So(usage.Entries, ShouldNotBeNil)
		So(usage.Summary, ShouldNotBeNil)
	})
}

func TestGetUsageEntriesSummaryWithUIDStartTime(t *testing.T) {
	Convey("Testing Get Usage entries/summary field with specified uid, and start Time", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		now := time.Now().AddDate(0, 0, -1)
		usage, err := api.GetUsage(UsageConfig{
			UID:         "UnitTest",
			ShowEntries: true,
			ShowSummary: true,
			Start:       &now,
		})
		So(err, ShouldBeNil)
		So(usage, ShouldNotBeNil)
		So(usage.Entries, ShouldNotBeNil)
		So(usage.Summary, ShouldNotBeNil)
	})
}

func TestDeleteAllUsages(t *testing.T) {
	Convey("Testing Delete all usages", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		err := api.DeleteUsage(UsageConfig{
			UID:       "UnitTest",
			RemoveAll: true,
		})
		So(err, ShouldBeNil)
	})
}

func TestCreateKey(t *testing.T) {
	Convey("Testing Create Key", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		keys, err := api.CreateKey(KeyConfig{
			UID: "UnitTest",
		})
		So(err, ShouldBeNil)
		So(keys, ShouldNotBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})
}

func TestCreateKeyWithoutUID(t *testing.T) {
	Convey("Testing Create Key without UID", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		keys, err := api.CreateKey(KeyConfig{})
		So(err, ShouldNotBeNil)
		So(keys, ShouldBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})
}

func TestRemoveKey(t *testing.T) {
	Convey("Testing Remove Key", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		keys, err := api.CreateKey(KeyConfig{
			UID: "UnitTest",
		})
		So(err, ShouldBeNil)
		So(keys, ShouldNotBeNil)

		err = api.RemoveKey(KeyConfig{
			AccessKey: (*keys)[0].AccessKey,
		})

		So(err, ShouldBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})
}

func TestRemoveKeyWithoutAccessKey(t *testing.T) {
	Convey("Testing Remove Key without access key", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		keys, err := api.CreateKey(KeyConfig{
			UID: "UnitTest",
		})
		So(err, ShouldBeNil)
		So(keys, ShouldNotBeNil)

		err = api.RemoveKey(KeyConfig{})

		So(err, ShouldNotBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})
}

func TestCreateSubUser(t *testing.T) {
	Convey("Testing CreateSubUser", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		sub, err := api.CreateSubUser(SubUserConfig{
			UID:     "UnitTest",
			SubUser: "SubUnitTest",
		})
		So(err, ShouldBeNil)
		So(sub, ShouldNotBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})
}

func TestCreateSubUserWithoutUID(t *testing.T) {
	Convey("Testing CreateSubUser without UID", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		sub, err := api.CreateSubUser(SubUserConfig{
			SubUser: "SubUnitTest",
		})
		So(err, ShouldNotBeNil)
		So(sub, ShouldBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})
}

func TestUpdateSubUser(t *testing.T) {
	Convey("Testing UpdateSubUser", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		sub, err := api.CreateSubUser(SubUserConfig{
			UID:     "UnitTest",
			SubUser: "SubUnitTest",
		})
		So(err, ShouldBeNil)
		So(sub, ShouldNotBeNil)

		sub, err = api.UpdateSubUser(SubUserConfig{
			UID:     "UnitTest",
			SubUser: (*sub)[0].ID,
		})
		So(err, ShouldBeNil)
		So(sub, ShouldNotBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})
}

func TestUpdateSubUserWithoutUID(t *testing.T) {
	Convey("Testing UpdateSubUser without UID", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		sub, err := api.CreateSubUser(SubUserConfig{
			UID:     "UnitTest",
			SubUser: "SubUnitTest",
		})
		So(err, ShouldBeNil)
		So(sub, ShouldNotBeNil)

		sub, err = api.UpdateSubUser(SubUserConfig{
			SubUser: (*sub)[0].ID,
		})
		So(err, ShouldNotBeNil)
		So(sub, ShouldBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})
}

func TestUpdateSubUserWithoutSubUID(t *testing.T) {
	Convey("Testing UpdateSubUser without SubUID", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		sub, err := api.CreateSubUser(SubUserConfig{
			UID:     "UnitTest",
			SubUser: "SubUnitTest",
		})
		So(err, ShouldBeNil)
		So(sub, ShouldNotBeNil)

		sub, err = api.UpdateSubUser(SubUserConfig{
			UID: "UnitTest",
		})
		So(err, ShouldNotBeNil)
		So(sub, ShouldBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})
}

func TestRemoveSubUser(t *testing.T) {
	Convey("Testing RemoveSubUser", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		sub, err := api.CreateSubUser(SubUserConfig{
			UID:     "UnitTest",
			SubUser: "SubUnitTest",
		})
		So(err, ShouldBeNil)
		So(sub, ShouldNotBeNil)

		err = api.RemoveSubUser(SubUserConfig{
			UID:     "UnitTest",
			SubUser: (*sub)[0].ID,
		})
		So(err, ShouldBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})
}

func TestRemoveSubUserWithoutUID(t *testing.T) {
	Convey("Testing RemoveSubUser without UID", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		sub, err := api.CreateSubUser(SubUserConfig{
			UID:     "UnitTest",
			SubUser: "SubUnitTest",
		})
		So(err, ShouldBeNil)
		So(sub, ShouldNotBeNil)

		err = api.RemoveSubUser(SubUserConfig{
			SubUser: (*sub)[0].ID,
		})
		So(err, ShouldNotBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})
}

func TestRemoveSubUserWithoutSubUID(t *testing.T) {
	Convey("Testing RemoveSubUser without SubUID", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		sub, err := api.CreateSubUser(SubUserConfig{
			UID:     "UnitTest",
			SubUser: "SubUnitTest",
		})
		So(err, ShouldBeNil)
		So(sub, ShouldNotBeNil)

		err = api.RemoveSubUser(SubUserConfig{
			UID: "UnitTest",
		})
		So(err, ShouldNotBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})
}
*/
