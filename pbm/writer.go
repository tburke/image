package pbm

import (
	"image"
	"io"
)

type Encoder struct {
}

func Encode(w io.Writer, m image.Image) error {
	var e Encoder
	return e.Encode(w, m)
}

func (enc *Encoder) Encode(w io.Writer, m image.Image) error {
	b := m.Bounds()
	img := make([]byte, (b.Max.Y - b.Min.Y) * (b.Max.X - b.Min.X) / 8)
	idx := 0
	bit := 0
	var byt byte
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			_, _, _, a := m.At(x, y).RGBA()
			if a > 0 { byt |= 1 }
			bit++
			if bit > 7 {
				img[idx] = byt
				byt = 0
				bit = 0
				idx++
			} else { byt = byt << 1 }
		}
	}
	_, err := w.Write(img)
	return err
}

