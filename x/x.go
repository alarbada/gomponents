package x

import (
	g "github.com/alarbada/gomponents"
)

func Data(value string) g.Node { return g.Attr("x-data", value) }
func On(value string) g.Node   { return g.Attr("x-on", value) }
func Init(value string) g.Node { return g.Attr("x-init", value) }
