//go:build go1.18
// +build go1.18

package main

import (
	"net/http"

	g "github.com/alarbada/gomponents"
	c "github.com/alarbada/gomponents/components"
	. "github.com/alarbada/gomponents/html"
)

func main() {
	_ = http.ListenAndServe("localhost:8080", http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	_ = Page(props{
		title: r.URL.Path,
		path:  r.URL.Path,
	}).Render(w)
}

type props struct {
	title string
	path  string
}

// Page is a whole document to output.
func Page(p props) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    p.title,
		Language: "en",
		Head: []g.Node{
			StyleEl(Type("text/css"),
				Raw("html { font-family: sans-serif; }"),
				Raw("ul { list-style-type: none; margin: 0; padding: 0; overflow: hidden; }"),
				Raw("ul li { display: block; padding: 8px; float: left; }"),
				Raw(".is-active { font-weight: bold; }"),
			),
		},
		Body: []g.Node{
			Navbar(p.path, []PageLink{
				{Path: "/foo", Name: "Foo"},
				{Path: "/bar", Name: "Bar"},
			}),
			H1(Text(p.title)),
			P(Textf("Welcome to the page at %v.", p.path)),
		},
	})
}

type PageLink struct {
	Path string
	Name string
}

func Navbar(currentPath string, links []PageLink) g.Node {
	return Div(
		Ul(
			NavbarLink("/", "Home", currentPath),

			Foreach(links, func(pl PageLink) g.Node {
				return NavbarLink(pl.Path, pl.Name, currentPath)
			}),
		),

		Hr(),
	)
}

func NavbarLink(href, name, currentPath string) g.Node {
	return Li(A(Href(href), c.Classes{"is-active": currentPath == href}, Text(name)))
}
