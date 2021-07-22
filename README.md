# Affine
Affine transformation matrices for Golang

# Install
```bash
go get github.com/sshuair/affine
```

# Usage
The Affine matrices can be created by passing the values `A,B,C,D,E,F` to `Affine`.
It's also can be constructed by `Identity`, `Translation`, `Scale` and `Rotation`.

```golang
aff1 := Affine{2e-5, 0, 116.5, 0, 2e-5, 36.8}
fmt.Println(aff1) //{2e-05 0 116.5 0 2e-05 36.8}

aff2 := Affine{}
aff2 = aff2.Identity() //{1 0 0 0 1 0}

aff3 := Affine{}
aff3 = aff3.Translation(1.0, 5.0) //{1 0 1 0 1 5}

aff4 := Affine{}
aff4 = aff4.Scale(2.0) //{2 0 0 0 2 0}

aff5 := Affine{}
aff5 = aff5.Rotation(45.0, [2]float64{0.0, 0.0}) //{0.7071067811865476 -0.7071067811865475 0 0.7071067811865475 20.7071067811865476 0}
```

You can also get the x,y by using the pixel row,col.

```golang
aff := Affine{}
aff.FromGdal([6]float64{-237481.5, 425.0, 0.0, 237536.4, 0.0, -425.0})
x, y := aff.XY(100, 0)
fmt.Println(x,y) //-237481.5 195036.4
```