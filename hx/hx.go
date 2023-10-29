// Package hx provides utilities to work with htmx
package hx

import (
	g "github.com/alarbada/gomponents"
)

func Boost() g.Node { return g.Attr("hx-boost", "true") }

func Get(path string) g.Node    { return g.Attr("hx-get", path) }
func Post(path string) g.Node   { return g.Attr("hx-post", path) }
func Put(path string) g.Node    { return g.Attr("hx-put", path) }
func Delete(path string) g.Node { return g.Attr("hx-delete", path) }

func On(evt, code string) g.Node     { return g.Attr("hx-on:"+evt, code) }
func PushUrl(val string) g.Node      { return g.Attr("hx-push-url", val) }
func PushUrlT() g.Node               { return g.Attr("hx-push-url", "true") }
func Select(target string) g.Node    { return g.Attr("hx-select", target) }
func SelectOob(target string) g.Node { return g.Attr("hx-select-oob", target) }
func Swap(how string) g.Node         { return g.Attr("hx-swap", how) }
func SwapOob(how string) g.Node      { return g.Attr("hx-swap-oob", how) }
func Target(target string) g.Node    { return g.Attr("hx-target", target) }
func Trigger(trigger string) g.Node  { return g.Attr("hx-trigger", trigger) }
func Vals(vals string) g.Node        { return g.Attr("hx-vals", vals) }
func Ext(ext string) g.Node          { return g.Attr("hx-ext", ext) }
