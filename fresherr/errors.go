package fresherr

import "math/rand/v2"

var preambles []string = []string{
	"Well, there's your problem",
	"Ker-blam",
	"Womp womp",
	"*ambulance noises*",
}

func GetFresh() string {
	return preambles[rand.IntN(len(preambles))]
}
