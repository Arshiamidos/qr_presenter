package painter

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"qr/utils"
)

type Boom struct {
	boom        *image.RGBA
	brush       *Brush
	dirty       [][]bool
	BlackWhites [][]bool
}

func NewBoom(w int, h int, v int) *Boom {

	dirty := make([][]bool, utils.QrSize(v)+1)
	BlackWhites := make([][]bool, utils.QrSize(v)+1)

	for i := range dirty {
		dirty[i] = make([]bool, utils.QrSize(v)+1)
		BlackWhites[i] = make([]bool, utils.QrSize(v)+1)

	}
	for i := 0; i < len(dirty); i++ {
		for j := 0; j < len(dirty[0]); j++ {
			if i == 0 || j == 0 {
				dirty[i][j] = true
				BlackWhites[i][j] = true
			}
			dirty[i][j] = false
			BlackWhites[i][j] = false
		}
	}

	return &Boom{
		image.NewRGBA(image.Rect(0, 0, w, h)),
		&Brush{},
		dirty,
		BlackWhites,
	}
}
func (b *Boom) SetBrush(br *Brush) {
	b.brush = br
}
func (b *Boom) SaveBoom(s string) {
	outputFile, _ := os.Create("QR_" + s + ".png")
	png.Encode(outputFile, b.boom)
	outputFile.Close()
}
func (b *Boom) DrawRect(x int, y int, xx int, yy int) {
	for i := x; i <= xx; i++ {
		for j := y; j <= yy; j++ {
			boom.Draw(i, j)
		}
	}
}
func (b *Boom) DrawLineH(x int, y int, xx int) {
	for i := x; i <= xx; i++ {
		b.Draw(i, y)
	}
}
func (b *Boom) DrawLineV(x int, y int, yy int) {
	for i := y; i <= yy; i++ {
		b.Draw(x, i)
	}
}
func (b *Boom) DrawLineZebraH(x int, y int, xx int) {

	for i := x; i <= xx; i++ {
		if i%2 != 0 {
			b.brush.ChangeColor(color.RGBA{0, 0, 0, 255})
			b.Draw(i, y)
		} else {
			b.brush.ChangeColor(color.RGBA{255, 255, 255, 255})
			b.Draw(i, y)
		}
	}
}
func (b *Boom) DrawLineZebraV(x int, y int, yy int) {
	for i := y; i <= yy; i++ {
		if i%2 != 0 {
			b.brush.ChangeColor(color.RGBA{0, 0, 0, 255})
			b.Draw(x, i)
		} else {
			b.brush.ChangeColor(color.RGBA{255, 255, 255, 255})
			b.Draw(x, i)
		}
	}
}
func (b *Boom) DrawAlign(x int, y int) {
	b.brush.ChangeColor(color.RGBA{255, 0, 0, 255})
	b.DrawRect(
		x-2*b.brush.siz.Bounds().Dx(),
		y-2*b.brush.siz.Bounds().Dx(),
		x+2*b.brush.siz.Bounds().Dx(),
		y+2*b.brush.siz.Bounds().Dx(),
	)
}
func (b *Boom) Draw(x int, y int) {
	brs := b.brush.siz.Bounds().Dx()
	r := image.Rect(x*brs, y*brs, x*brs+brs, y*brs+brs)
	b.dirty[x][y] = true
	if b.brush.clr.R == 255 {
		b.BlackWhites[x][y] = true

	} else {
		b.BlackWhites[x][y] = false
	}
	draw.Draw(b.boom, r.Bounds(), &image.Uniform{b.brush.clr}, image.ZP, draw.Src)
}
