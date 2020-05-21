package picture

import (
	"image"
	"io/ioutil"
	"image/draw"

	"github.com/golang/freetype"
	"golang.org/x/image/math/fixed"
	imageFont "golang.org/x/image/font"
)

// Context text setting
type Context struct {
	color   *image.Uniform
	context *freetype.Context
	size    float64
	spacing float64
}

// NewContext new a text setting
func NewContext(fontfile string, size, dpi, spacing float64, color *image.Uniform) (*Context, error) {
	fontBytes, err := ioutil.ReadFile(fontfile)
	if err != nil {
		return nil, err
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}

	context := freetype.NewContext()
	context.SetDPI(dpi)
	context.SetHinting(imageFont.HintingNone)
	context.SetFont(font)
	context.SetFontSize(size)
	context.SetSrc(color)

	return &Context{
		size:    size,
		color:   color,
		spacing: spacing,
		context: context,
	}, nil
}

// Size font size
func (c *Context) Size() float64 {
	return c.size
}

// PointToFixed converts the given number of points (as in "a 12 point font")
// into a 26.6 fixed point number of pixels.
func (c *Context) PointToFixed(x float64) fixed.Int26_6 {
	return c.context.PointToFixed(x)
}

// SetDst set draw dst
func (c *Context) SetDst(clip image.Rectangle, dst draw.Image) {
	c.context.SetClip(clip)
	c.context.SetDst(dst)
}

// DrawString draw string
func (c *Context) DrawString(s string, p fixed.Point26_6) (fixed.Point26_6, error) {
	return c.context.DrawString(s, p)
}
