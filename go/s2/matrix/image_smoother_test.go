package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageSmoother(t *testing.T) {
	assert.Equal(t,
		[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		imageSmoother([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))

	assert.Equal(t,
		[][]int{{137, 141, 137}, {141, 138, 141}, {137, 141, 137}},
		imageSmoother([][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}))
}

func imageSmoother(img [][]int) [][]int {
	if len(img) == 0 {
		return img
	}

	h := len(img) - 1
	w := len(img[0]) - 1

	s := make([][]int, len(img))
	for i := range s {
		s[i] = make([]int, w+1)
	}

	p1x := 0
	p1y := 0
	p2x := 0
	p2y := 0

	for y, row := range img {
		for x := range row {
			p1y = y - 1
			if p1y < 0 {
				p1y = 0
			}

			p1x = x - 1
			if p1x < 0 {
				p1x = 0
			}

			p2x = x + 1
			if p2x > w {
				p2x = w
			}
			p2y = y + 1
			if p2y > h {
				p2y = h
			}

			sum, count := 0, 0
			for x := p1x; x <= p2x; x++ {
				for y := p1y; y <= p2y; y++ {
					sum += img[y][x]
					count++
				}
			}

			s[y][x] = sum / count
		}
	}

	return s

}
