package picture

import (
	"os"
	"io"
	"image"
	"image/png"
	"image/draw"

	"github.com/golang/freetype"
)

// Canvas wrapper image
type Canvas struct {
	img draw.Image
}

// NewCanvas return unicolor image
func NewCanvas(width, height int, color *image.Uniform) *Canvas {
	img := image.NewRGBA(image.Rectangle{Min: image.Pt(0, 0), Max: image.Pt(width, height)})
	draw.Draw(img, img.Bounds(), color, image.ZP, draw.Src)

	return &Canvas{img: img}
}

// NewCanvasWith make image with io.Reader
func NewCanvasWith(r io.Reader) (*Canvas, error) {
	baseImg, err := png.Decode(r)
	if err != nil {
		return nil, err
	}

	img := image.NewRGBA(baseImg.Bounds())
	draw.Draw(img, baseImg.Bounds(), image.White, image.ZP, draw.Src)
	draw.Draw(img, baseImg.Bounds(), baseImg, image.ZP, draw.Src)

	return &Canvas{img: img}, nil
}

// New return a new Canvas
func (c *Canvas) New() *Canvas {
	img := image.NewRGBA(c.img.Bounds())
	draw.Draw(img, c.img.Bounds(), c.img, image.ZP, draw.Src)
	return &Canvas{img: img}
}

// Image return draw.Image
func (c *Canvas) Image() draw.Image {
	return c.img
}

// Draw use other image mask to canvas
func (c *Canvas) Draw(pos image.Point, m image.Image) error {
	rect := m.Bounds().Add(pos)
	draw.Draw(c.img, rect, m, image.ZP, draw.Src)
	return nil
}

// DrawText draw text to canvas
func (c *Canvas) DrawText(pos image.Point, context *Context, strs []string) error {
	context.SetDst(c.img.Bounds(), c.img)
	pt := freetype.Pt(pos.X, pos.Y+int(context.PointToFixed(context.Size())>>6))
	for _, s := range strs {
		if _, err := context.DrawString(s, pt); err != nil {
			return err
		}

		pt.Y += context.PointToFixed(context.Size() * context.spacing)
	}
	return nil
}

// WriteTo write image to file, the file arg without expanded-name
func (c *Canvas) WriteTo(file string) (fullName string, err error) {
	fullName = file + ".png"
	fd, err := os.Create(fullName)
	if err != nil {
		return "", err
	}

	defer fd.Close()
	if err := png.Encode(fd, c.img); err != nil {
		return "", err
	}

	return fullName, nil
}
