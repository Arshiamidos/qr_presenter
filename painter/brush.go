package painter

import "image"

import "image/color"

type Brush struct {
	siz *image.RGBA
	clr color.RGBA
}

func  NewBrush(w int,h int ) Brush{
	return Brush{
		image.NewRGBA(image.Rect(0, 0, w, h)),
		color.RGBA{255, 0, 0, 255},
	}
}
func (b *Brush) ChangeColor(c color.RGBA){
	b.clr= c
}
