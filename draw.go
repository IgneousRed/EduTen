package EduTen

import (
	"image/color"

	m "github.com/IgneousRed/gomisc"
	eb "github.com/hajimehoshi/ebiten/v2"
)

type Image = eb.Image

var emptyImg = eb.NewImage(1, 1)

func init() {
	emptyImg.Fill(color.White)
}

func DrawTriangles(scr *Image, trigs Trigs, clr Color) {
	colorR := f32(clr.R) / 255
	colorG := f32(clr.G) / 255
	colorB := f32(clr.B) / 255
	colorA := f32(clr.A) / 255
	verts := make([]eb.Vertex, len(trigs.Verts))
	for i, v := range trigs.Verts {
		verts[i].DstX = f32(v[0])
		verts[i].DstY = f32(windowSize[1] - v[1])
		verts[i].ColorR = colorR
		verts[i].ColorG = colorG
		verts[i].ColorB = colorB
		verts[i].ColorA = colorA
	}
	scr.DrawTriangles(verts, trigs.Inds, emptyImg, &eb.DrawTrianglesOptions{})
}
func DrawRectangle(scr *Image, pos, size v2, clr Color) {
	DrawTriangles(scr, Trigs{[]v2{
		pos,
		pos.Add(V2(size[0], 0)),
		pos.Add(V2(0, size[1])),
		pos.Add(V2(size[0], size[1])),
	}, []u16{0, 1, 2, 1, 2, 3}}, clr)
}
func DrawLine(scr *Image, a, b v2, thickness f64, clr Color) {
	normal := b.Sub(a).Rot90().MagSet(thickness * .5)
	DrawTriangles(scr, Trigs{[]v2{
		a.Sub(normal),
		a.Add(normal),
		b.Sub(normal),
		b.Add(normal),
	}, []u16{0, 1, 2, 1, 2, 3}}, clr)
}
func DrawCircle(scr *Image, pos v2, size f64, points int, clr Color) {
	verts := make([]v2, points)
	for i := range verts {
		verts[i] = rad(m.Tau * f64(i) / f64(points)).Vec2().Mul1(size).Add(pos)
	}
	inds := make([]u16, 0, (points-2)*3)
	for i := 2; i < points; i++ {
		inds = append(inds, 0, u16(i-1), u16(i))
	}
	DrawTriangles(scr, Trigs{verts, inds}, clr)
}
