package html

import (
	"fmt"
	"io"
	"strconv"
	"text/template"

	g "github.com/alarbada/gomponents"
)

// Foreach renders a slice of anything to a single Node.
func Foreach[T any](s []T, cb func(T) g.Node) g.Node {
	return g.NodeFunc(func(w io.Writer) error {
		for _, val := range s {
			c := cb(val)
			if c == nil {
				continue
			}

			if err := c.Render(w); err != nil {
				return err
			}
		}
		return nil
	})
}

// Foreach renders a slice of anything to a single Node.
func ForeachI[T any](s []T, cb func(string, T) g.Node) g.Node {
	return g.NodeFunc(func(w io.Writer) error {
		for i, val := range s {
			iStr := strconv.Itoa(i)
			c := cb(iStr, val)
			if c == nil {
				continue
			}

			if err := c.Render(w); err != nil {
				return err
			}
		}
		return nil
	})
}

// LoopTimes renders a callback function n times.
func LoopTimes(times int, cb func(i int) g.Node) g.Node {
	return g.NodeFunc(func(w io.Writer) error {
		for i := 0; i < times; i++ {
			c := cb(i)
			if c == nil {
				continue
			}

			if err := c.Render(w); err != nil {
				return err
			}
		}
		return nil
	})
}


// Text creates a text DOM Node that Renders the escaped string t.
func Text(t string) g.Node {
	return g.NodeFunc(func(w io.Writer) error {
		bs := g.StringToBytes(template.HTMLEscapeString(t))
		_, err := w.Write(bs)
		return err
	})
}

// Textf creates a text DOM Node that Renders the interpolated and escaped string format.
func Textf(format string, a ...interface{}) g.Node {
	return g.NodeFunc(func(w io.Writer) error {
		s := template.HTMLEscapeString(fmt.Sprintf(format, a...))
		bs := g.StringToBytes(s)
		_, err := w.Write(bs)
		return err
	})
}

// Raw creates a text DOM Node that just Renders the unescaped string t.
func Raw(t string) g.Node {
	return g.NodeFunc(func(w io.Writer) error {
		bs := g.StringToBytes(t)
		_, err := w.Write(bs)
		return err
	})
}

// Rawf creates a text DOM Node that just Renders the interpolated and unescaped string format.
func Rawf(format string, a ...interface{}) g.Node {
	return g.NodeFunc(func(w io.Writer) error {
		bs := g.StringToBytes(fmt.Sprintf(format, a...))
		_, err := w.Write(bs)
		return err
	})
}

