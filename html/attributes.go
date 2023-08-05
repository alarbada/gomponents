package html

import (
	g "github.com/alarbada/gomponents"
)

func Accept(v string) g.Node         { return g.Attr("accept", v) }
func Action(v string) g.Node         { return g.Attr("action", v) }
func Alt(v string) g.Node            { return g.Attr("alt", v) }
func Aria(name, v string) g.Node     { return g.Attr("aria-"+name, v) }
func As(v string) g.Node             { return g.Attr("as", v) }
func Async() g.Node                  { return g.Attr("async") }
func AutoComplete(v string) g.Node   { return g.Attr("autocomplete", v) }
func AutoFocus() g.Node              { return g.Attr("autofocus") }
func AutoPlay() g.Node               { return g.Attr("autoplay") }
func Charset(v string) g.Node        { return g.Attr("charset", v) }
func Checked() g.Node                { return g.Attr("checked") }
func Class(v string) g.Node          { return g.Attr("class", v) }
func ColSpan(v string) g.Node        { return g.Attr("colspan", v) }
func Cols(v string) g.Node           { return g.Attr("cols", v) }
func Content(v string) g.Node        { return g.Attr("content", v) }
func Controls() g.Node               { return g.Attr("controls") }
func DataAttr(name, v string) g.Node { return g.Attr("data-"+name, v) }
func Defer() g.Node                  { return g.Attr("defer") }
func Disabled() g.Node               { return g.Attr("disabled") }
func EncType(v string) g.Node        { return g.Attr("enctype", v) }
func For(v string) g.Node            { return g.Attr("for", v) }
func FormAttr(v string) g.Node       { return g.Attr("form", v) }
func Height(v string) g.Node         { return g.Attr("height", v) }
func Href(v string) g.Node           { return g.Attr("href", v) }
func ID(v string) g.Node             { return g.Attr("id", v) }
func Lang(v string) g.Node           { return g.Attr("lang", v) }
func Loading(v string) g.Node        { return g.Attr("loading", v) }
func Loop() g.Node                   { return g.Attr("loop") }
func Max(v string) g.Node            { return g.Attr("max", v) }
func MaxLength(v string) g.Node      { return g.Attr("maxlength", v) }
func Method(v string) g.Node         { return g.Attr("method", v) }
func Min(v string) g.Node            { return g.Attr("min", v) }
func MinLength(v string) g.Node      { return g.Attr("minlength", v) }
func Multiple() g.Node               { return g.Attr("multiple") }
func Muted() g.Node                  { return g.Attr("muted") }
func Name(v string) g.Node           { return g.Attr("name", v) }
func Pattern(v string) g.Node        { return g.Attr("pattern", v) }
func Placeholder(v string) g.Node    { return g.Attr("placeholder", v) }
func PlaysInline() g.Node            { return g.Attr("playsinline") }
func Poster(v string) g.Node         { return g.Attr("poster", v) }
func Preload(v string) g.Node        { return g.Attr("preload", v) }
func ReadOnly() g.Node               { return g.Attr("readonly") }
func Rel(v string) g.Node            { return g.Attr("rel", v) }
func Required() g.Node               { return g.Attr("required") }
func Role(v string) g.Node           { return g.Attr("role", v) }
func RowSpan(v string) g.Node        { return g.Attr("rowspan", v) }
func Rows(v string) g.Node           { return g.Attr("rows", v) }
func Selected() g.Node               { return g.Attr("selected") }
func Slot(v string) g.Node           { return g.Attr("slot", v) }
func Src(v string) g.Node            { return g.Attr("src", v) }
func SrcSet(v string) g.Node         { return g.Attr("srcset", v) }
func Step(v string) g.Node           { return g.Attr("step", v) }
func StyleAttr(v string) g.Node      { return g.Attr("style", v) }
func TabIndex(v string) g.Node       { return g.Attr("tabindex", v) }
func Target(v string) g.Node         { return g.Attr("target", v) }
func TitleAttr(v string) g.Node      { return g.Attr("title", v) }
func Type(v string) g.Node           { return g.Attr("type", v) }
func Value(v string) g.Node          { return g.Attr("value", v) }
func Width(v string) g.Node          { return g.Attr("width", v) }
