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
func (b *Boom) DrawRect(x int,y int , xx int , yy int){
	for i := x; i <= xx; i++ {
		for j := y; j <= yy; j++ {
			boom.Draw(i, j)
		}
	}
}
func (b *Boom) DrawLineH(x int,y int , xx int ){
	for i := x; i <= xx; i++ {
		b.Draw(i,y)
	}
}
func (b *Boom) DrawLineV(x int,y int , yy int){
	for i := y; i <= yy; i++ {
		b.Draw(x,i)
	}
}
func (b *Boom) Draw(x int,y int) {
	brs:=b.brush.siz.Bounds().Dx()
	r:=image.Rect(x*brs, y*brs, x*brs+brs,y*brs+brs )
	draw.Draw(b.boom, r.Bounds(), &image.Uniform{b.brush.clr}, image.ZP, draw.Src)
}
