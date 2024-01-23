package openapi

import (
	"embed"
)

//go:embed index.html
//go:embed bookstore.swagger.json
var SwaggerUI embed.FS
