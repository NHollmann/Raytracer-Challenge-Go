package canvas

import "nicolashollmann.de/raytracer-challange/color"

type Canvas struct {
	width  uint32
	height uint32
	pixels []color.Color
}

func New(width, height uint32) Canvas {
	return Canvas{
		width:  width,
		height: height,
		pixels: make([]color.Color, width*height),
	}
}

func (c *Canvas) Width() uint32 {
	return c.width
}

func (c *Canvas) Height() uint32 {
	return c.height
}

func (c *Canvas) PixelAt(x, y uint32) color.Color {
	idx := c.width*y + x
	return c.pixels[idx]
}

func (c *Canvas) SetPixel(x, y uint32, p color.Color) {
	idx := c.width*y + x
	c.pixels[idx] = p
}
