package utils

import (
	"github.com/thedevsaddam/renderer"
)

type resp *renderer.Render

var (
	instance resp
)

func ResponseRender() resp {

	if instance == nil {
		instance = renderer.New()
	}

	return instance
}
