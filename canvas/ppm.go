package canvas

import (
	"fmt"
	"io"
	"math"
	"os"
)

func (c *Canvas) SavePpm(w io.Writer) {
	fmt.Fprintf(w, "P3\n%d %d\n255\n", c.width, c.height)

	for y := uint32(0); y < c.height; y++ {
		for x := uint32(0); x < c.width; x++ {
			idx := c.width*y + x
			color := c.pixels[idx]
			r := int(math.Min(math.Max(color.R()*255, 0), 255) + 0.5)
			g := int(math.Min(math.Max(color.G()*255, 0), 255) + 0.5)
			b := int(math.Min(math.Max(color.B()*255, 0), 255) + 0.5)
			fmt.Fprintf(w, "%d %d %d ", r, g, b)
		}
		fmt.Fprintf(w, "\n")
	}
}

func (c *Canvas) SavePpmToFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	c.SavePpm(f)

	return f.Close()

}
