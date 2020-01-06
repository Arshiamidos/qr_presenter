package painter

import (
	"fmt"
	"image/color"
	"qr/utils"
)
var boom *Boom
var brush Brush
var BRUSH_SIZE=10
var bits="00100000000100011110110011101100010110111110110000010001000100010000101100010001111011001110110001111000111011000001000100010001110100010001000111101100111011000111001011101100000100010001000111011100000100011110110011101100000100010001000101001101111011000100001100010001111011001110110001000000111011000001000100010001111011000001000111101100111011000001000100010001000100011110110011101100000100011110110011101100000100011110110000010001000100011110110000010001111011001110110000010001000100011000110100001000111111011111110111001000011111111101000011010000011100010011101011010000110100001001101111010101110111101101111001100101010001011001010010010100001001010010010111111101010101101101001111010001100011011000110110001001001000101000001010000010111001100000010011100011111000110000110011001110001100000011000000100101010010111011011010110110111100011111000101011101111001010011010001010010011001110110011100000110011101011111110111111101001001010010010110110100111101011001000011111101000011010000110110101011101010111110100111011111110101010000000000010000000100000000000"
func locateDarkPattern(v int) (int, int) {
	return (4 * v) + 9, 8
}

func Clear() {

}

func AddFinderPattern() {

	//top left
	brush.ChangeColor(color.RGBA{0,0,0,255})
	boom.DrawRect(1,1,7,7)
	brush.ChangeColor(color.RGBA{255,255,255,255})
	boom.DrawRect(2,2,6,6)
	brush.ChangeColor(color.RGBA{0,0,0,255})
	boom.DrawRect(3,3,5,5)

	//bottom left
	brush.ChangeColor(color.RGBA{0,0,0,255})
	boom.DrawRect(
		1,
		boom.boom.Rect.Dy()/BRUSH_SIZE-7,
		7,
		boom.boom.Rect.Dy()/BRUSH_SIZE)

	brush.ChangeColor(color.RGBA{255,255,255,255})
	boom.DrawRect(
		2,
		boom.boom.Rect.Dy()/BRUSH_SIZE-6,
		6,
		boom.boom.Rect.Dy()/BRUSH_SIZE-2)

	brush.ChangeColor(color.RGBA{0,0,0,255})
	boom.DrawRect(
		3,
		boom.boom.Rect.Dy()/BRUSH_SIZE-5,
		5,
		boom.boom.Rect.Dy()/BRUSH_SIZE-3)
	


	//right top	
	brush.ChangeColor(color.RGBA{0,0,0,255})
	boom.DrawRect(
		boom.boom.Rect.Dx()/BRUSH_SIZE-7,
		1,
		boom.boom.Rect.Dx()/BRUSH_SIZE,
		7)

	brush.ChangeColor(color.RGBA{255,255,255,255})
	boom.DrawRect(
		boom.boom.Rect.Dy()/BRUSH_SIZE-6,
		2,
		boom.boom.Rect.Dy()/BRUSH_SIZE-2,
		6)

	brush.ChangeColor(color.RGBA{0,0,0,255})
	boom.DrawRect(
		boom.boom.Rect.Dy()/BRUSH_SIZE-5,
		3,
		boom.boom.Rect.Dy()/BRUSH_SIZE-3,
		5)
}
func AddSepratorPattern() {

	brush.ChangeColor(color.RGBA{255,255,255,255})
	//top left seps
	boom.DrawLineH(1,8,8)
	boom.DrawLineV(8,1,8)
	//top bottom seps
	boom.DrawLineH(1,boom.boom.Rect.Dx()/BRUSH_SIZE-8,8)
	boom.DrawLineV(8,boom.boom.Rect.Dy()/BRUSH_SIZE-8,boom.boom.Rect.Dx()/BRUSH_SIZE)
	//top right sep
	boom.DrawLineH(boom.boom.Rect.Dx()/BRUSH_SIZE-8,8,boom.boom.Rect.Dx()/BRUSH_SIZE)
	boom.DrawLineV(boom.boom.Rect.Dx()/BRUSH_SIZE-8,1,8)
}

