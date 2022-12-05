package EduTen

import (
	"fmt"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	m "github.com/IgneousRed/gomisc"
	text "github.com/hajimehoshi/ebiten/v2/text"
)

var fontId = 0

type Font struct {
	id   int
	data []byte
}

func FontNew(path string) (Font, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Font{}, err
	}
	result := Font{fontId, data}
	fontId++
	return result, nil
}

var fonts = map[string]font.Face{}

func DrawText(f Font, size float64, pos m.Vec2F, txt string, clr Color) {
	str := fmt.Sprint(f.id, size)
	face, ok := fonts[str]
	if !ok {
		ttf, _ := opentype.Parse(f.data)
		face, _ = opentype.NewFace(ttf, &opentype.FaceOptions{
			Size:    float64(size),
			DPI:     72,
			Hinting: font.HintingFull,
		})
		fonts[str] = face
	}
	text.Draw(Screen, txt, face, int(pos.X), int(windowSizeY64-pos.Y), clr.Color())
}
