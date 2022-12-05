// package userentry provides type Entry and its affiliates,
// an entry is record corresponds to infomation of a person
// to be submitted.
package userentry

type Entry struct {
	StudentId      string
	LocationString string
	Latitude       string
	Longitude      string
}

func NewEntry(id, loc, lat, lon string) (e *Entry) {
	e.StudentId = id
	e.LocationString = loc
	e.Latitude = lat
	e.Longitude = lon
}
