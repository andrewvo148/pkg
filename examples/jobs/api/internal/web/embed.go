package web

import (
	"embed"
)

//go:embed swagger-ui/*
//go:embed index.html
//go:embed openapi/*
var WebUI embed.FS
