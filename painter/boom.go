package painter

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

type Boom struct {
	boom *image.RGBA
	brush *Brush
}

func NewBoom(w int, h int) *Boom {

	return &Boom{
		image.NewRGBA(image.Rect(0, 0, w, h)),
		&Brush{},
	}
}
func (b *Boom) SetBrush(br *Brush) {
	b.brush = br
}
func (b *Boom) SaveBoom() {
	outputFile, _ := os.Create("QR.png")
	png.Encode(outputFile, b.boom)
	outputFile.Close()
}

func (b *Boom) Draw(x int,y int) {

	r:=image.Rect(x, y, x+b.brush.siz.Bounds().Dx(),y+b.brush.siz.Bounds().Dy() )
	draw.Draw(b.boom, r.Bounds(), &image.Uniform{b.brush.clr}, image.ZP, draw.Src)
}
