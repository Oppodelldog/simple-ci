package assets

import (
	"github.com/go-playground/statics/static"
)

const assetsPath = "./webview/assets"

var Templates *static.Files
var Images *static.Files

func init() {
	var err error

	config := &static.Config{
		UseStaticFiles: true,
		FallbackToDisk: false,
		AbsPkgPath:     assetsPath,
	}

	Templates, err = newStaticTemplates(config)
	if err != nil {
		panic(err)
	}

	Images, err = newStaticImages(config)
	if err != nil {
		panic(err)
	}
}
