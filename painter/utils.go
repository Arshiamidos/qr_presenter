package painter

import (
	"fmt"
	"image/color"
	"math"
	"qr/utils"
)

var boom *Boom
var brush Brush
var BRUSH_SIZE = 10

func locateDarkPattern(v int) (int, int) {
	return (4 * v) + 9, 8
}

func Clear() {

}

func AddFinderPattern() {

	//top left
	brush.ChangeColor(color.RGBA{0, 0, 0, 255})
	boom.DrawRect(1, 1, 7, 7)
	brush.ChangeColor(color.RGBA{255, 255, 255, 255})
	boom.DrawRect(2, 2, 6, 6)
	brush.ChangeColor(color.RGBA{0, 0, 0, 255})
	boom.DrawRect(3, 3, 5, 5)

	//bottom left
	brush.ChangeColor(color.RGBA{0, 0, 0, 255})
	boom.DrawRect(
		1,
		boom.boom.Rect.Dy()/BRUSH_SIZE-7,
		7,
		boom.boom.Rect.Dy()/BRUSH_SIZE-1)

	brush.ChangeColor(color.RGBA{255, 255, 255, 255})
	boom.DrawRect(
		2,
		boom.boom.Rect.Dy()/BRUSH_SIZE-6,
		6,
		boom.boom.Rect.Dy()/BRUSH_SIZE-2)

	brush.ChangeColor(color.RGBA{0, 0, 0, 255})
	boom.DrawRect(
		3,
		boom.boom.Rect.Dy()/BRUSH_SIZE-5,
		5,
		boom.boom.Rect.Dy()/BRUSH_SIZE-3)

	//right top
	brush.ChangeColor(color.RGBA{0, 0, 0, 255})
	boom.DrawRect(
		boom.boom.Rect.Dx()/BRUSH_SIZE-8,
		1,
		boom.boom.Rect.Dx()/BRUSH_SIZE-1,
		7)

	brush.ChangeColor(color.RGBA{255, 255, 255, 255})
	boom.DrawRect(
		boom.boom.Rect.Dy()/BRUSH_SIZE-6,
		2,
		boom.boom.Rect.Dy()/BRUSH_SIZE-2,
		6)

	brush.ChangeColor(color.RGBA{0, 0, 0, 255})
	boom.DrawRect(
		boom.boom.Rect.Dy()/BRUSH_SIZE-5,
		3,
		boom.boom.Rect.Dy()/BRUSH_SIZE-3,
		5)
}
func AddSepratorPattern() {

	brush.ChangeColor(color.RGBA{255, 255, 255, 255})
	//top left seps
	boom.DrawLineH(1, 8, 8)
	boom.DrawLineV(8, 1, 8)
	//top bottom seps
	boom.DrawLineH(1, boom.boom.Rect.Dx()/BRUSH_SIZE-8, 8)
	boom.DrawLineV(8, boom.boom.Rect.Dy()/BRUSH_SIZE-8, boom.boom.Rect.Dx()/BRUSH_SIZE-1)
	//top right sep
	boom.DrawLineH(boom.boom.Rect.Dx()/BRUSH_SIZE-8, 8, boom.boom.Rect.Dx()/BRUSH_SIZE-1)
	boom.DrawLineV(boom.boom.Rect.Dx()/BRUSH_SIZE-8, 1, 8)
}

