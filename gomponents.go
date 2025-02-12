// Package gomponents provides view components in Go, that render to HTML 5.
//
// The primary interface is a Node. It describes a function Render, which should render the Node
// to the given writer as a string.
//
// All DOM elements and attributes can be created by using the El and Attr functions.
// The functions Text, Textf, Raw, and Rawf can be used to create text nodes, either HTML-escaped or unescaped.
// See also helper functions Group, Map, and If for mapping data to Nodes and inserting them conditionally.
//
// For basic HTML elements and attributes, see the package html.
// For higher-level HTML components, see the package components.
// For SVG elements and attributes, see the package svg.
// For HTTP helpers, see the package http.
package gomponents

import (
	"html/template"
	"io"
	"strings"
	"unsafe"
)

// Node is a DOM node that can Render itself to a io.Writer.
type Node interface {
	Render(w io.Writer) error
}

type TypedNode interface {
	Node
	Type() NodeType
}

// NodeType describes what type of Node it is, currently either an ElementType or an AttributeType.
// This decides where a Node should be rendered.
// Nodes default to being ElementType.
type NodeType uint8

const (
	ElementType NodeType = iota
	AttributeType
)

// NodeFunc is a render function that is also a Node of ElementType.
type NodeFunc func(io.Writer) error

// Render satisfies Node.
func (n NodeFunc) Render(w io.Writer) error {
	return n(w)
}

func (n NodeFunc) Type() NodeType {
	return ElementType
}

// String satisfies fmt.Stringer.
func (n NodeFunc) String() string {
	var b strings.Builder
	_ = n.Render(&b)
	return b.String()
}

// El creates an element DOM Node with a name and child Nodes.
// See https://dev.w3.org/html5/spec-LC/syntax.html#elements-0 for how elements are rendered.
// No tags are ever omitted from normal tags, even though it's allowed for elements given at
// https://dev.w3.org/html5/spec-LC/syntax.html#optional-tags
// If an element is a void element, non-attribute children nodes are ignored.
// Use this if no convenience creator exists.
func El(name string, children ...Node) Node {
	return NodeFunc(func(w2 io.Writer) error {
		w := &statefulWriter{w: w2}

		w.WriteString("<")
		w.WriteString(name)

		classValues := []string{}
		for _, c := range children {
			renderAttributes(w, c, &classValues)
		}

		if l := len(classValues); l > 0 {
			w.WriteString(` class="`)
			for i, v := range classValues {
				w.WriteString(v)
				if i < l-1 {
					w.WriteString(" ")
				}
			}
			w.WriteString(`"`)
		}

		w.WriteString(">")

		if isVoidElement(name) {
			return w.err
		}

		for _, c := range children {
			renderChild(w, c)
		}

		w.WriteString("</")
		w.WriteString(name)
		w.WriteString(">")

		return w.err
	})
}

func renderAttributes(w *statefulWriter, n Node, classValues *[]string) {
	if w.err != nil || n == nil {
		return
	}

	if g, ok := n.(group); ok {
		for _, groupC := range g.children {
			renderAttributes(w, groupC, classValues)
		}
		return
	}

	if attr, ok := n.(*attr); ok && attr.name == "class" {
		*classValues = append(*classValues, *attr.value)
		return
	}

	if n, ok := n.(TypedNode); ok && n.Type() == AttributeType {
		w.err = n.Render(w.w)
	}
}

func renderChild(w *statefulWriter, n Node) {
	if w.err != nil || n == nil {
		return
	}

	if g, ok := n.(group); ok {
		for _, groupC := range g.children {
			renderChild(w, groupC)
		}
		return
	}

	typed, ok := n.(TypedNode)
	if !ok || typed.Type() == ElementType {
		w.err = n.Render(w.w)
		return
	}
}

// statefulWriter only writes if no errors have occurred earlier in its lifetime.
type statefulWriter struct {
	w   io.Writer
	err error
}

