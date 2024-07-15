package types

type status int

const (
	backlog status = iota
	waiting
	todo
	inDesign
	inProgress
	done
)

func (s status) String() string {
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
func (s status) Next() int {
	if s == done {
		return int(todo)
	}

	return int(s + 1)
}

func (s status) Prev() int {
	if s == backlog {
		return int(done)
	}

	return int(s - 1)
}

func (s status) Int() int {
	return int(s)
}