func AddAlignmentPattern(v int) {
	if v == 1 {
		return
	}
	aligns_ := utils.AlignmentPatternLocation(v)
	aligns := []int{}
	for _, k := range aligns_ {
		if k != 0 {
			aligns = append(aligns, k)
		}
	}
	for i := 0; i < len(aligns); i++ {
		for j := 0; j < len(aligns); j++ {
			if !((i == 0 && j == 0) || (i == 0 && j == len(aligns)-1) || (j == 0 && i == len(aligns)-1)) {
				//fmt.Println(aligns[i]+1,aligns[j]+1)
				brush.ChangeColor(color.RGBA{0, 0, 0, 255})
				boom.DrawRect(aligns[i]+1-2, aligns[j]+1-2, aligns[i]+1+2, aligns[j]+1+2)
				brush.ChangeColor(color.RGBA{255, 255, 255, 255})
				boom.DrawRect(aligns[i]+1-1, aligns[j]+1-1, aligns[i]+1+1, aligns[j]+1+1)
				brush.ChangeColor(color.RGBA{0, 0, 0, 255})
				boom.Draw(aligns[i]+1, aligns[j]+1)

			}
		}
	}
}
func AddTimingPattern() {
	boom.DrawLineZebraH(9, 7, boom.boom.Rect.Dx()/BRUSH_SIZE-9)
	boom.DrawLineZebraV(7, 9, boom.boom.Rect.Dy()/BRUSH_SIZE-9)
}
func AddDarkModule() {
	brush.ChangeColor(color.RGBA{0, 0, 0, 255})
	boom.Draw(9, boom.boom.Rect.Dy()/BRUSH_SIZE-8)
}
func DrawOneColor(x int,y int,c rune ){
	if c=='1'{
		brush.ChangeColor(color.RGBA{0, 0, 0, 255})
	}else{
		brush.ChangeColor(color.RGBA{255, 255, 255, 255})
	}
	boom.Draw(x,y)
}
func VersionInfomationArea(v int){
	if v >= 7 {
		versionInformation:=utils.DrawVersionInformation(v)
		fmt.Println("<<< version ifo",versionInformation)
		DrawOneColor(  1,boom.boom.Bounds().Dy()/BRUSH_SIZE-11		,rune(versionInformation[17]		))
		DrawOneColor(  1,boom.boom.Bounds().Dy()/BRUSH_SIZE-10		,rune(versionInformation[16]		))
		DrawOneColor(  1,boom.boom.Bounds().Dy()/BRUSH_SIZE-9		,rune(versionInformation[15]		))
		DrawOneColor(  2,boom.boom.Bounds().Dy()/BRUSH_SIZE-11		,rune(versionInformation[14]		))
		DrawOneColor(  2,boom.boom.Bounds().Dy()/BRUSH_SIZE-10		,rune(versionInformation[13]		))
		DrawOneColor(  2,boom.boom.Bounds().Dy()/BRUSH_SIZE-9		,rune(versionInformation[12]		))
		DrawOneColor(  3,boom.boom.Bounds().Dy()/BRUSH_SIZE-11		,rune(versionInformation[11]		))
		DrawOneColor(  3,boom.boom.Bounds().Dy()/BRUSH_SIZE-10		,rune(versionInformation[10]		))
		DrawOneColor(  3,boom.boom.Bounds().Dy()/BRUSH_SIZE-9		,rune(versionInformation[9]		))
		DrawOneColor(  4,boom.boom.Bounds().Dy()/BRUSH_SIZE-11		,rune(versionInformation[8]		))
		DrawOneColor(  4,boom.boom.Bounds().Dy()/BRUSH_SIZE-10		,rune(versionInformation[7]		))
		DrawOneColor(  4,boom.boom.Bounds().Dy()/BRUSH_SIZE-9		,rune(versionInformation[6]		))
		DrawOneColor(  5,boom.boom.Bounds().Dy()/BRUSH_SIZE-11		,rune(versionInformation[5]		))
		DrawOneColor(  5,boom.boom.Bounds().Dy()/BRUSH_SIZE-10		,rune(versionInformation[4]		))
		DrawOneColor(  5,boom.boom.Bounds().Dy()/BRUSH_SIZE-9		,rune(versionInformation[3]		))
		DrawOneColor(  6,boom.boom.Bounds().Dy()/BRUSH_SIZE-11		,rune(versionInformation[2]		))
		DrawOneColor(  6,boom.boom.Bounds().Dy()/BRUSH_SIZE-10		,rune(versionInformation[1]		))
		DrawOneColor(  6,boom.boom.Bounds().Dy()/BRUSH_SIZE-9		,rune(versionInformation[0]		))

		DrawOneColor( boom.boom.Bounds().Dx()/BRUSH_SIZE-11 ,1 ,rune(versionInformation[17]		))
		DrawOneColor( boom.boom.Bounds().Dx()/BRUSH_SIZE-10 ,1 ,rune(versionInformation[16]		))
		DrawOneColor( boom.boom.Bounds().Dx()/BRUSH_SIZE-9	,1 ,rune(versionInformation[15]		))
		DrawOneColor( boom.boom.Bounds().Dx()/BRUSH_SIZE-11 ,2 ,rune(versionInformation[14]		))
		DrawOneColor( boom.boom.Bounds().Dx()/BRUSH_SIZE-10 ,2 ,rune(versionInformation[13]		))
		DrawOneColor( boom.boom.Bounds().Dx()/BRUSH_SIZE-9	,2 ,rune(versionInformation[12]		))
		DrawOneColor( boom.boom.Bounds().Dx()/BRUSH_SIZE-11 ,3 ,rune(versionInformation[11]		))
		DrawOneColor( boom.boom.Bounds().Dx()/BRUSH_SIZE-10 ,3 ,rune(versionInformation[10]		))
		DrawOneColor( boom.boom.Bounds().Dx()/BRUSH_SIZE-9	,3 ,rune(versionInformation[9]		))
		DrawOneColor( boom.boom.Bounds().Dx()/BRUSH_SIZE-11 ,4 ,rune(versionInformation[8]		))
		DrawOneColor( boom.boom.Bounds().Dx()/BRUSH_SIZE-10 ,4 ,rune(versionInformation[7]		))
		DrawOneColor( boom.boom.Bounds().Dx()/BRUSH_SIZE-9	,4 ,rune(versionInformation[6]		))
		DrawOneColor( boom.boom.Bounds().Dx()/BRUSH_SIZE-11 ,5 ,rune(versionInformation[5]		))
		DrawOneColor( boom.boom.Bounds().Dx()/BRUSH_SIZE-10 ,5 ,rune(versionInformation[4]		))
		DrawOneColor( boom.boom.Bounds().Dx()/BRUSH_SIZE-9	,5 ,rune(versionInformation[3]		))
		DrawOneColor( boom.boom.Bounds().Dx()/BRUSH_SIZE-11	,6 ,rune(versionInformation[2]		))
		DrawOneColor( boom.boom.Bounds().Dx()/BRUSH_SIZE-10	,6 ,rune(versionInformation[1]		))
		DrawOneColor( boom.boom.Bounds().Dx()/BRUSH_SIZE-9	,6 ,rune(versionInformation[0]		))


	}
}
func FormatInformationArea(v int, e rune,m int) {
	brush.ChangeColor(color.RGBA{255, 255, 255, 255})
	formatInfo:=utils.DrawFormatInformation(e,m)
	fmt.Println("<<< formatinfo",formatInfo,utils.QrSize(v),boom.boom.Bounds().Dx()/BRUSH_SIZE)
	
	DrawOneColor(1,9,rune(formatInfo[0]		))
	DrawOneColor(2,9,rune(formatInfo[1]		))
	DrawOneColor(3,9,rune(formatInfo[2]		))
	DrawOneColor(4,9,rune(formatInfo[3]		))
	DrawOneColor(5,9,rune(formatInfo[4]		))
	DrawOneColor(6,9,rune(formatInfo[5]		))
	DrawOneColor(8,9,rune(formatInfo[6]		))
	DrawOneColor(9,9,rune(formatInfo[7]		))
	DrawOneColor(9,8,rune(formatInfo[8]		))
	DrawOneColor(9,6,rune(formatInfo[9]		))
	DrawOneColor(9,5,rune(formatInfo[10]	))
	DrawOneColor(9,4,rune(formatInfo[11]	))
	DrawOneColor(9,3,rune(formatInfo[12]	))
	DrawOneColor(9,2,rune(formatInfo[13]	))
	DrawOneColor(9,1,rune(formatInfo[14]	))

	DrawOneColor(9,boom.boom.Bounds().Dy()/BRUSH_SIZE-0-1,		rune(formatInfo[0]			))
	DrawOneColor(9,boom.boom.Bounds().Dy()/BRUSH_SIZE-1-1,		rune(formatInfo[1]			))
	DrawOneColor(9,boom.boom.Bounds().Dy()/BRUSH_SIZE-2-1,		rune(formatInfo[2]			))
	DrawOneColor(9,boom.boom.Bounds().Dy()/BRUSH_SIZE-3-1,		rune(formatInfo[3]			))
	DrawOneColor(9,boom.boom.Bounds().Dy()/BRUSH_SIZE-4-1,		rune(formatInfo[4]			))
	DrawOneColor(9,boom.boom.Bounds().Dy()/BRUSH_SIZE-5-1,		rune(formatInfo[5]			))
	DrawOneColor(9,boom.boom.Bounds().Dy()/BRUSH_SIZE-6-1,		rune(formatInfo[6]			))

	DrawOneColor(  boom.boom.Bounds().Dx()/BRUSH_SIZE-7-1,9,	rune(formatInfo[7]		))
	DrawOneColor(  boom.boom.Bounds().Dx()/BRUSH_SIZE-6-1,9,	rune(formatInfo[8]		))
	DrawOneColor(  boom.boom.Bounds().Dx()/BRUSH_SIZE-5-1,9,	rune(formatInfo[9]		))
	DrawOneColor(  boom.boom.Bounds().Dx()/BRUSH_SIZE-4-1,9,	rune(formatInfo[10]		))
	DrawOneColor(  boom.boom.Bounds().Dx()/BRUSH_SIZE-3-1,9,	rune(formatInfo[11]		))
	DrawOneColor(  boom.boom.Bounds().Dx()/BRUSH_SIZE-2-1,9,	rune(formatInfo[12]		))
	DrawOneColor(  boom.boom.Bounds().Dx()/BRUSH_SIZE-1-1,9,	rune(formatInfo[13]		))
	DrawOneColor(  boom.boom.Bounds().Dx()/BRUSH_SIZE-0-1,9,	rune(formatInfo[14]		))

}
func DrawBoubleColumn(r int, i int, bits string, m int) string {

	for offset := 0; offset < 2; offset++ {
		if !boom.dirty[r-offset][i] {
			if int(int(bits[0])-'0')^GetMask(i, r-offset, m) == 1 {
				brush.ChangeColor(color.RGBA{0, 0, 0, 255})
			} else {
				brush.ChangeColor(color.RGBA{255, 255, 255, 255})
			}
			boom.Draw(r-offset, i)
			bits = bits[1:]
		}
	}
	
	return bits

}
func GetMask(row int, column int, m int) int {
	row = row - 1
	column = column - 1
	result := -1
	switch m {
	case 0:
		result = (row + column) % 2
	case 1:
		result = (row) % 2
	case 2:
		result = (column) % 3
	case 3:
		result = (row + column) % 3
	case 4:
		result = int(math.Floor(float64(row/2))+math.Floor(float64(column/3))) % 2
	case 5:
		result = ((row * column) % 2) + ((row * column) % 3)
	case 6:
		result = (((row * column) % 2) + ((row * column) % 3)) % 2
	case 7:
		result = (((row + column) % 2) + ((row * column) % 3)) % 2
	default:
		panic("not found mask pattern")
	}
	if result == 0 {
		return 1
	}
	return 0

}
func FillDataWithMask(_bits string, mask int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	r := len(boom.dirty) - 1
	c := r
	columnReverse := true
	bits:=_bits
	for len(bits)!= 0 {
		if columnReverse {
			for i := c; i > 0; i-- {
					bits = DrawBoubleColumn(r, i, bits, mask)
			}
		} else {
			for i := 1; i <= c; i++ {
					bits = DrawBoubleColumn(r, i, bits, mask)

			}
		}
		r = r - 2
		columnReverse = !columnReverse
		if r == 7 {
			r = r - 1
		}
	}
}