func AddAlignmentPattern(v int) {
	if v==1{
		return
	}
	aligns_:=utils.AlignmentPatternLocation(v)
	aligns:=[]int{}
	for  _,k := range aligns_ {
		if k!=0{
			aligns=append(aligns,k)
		}
	}
	for i := 0; i < len(aligns); i++ {
		for j := 0; j < len(aligns); j++ {
			if !((i==0 && j==0) || (i==0 &&j==len(aligns)-1) || (j==0 &&i==len(aligns)-1) ) {
				//fmt.Println(aligns[i]+1,aligns[j]+1)
				brush.ChangeColor(color.RGBA{0,0,0,255})
				boom.DrawRect(aligns[i]+1-2,aligns[j]+1-2,aligns[i]+1+2,aligns[j]+1+2)
				brush.ChangeColor(color.RGBA{255,255,255,255})
				boom.DrawRect(aligns[i]+1-1,aligns[j]+1-1,aligns[i]+1+1,aligns[j]+1+1)	
				brush.ChangeColor(color.RGBA{0,0,0,255})
				boom.Draw(aligns[i]+1,aligns[j]+1)	

			}
		}
	}
}
func AddTimingPattern() {
	boom.DrawLineZebraH(9,7,boom.boom.Rect.Dx()/BRUSH_SIZE-9)
	boom.DrawLineZebraV(7,9,boom.boom.Rect.Dy()/BRUSH_SIZE-9)
}
func AddDarkModule() {
	brush.ChangeColor(color.RGBA{0,0,0,255})
	boom.Draw(9,boom.boom.Rect.Dy()/BRUSH_SIZE-8)	
}

func Reserve_VersionInformationArea_FormatInformationArea(v int) {
	brush.ChangeColor(color.RGBA{0,0,255,255})

	boom.DrawLineH(boom.boom.Rect.Dx()/BRUSH_SIZE-8,9,boom.boom.Rect.Dx()/BRUSH_SIZE)
	boom.DrawLineV(9,boom.boom.Rect.Dy()/BRUSH_SIZE-7,boom.boom.Rect.Dy()/BRUSH_SIZE)

	if v>=7 {
		boom.DrawRect(1,boom.boom.Rect.Dy()/BRUSH_SIZE-11,6,boom.boom.Rect.Dy()/BRUSH_SIZE-9)
		boom.DrawRect(boom.boom.Rect.Dx()/BRUSH_SIZE-11,1,boom.boom.Rect.Dx()/BRUSH_SIZE-9,6)
	}
}
func FillData(){

	r:=len(boom.dirty)-1
	c:=len(boom.dirty[0])-1
	fmt.Println("...",len(boom.dirty),len(boom.dirty[0]))

	for c!=0{
		for i := r; i >0; i-- {
			fmt.Println("> ",c,i)
			if(!boom.dirty[i][c]){
				if bits[0]=='0'{
					brush.ChangeColor(color.RGBA{255,0,0,255})
				}else{
					brush.ChangeColor(color.RGBA{0,255,0,255})
				}
				fmt.Println(c,i)
				boom.Draw(i,c)
				bits=bits[1:]
			}
		}
		c=c-1
		if c == 7{
			c=c-1
		}

	}
	
}
func CalcPenalties() {

}
func CalcPenalty_1() {}
func CalcPenalty_2() {}
func CalcPenalty_3() {}
func CalcPenalty_4() {}

func PaintV(v int,brsh_size int ) {

	BRUSH_SIZE=brsh_size
	V := utils.QrSize(v)
	boom = NewBoom((V+1)*BRUSH_SIZE, (V+1)*BRUSH_SIZE,v)
	brush = NewBrush(BRUSH_SIZE, BRUSH_SIZE)
	boom.SetBrush(&brush)
	brush.ChangeColor(color.RGBA{255,255,255,255})

	AddFinderPattern()
	AddSepratorPattern()
	AddAlignmentPattern(v)
	AddTimingPattern()
	AddDarkModule()
	Reserve_VersionInformationArea_FormatInformationArea(v)
	FillData();

	boom.SaveBoom()
	fmt.Println(V)

}
