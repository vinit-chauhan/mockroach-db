package types

type Status int

const (
	Backlog Status = iota
	Waiting
	Todo
	InDesign
	InProgress
	Done
)

func (s Status) String() string {
	return [...]string{
		"Backlog",
		"Waiting",
		"To-Do",
		"In Design",
		"In Progress",
		"Done",
	}[s]
}

// Following functions are there to implement kancli.Status interface
// - ( Next, Prev, Int )
func (s Status) Next() int {
	if s == Done {
		return int(Todo)
	}

	return int(s + 1)
}

func (s Status) Prev() int {
	if s == Backlog {
		return int(Done)
	}

	return int(s - 1)
}

func (s Status) Int() int {
	return int(s)
}