func CalcPenalties() int{
	return CalcPenalty_1()+CalcPenalty_2()+CalcPenalty_3()+CalcPenalty_4()
}
func CalcPenalty_1() int{
	return 0
}
func CalcPenalty_2() int{
	return 0

}
func CalcPenalty_3() int{
	return 0

}
func CalcPenalty_4() int{
	return 0

}

func freshCopy(opy [][]bool) {
	boom.dirty = make([][]bool, len(opy))
	for i := range opy {
		boom.dirty[i] = make([]bool, len(opy[i]))
		copy(boom.dirty[i], opy[i])
	}
}
func prettyPrint(){
	for i := 1; i <= len(boom.dirty)-1; i++ {
		for j := 1; j <= len(boom.dirty[i])-1; j++ {
			if(boom.BlackWhites[i][j]==true){
				fmt.Print(" ⬜️")
			}else{
				fmt.Print(" ⬛️")
			}
		}
		fmt.Println()
	}
	fmt.Println("\n")
}
func PaintV(v int,e rune, brsh_size int, bits string) {

	BRUSH_SIZE = brsh_size
	V := utils.QrSize(v)
	boom = NewBoom((V+1)*BRUSH_SIZE, (V+1)*BRUSH_SIZE, v)
	brush = NewBrush(BRUSH_SIZE, BRUSH_SIZE)
	boom.SetBrush(&brush)
	brush.ChangeColor(color.RGBA{255, 255, 255, 255})

	AddFinderPattern()
	AddSepratorPattern()
	AddAlignmentPattern(v)
	AddTimingPattern()
	AddDarkModule()
	VersionInfomationArea(v)
	duplicate := make([][]bool, len(boom.dirty))
	for i := range boom.dirty {
		duplicate[i] = make([]bool, len(boom.dirty[i]))
		copy(duplicate[i], boom.dirty[i])
	}

	p:=[8]int{}

	FormatInformationArea(v,e,0)
	FillDataWithMask(bits, 0)
	boom.SaveBoom("0")
	p[0]=CalcPenalties()
	prettyPrint()
	freshCopy(duplicate)

	FormatInformationArea(v,e,1)
	FillDataWithMask(bits, 1)
	boom.SaveBoom("1")
	p[1]=CalcPenalties()
	prettyPrint()
	freshCopy(duplicate)

	FormatInformationArea(v,e,2)
	FillDataWithMask(bits, 2)
	boom.SaveBoom("2")
	p[2]=CalcPenalties()
	prettyPrint()
	freshCopy(duplicate)

	FormatInformationArea(v,e,3)
	FillDataWithMask(bits, 3)
	boom.SaveBoom("3")
	p[3]=CalcPenalties()
	prettyPrint()
	freshCopy(duplicate)

	FormatInformationArea(v,e,4)
	FillDataWithMask(bits, 4)
	boom.SaveBoom("4")
	p[4]=CalcPenalties()
	prettyPrint()
	freshCopy(duplicate)

	FormatInformationArea(v,e,5)
	FillDataWithMask(bits, 5)
	boom.SaveBoom("5")
	p[5]=CalcPenalties()
	prettyPrint()
	freshCopy(duplicate)

	FormatInformationArea(v,e,6)
	FillDataWithMask(bits, 6)
	boom.SaveBoom("6")
	p[6]=CalcPenalties()
	prettyPrint()
	freshCopy(duplicate)

	FormatInformationArea(v,e,7)
	FillDataWithMask(bits, 7)
	boom.SaveBoom("7")
	p[7]=CalcPenalties()
	prettyPrint()
	freshCopy(duplicate)

}
