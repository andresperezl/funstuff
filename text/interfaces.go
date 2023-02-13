package text

type RuneWriter interface {
	WriteRune(r rune) (int, error)
}
