package enum

type Machine_state int
const (
	Running Machine_state = 0
	Paused  Machine_state = 1
	Debug   Machine_state = 2
	Stop    Machine_state = 3
)
