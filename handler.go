package function

import (
	"bytes"
	"crypto/md5"
	"image"
	"image/color"
	"image/png"

	"github.com/llgcode/draw2d/draw2dimg"
)

// Point holds the co-ordinated in 2D
type Point struct {
	x, y int
}

// DrawingPoint holds topLeft and bottomRight of viewport
type DrawingPoint struct {
	topLeft     Point
	bottomRight Point
}

// GridPoint represents a grid in the identicon
type GridPoint struct {
	value byte
	index int
}

// Identicon stores the data necessary to draw identicon
type Identicon struct {
	name       string
	hash       [16]byte
	color      [3]byte
	grid       []byte
	gridPoints []GridPoint
	pixelMap   []DrawingPoint
}

// Apply is a consumes and produces of Identicon
type Apply func(Identicon) Identicon

// Handle a serverless request
func Handle(req []byte) string {
	identicon := hashInput(req)
	identicon = pipe(identicon, pickColor, buildGrid, filterOddSquares, buildPixelMap)

	img := drawRectangle(identicon)
	buf := bytes.NewBufferString("")
	png.Encode(buf, img)
	return buf.String()
}

func pipe(identicon Identicon, funcs ...Apply) Identicon {
	for _, applyer := range funcs {
		identicon = applyer(identicon)
	}
	return identicon
}

func hashInput(input []byte) Identicon {
	checkSum := md5.Sum(input)
	return Identicon{
		name: string(input),
		hash: checkSum,
	}
}

func pickColor(identicon Identicon) Identicon {
	rgb := [3]byte{}
	copy(rgb[:], identicon.hash[:3])
	identicon.color = rgb
	return identicon
}

func buildGrid(identicon Identicon) Identicon {
	grid := []byte{}
	for i := 0; i < len(identicon.hash) && i+3 <= len(identicon.hash)-1; i += 3 {
		chunk := make([]byte, 5)
		copy(chunk, identicon.hash[i:i+3])
		chunk[3] = chunk[1]
		chunk[4] = chunk[0]
		grid = append(grid, chunk...)

	}
	identicon.grid = grid
	return identicon
}

func filterOddSquares(identicon Identicon) Identicon {
	grid := []GridPoint{}
	for i, code := range identicon.grid {
		if code%2 == 0 {
			point := GridPoint{
				value: code,
				index: i,
			}
			grid = append(grid, point)
		}
	}
	identicon.gridPoints = grid
	return identicon
}

func buildPixelMap(identicon Identicon) Identicon {
	drawingPoints := []DrawingPoint{}

	pixelFunc := func(p GridPoint) DrawingPoint {
		horizontal := (p.index % 5) * 50
		vertical := (p.index / 5) * 50
		topLeft := Point{horizontal, vertical}
		bottomRight := Point{horizontal + 50, vertical + 50}

		return DrawingPoint{
			topLeft,
			bottomRight,
		}
	}

	for _, gridPoint := range identicon.gridPoints {
		drawingPoints = append(drawingPoints, pixelFunc(gridPoint))
	}
	identicon.pixelMap = drawingPoints
	return identicon
}

func drawRectangle(identicon Identicon) image.Image {
	var img = image.NewRGBA(image.Rect(0, 0, 250, 250))
	col := color.RGBA{identicon.color[0], identicon.color[1], identicon.color[2], 255}

	for _, pixel := range identicon.pixelMap {
		rect(img, col, float64(pixel.topLeft.x), float64(pixel.topLeft.y), float64(pixel.bottomRight.x), float64(pixel.bottomRight.y))
	}
	return img
}

func rect(img *image.RGBA, col color.Color, x1, y1, x2, y2 float64) {
	gc := draw2dimg.NewGraphicContext(img)
	gc.SetFillColor(col)
	gc.MoveTo(x1, y1)
	gc.LineTo(x1, y1)
	gc.LineTo(x1, y2)
	gc.MoveTo(x2, y1)
	gc.LineTo(x2, y1)
	gc.LineTo(x2, y2)
	gc.SetLineWidth(0)
	gc.FillStroke()
}
