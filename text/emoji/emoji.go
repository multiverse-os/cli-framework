package emoji

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"regexp"
	"unicode"
)

const (
	ReplacePadding = " "
)

func CodeMap() map[string]string {
	return emojiCodeMap
}

var flagRegexp = regexp.MustCompile(":flag-([a-z]{2}):")

func emojize(x string) string {
	str, ok := emojiCodeMap[x]
	if ok {
		return str + ReplacePadding
	}
	if match := flagRegexp.FindStringSubmatch(x); len(match) == 2 {
		return regionalIndicator(match[1][0]) + regionalIndicator(match[1][1])
	}
	return x
}

func regionalIndicator(i byte) string {
	return string('\U0001F1E6' + rune(i) - 'a')
}

func replaseEmoji(input *bytes.Buffer) string {
	emoji := bytes.NewBufferString(":")
	for {
		i, _, err := input.ReadRune()
		if err != nil {
			// not replase
			return emoji.String()
		}

		if i == ':' && emoji.Len() == 1 {
			return emoji.String() + replaseEmoji(input)
		}

		emoji.WriteRune(i)
		switch {
		case unicode.IsSpace(i):
			return emoji.String()
		case i == ':':
			return emojize(emoji.String())
		}
	}
}

func compile(x string) string {
	if x == "" {
		return ""
	}

	input := bytes.NewBufferString(x)
	output := bytes.NewBufferString("")

	for {
		i, _, err := input.ReadRune()
		if err != nil {
			break
		}
		switch i {
		default:
			output.WriteRune(i)
		case ':':
			output.WriteString(replaseEmoji(input))
		}
	}
	return output.String()
}

func Print(a ...interface{}) (int, error)   { return fmt.Print(compile(fmt.Sprint(a...))) }
func Println(a ...interface{}) (int, error) { return fmt.Println(compile(fmt.Sprint(a...))) }
func Printf(format string, a ...interface{}) (int, error) {
	return fmt.Printf(compile(fmt.Sprintf(format, a...)))
}

func Fprint(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprint(w, compile(fmt.Sprint(a...)))
}

func Fprintln(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprintln(w, compile(fmt.Sprint(a...)))
}
func Fprintf(w io.Writer, format string, a ...interface{}) (int, error) {
	return fmt.Fprint(w, compile(fmt.Sprintf(format, a...)))
}
func Sprint(a ...interface{}) string                 { return compile(fmt.Sprint(a...)) }
func Sprintf(format string, a ...interface{}) string { return compile(fmt.Sprintf(format, a...)) }
func Errorf(format string, a ...interface{}) error   { return errors.New(compile(Sprintf(format, a...))) }
