package entities

// Theater anchors physical locations natively tracking available show schedules.
type Theater struct {
	ID       string
	Name     string
	Location string
	Shows    []*Show
}
