package glitch

import (
	"bytes"
	"math/rand"
	"strings"

	"github.com/andresperezl/funstuff/text"
)

var (
	diacriticsTop    []rune
	diacriticsMiddle []rune
	diacriticsBottom []rune
)

type options struct {
	top           bool
	middle        bool
	bottom        bool
	maxHeight     int     // How many diacritic marks shall we put on top/bottom?
	randomization float32 // maxHeight 100 and randomization 0.2: the height goes from 80 to 100. randomization 0.7, height goes from 30 to 100.
}

type optionHint int

const (
	optTop optionHint = iota
	optMiddle
	optBottom
	optMaxHeight
	optRandomization
)

type Option struct {
	hint  optionHint
	value any
}

// WithTop should the encoder put diacritics on top of the characters
func WithTop(v bool) Option {
	return Option{optTop, v}
}

// WithTop should the encoder put diacritics on the middle  of the characters
func WithMiddle(v bool) Option {
	return Option{optMiddle, v}
}

// WithTop should the encoder put diacritics on the bottom of the characters
func WithBottom(v bool) Option {
	return Option{optBottom, v}
}

// WithMaxHeight How many diacritic marks shall we put on top/bottom?
func WithMaxHeight(v int) Option {
	if v < 0 {
		panic("maxHeight cannot be less than 0")
	}
	return Option{optMaxHeight, v}
}

// WithRandomization 0-100%. maxHeight 100 and randomization 20%: the height goes from 80 to 100. randomization 70%, height goes from 30 to 100.
func WithRandomization(v int) Option {
	if v < 0 || v > 100 {
		panic("randomization needs to be between 0 and 100")
	}
	return Option{optRandomization, v}
}

var defaultOptions = options{
	top:           true,
	middle:        true,
	bottom:        true,
	maxHeight:     15,
	randomization: 1.0,
}

func init() {
	for i := 768; i <= 789; i++ {
		diacriticsTop = append(diacriticsTop, rune(i))
	}

	for i := 790; i <= 819; i++ {
		if i != 794 && i != 795 {
			diacriticsBottom = append(diacriticsBottom, rune(i))
		}
	}
	diacriticsTop = append(diacriticsTop, rune(794))
	diacriticsTop = append(diacriticsTop, rune(795))

	for i := 820; i <= 824; i++ {
		diacriticsMiddle = append(diacriticsMiddle, rune(i))
	}

	for i := 825; i <= 828; i++ {
		diacriticsBottom = append(diacriticsBottom, rune(i))
	}

	for i := 829; i <= 836; i++ {
		diacriticsTop = append(diacriticsTop, rune(i))
	}
	diacriticsTop = append(diacriticsTop, rune(836))
	diacriticsBottom = append(diacriticsBottom, rune(837))
	diacriticsTop = append(diacriticsTop, rune(838))
	diacriticsBottom = append(diacriticsBottom, rune(839))
	diacriticsBottom = append(diacriticsBottom, rune(840))
	diacriticsBottom = append(diacriticsBottom, rune(841))
	diacriticsTop = append(diacriticsTop, rune(842))
	diacriticsTop = append(diacriticsTop, rune(843))
	diacriticsTop = append(diacriticsTop, rune(844))
	diacriticsBottom = append(diacriticsBottom, rune(845))
	diacriticsBottom = append(diacriticsBottom, rune(846))
	// 847 (U+034F) is invisible http://en.wikipedia.org/wiki/Combining_grapheme_joiner
	diacriticsTop = append(diacriticsTop, rune(848))
	diacriticsTop = append(diacriticsTop, rune(849))
	diacriticsTop = append(diacriticsTop, rune(850))
	diacriticsBottom = append(diacriticsBottom, rune(851))
	diacriticsBottom = append(diacriticsBottom, rune(852))
	diacriticsBottom = append(diacriticsBottom, rune(853))
	diacriticsBottom = append(diacriticsBottom, rune(854))
	diacriticsTop = append(diacriticsTop, rune(855))
	diacriticsTop = append(diacriticsTop, rune(856))
	diacriticsBottom = append(diacriticsBottom, rune(857))
	diacriticsBottom = append(diacriticsBottom, rune(858))
	diacriticsTop = append(diacriticsTop, rune(859))
	diacriticsBottom = append(diacriticsBottom, rune(860))
	diacriticsTop = append(diacriticsTop, rune(861))
	diacriticsTop = append(diacriticsTop, rune(861))
	diacriticsBottom = append(diacriticsBottom, rune(863))
	diacriticsTop = append(diacriticsTop, rune(864))
	diacriticsTop = append(diacriticsTop, rune(865))
}

// Encode transforms a message into glitch text, with default options top=true,
// middle=true, bottom=true, maxHeight=15, randomization=100, unless modified by
// opts
func Encode(t []byte, opts ...Option) []byte {
	encOpts := parseOptions()

	newText := &bytes.Buffer{}
	reader := bytes.NewReader(t)
	for newChar, _, err := reader.ReadRune(); err == nil; newChar, _, err = reader.ReadRune() {
		newText.WriteRune(newChar)
		// Middle
		// Put just one of the middle characters there, or it gets crowded
		if encOpts.middle {
			newText.WriteRune(diacriticsMiddle[rand.Intn(len(diacriticsMiddle))])
		}

		if encOpts.top {
			addRunesFromList(newText, diacriticsTop, encOpts)
		}

		if encOpts.bottom {
			addRunesFromList(newText, diacriticsBottom, encOpts)
		}
	}
	return newText.Bytes()
}

// EncodeString transforms a message into glitch text, with default options top=true,
// middle=true, bottom=true, maxHeight=15, randomization=100, unless modified by
// opts
func EncodeString(t string, opts ...Option) string {
	encOpts := parseOptions()

	newText := &strings.Builder{}
	for _, newChar := range t {
		newText.WriteRune(newChar)
		// Middle
		// Put just one of the middle characters there, or it gets crowded
		if encOpts.middle {
			newText.WriteRune(diacriticsMiddle[rand.Intn(len(diacriticsMiddle))])
		}

		if encOpts.top {
			addRunesFromList(newText, diacriticsTop, encOpts)
		}

		if encOpts.bottom {
			addRunesFromList(newText, diacriticsBottom, encOpts)
		}
	}
	return newText.String()
}

func addRunesFromList(dst text.RuneWriter, list []rune, opts options) {
	qty := opts.maxHeight - int(rand.Float32()*opts.randomization*float32(optMaxHeight))
	for count := 0; count < qty; count++ {
		dst.WriteRune(list[rand.Intn(len(list))])
	}
}

func parseOptions(opts ...Option) options {
	encOpts := defaultOptions
	for _, o := range opts {
		switch o.hint {
		case optTop:
			encOpts.top = o.value.(bool)
		case optMiddle:
			encOpts.middle = o.value.(bool)
		case optBottom:
			encOpts.bottom = o.value.(bool)
		case optMaxHeight:
			encOpts.maxHeight = o.value.(int)
		case optRandomization:
			encOpts.randomization = o.value.(float32)
		}
	}
	return encOpts
}
