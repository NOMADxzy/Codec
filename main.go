package main

import "image"

func main() {
	w := 400
	h := 400
	var nal [][]byte

	c, _ := NewH264Encoder(w, h, image.YCbCrSubsampleRatio420)
	nal = append(nal, c.Header)

	for i := 0; i < 60; i++ {
		img := image.NewYCbCr(image.Rect(0, 0, w, h), image.YCbCrSubsampleRatio420)
		p, _ := c.Encode(img)
		if len(p.Data) > 0 {
			nal = append(nal, p.Data)
		}
	}
	for {
		// flush encoder
		p, err := c.Encode(nil)
		if err != nil {
			break
		}
		nal = append(nal, p.Data)
	}
	return
}
