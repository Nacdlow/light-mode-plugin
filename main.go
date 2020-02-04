package main

import (
	"fmt"
	macaron "gopkg.in/macaron.v1"
	"html/template"
)

var (
	ID      = "light-mode"         // Plugin ID (required)
	NAME    = "Light Mode"         // Plugin name (required)
	AUTHOR  = "Nacdlow"            // Plugin author (required)
	VERSION = "v0.1"               // Plugin version (required)
	ROUTE   = "/plugin/light-mode" // The route of the plugin page (optional)
)

// Load runs when the plugin is loaded.
func Load() {
	fmt.Println("I'm light! yeh :D")
}

func Middleware() macaron.Handler {
	return func(ctx *macaron.Context) {
		ctx.Data["HeadExtra"] = template.HTML(`
<style>
body {
    background-color: white!important;
    color: black!important;
}
</style>
		`)

		if ctx.Req.RequestURI == ROUTE {
			if ctx.Data["LoggedIn"] == 1 {
				ctx.Data["Content"] = template.HTML(`
<form class="text-center" action="#!">
    <p class="h4 mb-4">Light Mode</p>
<div class="custom-control custom-switch">
  <input type="checkbox" class="custom-control-input" id="customSwitches" checked="1">
  <label class="custom-control-label" for="customSwitches">Enable Light Mode</label>
</div>
</form>
`)
				ctx.HTML(200, "blank")
			} else {
				ctx.Redirect("/")
			}
		}
	}
}
