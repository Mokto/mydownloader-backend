package controllers

import (
	"github.com/revel/revel"
)

// Options Controller
type Options struct {
	*revel.Controller
}

// Get return empty reponse
func (c Options) Get() revel.Result {
	return c.RenderText("")
}
