package coinmarketcap

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	Platform          = "Linux i686"
	InnerHeight       = 355
	OuterHeight       = "355"
	PuzzlePieceWidth  = 44
	PuzzlePieceHeight = 44
	ImageWidth        = "310"
)

func (p *Payload) FillPayload(solvedX int) {
	// ENVIRONMENT
	p.Ev.Wd = genOddNum()
	p.Ev.Im = genEvenNum()
	p.Ev.Prde = fmt.Sprintf("%v,%v,%v,%v", genOddNum(), genOddNum(), genOddNum(), genOddNum())
	p.Ev.Brla = genOddNum()
	p.Ev.Pl = Platform
	p.Ev.Wiinhe = InnerHeight
	p.Ev.Wiouhe = OuterHeight

	MACT, finalDist := generateMACT(solvedX, 290)
	p.Be.El = MACT
	p.Be.Ec.Ts = 1
	p.Be.Ec.Tm = len(MACT) - 2
	p.Be.Ec.Te = 1

	// PUZZLE PIECE
	p.Be.Th.El = []string{}
	p.Be.Th.Si.W = PuzzlePieceWidth
	p.Be.Th.Si.H = PuzzlePieceHeight

	// FINAL X WHERE THE SLIDER LANDS, MINUS THE OFFSET
	p.Dist = finalDist

	// IMAGE WIDTH, HARD CODED
	p.ImageWidth = ImageWidth
}

func generateMACT(xEnd, yDom int) ([]string, int) {
	xStart := rand.Intn(40) + 20
	xStartCopy := xStart + 1
	xEndMinusOffset := xEnd - 60
	yStart := yDom + 1
	eventsNeeded := int(math.Ceil(float64(xEnd) / 6.0))
	timestampStart := time.Now().UnixMilli()
	totalTimeUsed := 0
	output := []string{}
	for i := 0; i < eventsNeeded; i++ {
		yStart = newY(yStart)
		xStart = newX(xStart)
		timeTaken := getTimeTaken(i)

		totalTimeUsed += timeTaken

		output = append(output, fmt.Sprintf("|tm|%v,%v|%v|1", xStart, yStart, timeTaken))

		if xStart >= xEnd {
			break
		}
	}
	timeTaken := getTimeTaken(999)
	output = append(output, fmt.Sprintf("|te||%v|1", timeTaken))
	totalTimeUsed += timeTaken
	time.Sleep(time.Duration(totalTimeUsed*2) * time.Millisecond)
	finalOutput := []string{fmt.Sprintf("|ts|%v,%v|%v|1", xStartCopy, yDom, timestampStart)}
	for _, o := range output {
		finalOutput = append(finalOutput, o)
	}
	return finalOutput, xEndMinusOffset
}

func getTimeTaken(i int) int {
	if i == 0 {
		return rand.Intn(20) + 80
	}
	if i == 999 {
		return rand.Intn(50) + 200
	}
	return rand.Intn(7) + 13
}

func newX(xStart int) int {
	amtToAdd := rand.Intn(6) + 3
	return xStart + amtToAdd
}

func newY(yStart int) int {
	// get if change y
	changeY := randBool()
	// get if change up or down
	var changeYUp bool
	if changeY {
		changeYUp = randBool()
	}
	// get how much to change
	var changeYAmount int
	if changeY {
		changeYAmount = rand.Intn(3)
		if changeYUp {
			yStart += changeYAmount
		} else {
			yStart -= changeYAmount
		}
	}
	return yStart
}

func randBool() bool {
	return rand.Intn(2) == 1
}

func genOddNum() int {
	return (rand.Intn(50) * 2) + 1
}

func genEvenNum() int {
	return rand.Intn(50) * 2
}