func StringToBytes(str string) []byte {
	if str == "" {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

func (w *statefulWriter) WriteString(s string) {
	if w.err != nil {
		return
	}

	_, w.err = w.w.Write(StringToBytes(s))
}

// voidElements don't have end tags and must be treated differently in the rendering.
// See https://dev.w3.org/html5/spec-LC/syntax.html#void-elements
var voidElements = map[string]struct{}{
	"area":    {},
	"base":    {},
	"br":      {},
	"col":     {},
	"command": {},
	"embed":   {},
	"hr":      {},
	"img":     {},
	"input":   {},
	"keygen":  {},
	"link":    {},
	"meta":    {},
	"param":   {},
	"source":  {},
	"track":   {},
	"wbr":     {},
}

func isVoidElement(name string) bool {
	_, ok := voidElements[name]
	return ok
}

// Attr creates an attribute DOM Node with a name and optional value.
// If only a name is passed, it's a name-only (boolean) attribute (like "required").
// If a name and value are passed, it's a name-value attribute (like `class="header"`).
// More than one value make Attr panic.
// Use this if no convenience creator exists.
func Attr(name string, value ...string) Node {
	switch len(value) {
	case 0:
		return &attr{name: name}
	case 1:
		return &attr{name: name, value: &value[0]}
	default:
		panic("attribute must be just name or name and value pair")
	}
}

type attr struct {
	name  string
	value *string
}

// Render satisfies Node.
func (a *attr) Render(w io.Writer) error {
	sw := &statefulWriter{w: w}

	if a.value == nil {
		sw.WriteString(" ")
		sw.WriteString(a.name)

		return sw.err
	}

	sw.WriteString(" ")
	sw.WriteString(a.name)
	sw.WriteString(`="`)
	sw.WriteString(template.HTMLEscapeString(*a.value))
	sw.WriteString(`"`)
	return sw.err
}

func (a *attr) Type() NodeType {
	return AttributeType
}

func joinClassAttrs(w *statefulWriter, classValues []string) {
	w.WriteString(` class="`)
	for _, v := range classValues {
		w.WriteString(v)
		w.WriteString(" ")
	}
	w.WriteString(`"`)
}

// String satisfies fmt.Stringer.
func (a *attr) String() string {
	var b strings.Builder
	_ = a.Render(&b)
	return b.String()
}

type group struct {
	children []Node
}

func (group) Type() NodeType {
	return ElementType
}

// String satisfies fmt.Stringer.
func (g group) String() string {
	panic("cannot render group directly")
}

// Render satisfies Node.
func (g group) Render(io.Writer) error {
	panic("cannot render group directly")
}

// Group multiple Nodes into one Node. Useful for concatenation of Nodes in variadic functions.
// The resulting Node cannot Render directly, trying it will panic.
// Render must happen through a parent element created with El or a helper.
func Group(children []Node) Node {
	return group{children: children}
}

// If condition is true, return the given Node. Otherwise, return nil.
// This helper function is good for inlining elements conditionally.
func If(condition bool, thenNode, elseNode Node) Node {
	if condition {
		return thenNode
	}
	return elseNode
}

type fragment struct {
	children []Node
}

func (fragment) Type() NodeType {
	return ElementType
}

// String satisfies fmt.Stringer.
func (f fragment) String() string {
	var b strings.Builder
	for _, c := range f.children {
		c.Render(&b)
	}
	return b.String()
}

// Render satisfies Node.
func (f *fragment) Render(w io.Writer) error {
	for _, c := range f.children {
		if err := c.Render(w); err != nil {
			return err
		}
	}
	return nil
}

// Fragment groups multiple nodes into one Node. Kind of like React.Fragment. 
// It has an easier api than Group for rendering a collection of nodes without
// specifiying a parent element.
func Fragment(children ...Node) Node {
	return &fragment{children: children}
}

func Static(children ...Node) Node {
	var sb strings.Builder
	err := Fragment(children...).Render(&sb)

	return NodeFunc(func(w io.Writer) error {
		if err != nil {
			return err
		}

		_, err := w.Write(StringToBytes(sb.String()))
		return err
	})
}
