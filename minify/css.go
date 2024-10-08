package minify

import (
	"bytes"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
)

type Params struct {
	IsMinifyCss   bool
	IsOptimizeCss bool
	IsMinifyJs    bool
	IsOptimizeJs  bool
}

func (p Params) IsOptiMiniCss() bool {
	return p.IsMinifyCss && p.IsOptimizeCss
}

func (p Params) IsOptiMiniJs() bool {
	return p.IsMinifyJs && p.IsOptimizeJs
}

func MinifyCSS(cssContent string) (string, error) {
	// Create a new minifier
	m := minify.New()

	// Add the CSS minifier
	m.AddFunc("text/css", css.Minify)

	// Create a buffer to store the minified CSS
	var b bytes.Buffer

	// Minify the CSS content
	if err := m.Minify("text/css", &b, bytes.NewBufferString(cssContent)); err != nil {
		return "", err
	}

	return b.String(), nil
}
