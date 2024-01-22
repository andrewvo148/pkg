package web

import (
	"embed"
)

//go:embed swagger-ui/*
//go:embed index.html
var SwaggerUI embed.FS
