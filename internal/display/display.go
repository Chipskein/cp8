package display

const Width = 64
const Height = 32

type Screen struct {
	Pixels [Width * Height]bool
}

func (s *Screen) ClearScreen() {
	s.Pixels = [Width * Height]bool{}
}
