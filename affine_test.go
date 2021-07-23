package affine

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	A float64 = 2e-5
	B float64 = 0
	C float64 = 116.5
	D float64 = 0
	E float64 = 2e-5
	F float64 = 36.8
)

func TestAffineFromGdal(t *testing.T) {
	aff := Affine{}
	gdalAff := [6]float64{C, A, B, F, D, E}
	aff.FromGdal(gdalAff)
	assert.Equal(t, aff.A, A)
	assert.Equal(t, aff.B, B)
	assert.Equal(t, aff.C, C)
	assert.Equal(t, aff.D, D)
	assert.Equal(t, aff.E, E)
	assert.Equal(t, aff.F, F)
}

func TestAffineToGdal(t *testing.T) {
	aff := Affine{A, B, C, D, E, F}
	affGdal := aff.ToGdal()
	assert.Equal(t, affGdal[0], C)
	assert.Equal(t, affGdal[1], A)
	assert.Equal(t, affGdal[2], B)
	assert.Equal(t, affGdal[3], F)
	assert.Equal(t, affGdal[4], D)
	assert.Equal(t, affGdal[5], E)
}

func TestAffineIdentity(t *testing.T) {
	aff := Affine{}
	aff = Identity()
	assert.Equal(t, aff.A, 1.0)
	assert.Equal(t, aff.B, 0.0)
	assert.Equal(t, aff.C, 0.0)
	assert.Equal(t, aff.D, 0.0)
	assert.Equal(t, aff.E, 1.0)
	assert.Equal(t, aff.F, 0.0)
}

func TestAffineTranslation(t *testing.T) {
	aff := Affine{}
	aff = Translation(1.0, 5.0)
	assert.Equal(t, aff.A, 1.0)
	assert.Equal(t, aff.B, 0.0)
	assert.Equal(t, aff.C, 1.0)
	assert.Equal(t, aff.D, 0.0)
	assert.Equal(t, aff.E, 1.0)
	assert.Equal(t, aff.F, 5.0)
}

func TestAffineScale(t *testing.T) {
	aff := Affine{}
	aff = Scale(2.0)
	assert.Equal(t, aff.A, 2.0)
	assert.Equal(t, aff.B, 0.0)
	assert.Equal(t, aff.C, 0.0)
	assert.Equal(t, aff.D, 0.0)
	assert.Equal(t, aff.E, 2.0)
	assert.Equal(t, aff.F, 0.0)
}

func TestAffineRotation(t *testing.T) {
	aff := Affine{}
	aff = Rotation(45.0, [2]float64{0.0, 0.0})
	assert.Equal(t, aff.A, 0.7071067811865476)
	assert.Equal(t, aff.B, -0.7071067811865475)
	assert.Equal(t, aff.C, 0.0)
	assert.Equal(t, aff.D, 0.7071067811865475)
	assert.Equal(t, aff.E, 0.7071067811865476)
	assert.Equal(t, aff.F, 0.0)
}

func TestAffineXY(t *testing.T) {
	aff := Affine{}
	aff.FromGdal([6]float64{-237481.5, 425.0, 0.0, 237536.4, 0.0, -425.0})
	x, y := aff.XY(0, 100)
	assert.Equal(t, x, -237481.5)
	assert.Equal(t, y, 195036.4)
}

func TestAffineColRow(t *testing.T) {
	aff := Affine{}
	aff.FromGdal([6]float64{-237481.5, 425.0, 0.0, 237536.4, 0.0, -425.0})
	col, row := aff.ColRow(-237481.5, 195036.4)
	assert.Equal(t, col, 0)
	assert.Equal(t, row, 100)
}

func TestAffine(t *testing.T) {
	aff := Affine{2e-5, 0, 116.5, 0, 2e-5, 36.8}
	fmt.Println(aff)
	aff2 := Identity()
	fmt.Println(aff2)
}

func TestAffineMul(t *testing.T) {
	aff := Translation(1.0, 5.0)
	affO := Rotation(45.0, [2]float64{0.0, 0.0})
	aff.Mul(affO)
	fmt.Println(aff)
	assert.Equal(t, aff,
		Affine{0.7071067811865476, -0.7071067811865475, 1.0,
			0.7071067811865475, 0.7071067811865476, 5.0})
}
