package html_test

import (
	"errors"
	"fmt"
	"io"
	"testing"

	g "github.com/alarbada/gomponents"
	. "github.com/alarbada/gomponents/html"
	"github.com/alarbada/gomponents/internal/assert"
)

type erroringWriter struct{}

func (w *erroringWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("don't want to write")
}

func TestDoctype(t *testing.T) {
	t.Run("returns doctype and children", func(t *testing.T) {
		assert.Equal(t, `<!doctype html><html></html>`, Doctype(g.El("html")))
	})

	t.Run("errors on write error in Render", func(t *testing.T) {
		err := Doctype(g.El("html")).Render(&erroringWriter{})
		assert.Error(t, err)
	})
}

func TestSimpleElements(t *testing.T) {
	cases := map[string]func(...g.Node) g.Node{
		"a":          A,
		"abbr":       Abbr,
		"address":    Address,
		"article":    Article,
		"aside":      Aside,
		"audio":      Audio,
		"b":          B,
		"blockquote": BlockQuote,
		"body":       Body,
		"button":     Button,
		"canvas":     Canvas,
		"caption":    Caption,
		"cite":       Cite,
		"code":       Code,
		"colgroup":   ColGroup,
		"data":       DataEl,
		"datalist":   DataList,
		"dd":         Dd,
		"del":        Del,
		"details":    Details,
		"dfn":        Dfn,
		"dialog":     Dialog,
		"div":        Div,
		"dl":         Dl,
		"dt":         Dt,
		"em":         Em,
		"fieldset":   FieldSet,
		"figcaption": FigCaption,
		"figure":     Figure,
		"footer":     Footer,
		"form":       FormEl,
		"h1":         H1,
		"h2":         H2,
		"h3":         H3,
		"h4":         H4,
		"h5":         H5,
		"h6":         H6,
		"head":       Head,
		"header":     Header,
		"hgroup":     HGroup,
		"html":       HTML,
		"i":          I,
		"iframe":     IFrame,
		"ins":        Ins,
		"kbd":        Kbd,
		"label":      Label,
		"legend":     Legend,
		"li":         Li,
		"main":       Main,
		"mark":       Mark,
		"menu":       Menu,
		"meter":      Meter,
		"nav":        Nav,
		"noscript":   NoScript,
		"object":     Object,
		"ol":         Ol,
		"optgroup":   OptGroup,
		"option":     Option,
		"p":          P,
		"picture":    Picture,
		"pre":        Pre,
		"progress":   Progress,
		"q":          Q,
		"s":          S,
		"samp":       Samp,
		"script":     Script,
		"section":    Section,
		"select":     Select,
		"small":      Small,
		"span":       Span,
		"strong":     Strong,
		"style":      StyleEl,
		"sub":        Sub,
		"summary":    Summary,
		"sup":        Sup,
		"svg":        SVG,
		"table":      Table,
		"tbody":      TBody,
		"td":         Td,
		"textarea":   Textarea,
		"tfoot":      TFoot,
		"th":         Th,
		"thead":      THead,
		"time":       Time,
		"title":      TitleEl,
		"tr":         Tr,
		"u":          U,
		"ul":         Ul,
		"var":        Var,
		"video":      Video,
	}

	for name, fn := range cases {
		t.Run(fmt.Sprintf("should output %v", name), func(t *testing.T) {
			n := fn(g.Attr("id", "hat"))
			assert.Equal(t, fmt.Sprintf(`<%v id="hat"></%v>`, name, name), n)
		})
	}
}

func TestSimpleVoidKindElements(t *testing.T) {
	cases := map[string]func(...g.Node) g.Node{
		"area":   Area,
		"base":   Base,
		"br":     Br,
		"col":    Col,
		"embed":  Embed,
		"hr":     Hr,
		"img":    Img,
		"input":  Input,
		"link":   Link,
		"meta":   Meta,
		"param":  Param,
		"source": Source,
		"wbr":    Wbr,
	}

	for name, fn := range cases {
		t.Run(fmt.Sprintf("should output %v", name), func(t *testing.T) {
			n := fn(g.Attr("id", "hat"))
			assert.Equal(t, fmt.Sprintf(`<%v id="hat">`, name), n)
		})
	}
}

type CustomComponent struct {
	PropA string
	PropB string
}

func (c CustomComponent) Render(w io.Writer) error {
	return Div(
		g.Attr("prop-a", c.PropA),
		g.Attr("prop-b", c.PropB),
	).Render(w)
}

func (c CustomComponent) Children(children ...g.Node) g.Node {
	return Div(
		g.Attr("prop-a", c.PropA),
		g.Attr("prop-b", c.PropB),
		g.Group(children),
	)
}

type shoelaceDivider struct{ g.AsEl }

func (s shoelaceDivider) Render(w io.Writer) error {
	return g.El("sl-divider").Render(w)
}

func TestCustomComponent(t *testing.T) {
	n := CustomComponent{
		PropA: "a",
		PropB: "b",
	}
	t.Run("should output custom component", func(t *testing.T) {
		assert.Equal(t, `<div prop-a="a" prop-b="b"></div>`, n)
	})

	t.Run("should output custom component with custom tag", func(t *testing.T) {
		n := n.Children(shoelaceDivider{})
		assert.Equal(t, `<div prop-a="a" prop-b="b"><sl-divider></sl-divider></div>`, n)
	})
}
