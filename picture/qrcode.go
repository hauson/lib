package picture

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/skip2/go-qrcode"
)

//QRCode make qr code demo:QRCode("https://bycoin.im?invitecode=XXXYYY", 256)
func QRCode(url string, size int) (image.Image, error) {
	qr, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		return nil, err
	}

	qr.BackgroundColor = color.White
	qr.ForegroundColor = color.Black

	margin := size / 7
	img := NewCanvas(size, size, image.White).Image()
	draw.Draw(img, img.Bounds(), qr.Image(size+2*margin), image.Pt(margin, margin), draw.Src)
	return img, nil
}
