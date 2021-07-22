package affine

import (
	"math"
)

func cosSinDeg(deg float64) (float64, float64) {
	deg = math.Mod(deg, 360.0)
	switch deg {
	case 90.0:
		return 0.0, 1.0
	case 180.0:
		return -1.0, 0.0
	case 270.0:
		return 0.0, -1.0
	}
	rad := deg * math.Pi / 180.0
	return math.Cos(rad), math.Sin(rad)
}

// https://www.zhihu.com/question/20666664
// Affine transform for translate bettwen spatial reference system and pixel reference system
type Affine struct {
	A float64 //width of a pixel
	B float64 //row rotation (typically zero)
	C float64 //x-coordinate of the upper-left corner of the upper-left pixel
	D float64 //column rotation (typically zero)
	E float64 //height of a pixel (typically negative)
	F float64 //y-coordinate of the of the upper-left corner of the upper-left pixel
}

// Create the identity transform
// | x' |   | 1  0  0 | | x |
// | y' | = | 0  1  0 | | y |
// | 1  |   | 0  0  1 | | 1 |
func (aff Affine) Identity() Affine {
	newAff := Affine{
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
	}
	return newAff
}

// Create a translation transform from an offset vector
// | x' |   | 1  0  xoff | | x |
// | y' | = | 0  1  yoff | | y |
// | 1  |   | 0  0  1    | | 1 |
func (aff Affine) Translation(xoff float64, yoff float64) Affine {
	newAff := Affine{
		1, 0, xoff,
		0, 1, yoff,
	}
	return newAff
}

// Create a scaling transform from a scalar
// | x' |   | scale  0  1 | | x |
// | y' | = | 0  scale  1 | | y |
// | 1  |   | 0      0  1 | | 1 |
func (aff Affine) Scale(scaling float64) Affine {
	newAff := Affine{
		scaling, 0, 0,
		0, scaling, 0,
	}
	return newAff
}

// Create a scaling transform from a scalar
// | x' |   | c   s  1 | | x |
// | y' | = | -s  c  1 | | y |
// | 1  |   | 0   0  1 | | 1 |
func (aff Affine) Rotation(angle float64, pivot [2]float64) Affine {
	ca, sa := cosSinDeg(angle)
	px, py := pivot[0], pivot[1]
	newAff := Affine{
		ca, -sa, px - px*ca + py*sa,
		sa, ca, py - px*sa - py*ca,
	}
	return newAff
}

// get the affine params from gdal
func (aff *Affine) FromGdal(affGdal [6]float64) {
	aff.A = affGdal[1]
	aff.B = affGdal[2]
	aff.C = affGdal[0]
	aff.D = affGdal[4]
	aff.E = affGdal[5]
	aff.F = affGdal[3]
}

// get convert the affine transform to gdal order
func (aff Affine) ToGdal() [6]float64 {
	var gt [6]float64
	gt[1] = aff.A
	gt[2] = aff.B
	gt[0] = aff.C
	gt[4] = aff.D
	gt[5] = aff.E
	gt[3] = aff.F
	return gt
}

// get the x,y from the pixel row,col
func (aff Affine) XY(row int, col int) (float64, float64) {
	var x, y float64
	x = aff.A*float64(col) + aff.C
	y = aff.E*float64(row) + aff.F
	return x, y
}

// convert the spatial reference system x,y to row,col
// note the x corspand to col, and y corespond to row
func (aff Affine) RowCol(x float64, y float64) (int, int) {
	var row, col int
	row = int(math.Floor((x - aff.C) / aff.A))
	col = int(math.Floor((y - aff.F) / aff.E))
	return row, col
}
