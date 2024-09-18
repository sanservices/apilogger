package apilogger

type StatusCat struct {
	Type string
}

var (
	StatusCatPassed  = StatusCat{Type: "Passed"}
	StatusCatPending = StatusCat{Type: "Pending"}
	StatusCatFailed  = StatusCat{Type: "Failed"}
)
