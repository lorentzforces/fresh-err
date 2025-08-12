package fresherr

import "math/rand/v2"

var preambles []string = []string{
	"Well, there's your problem",
	"Ker-blam",
	"Womp womp",
	"*ambulance noises*",
	"terrible error",
	"water u doin bud",
	"o no",
	"aw shucks",
	"stop right there criminal scum",
	"something's fucky",
	"biffed it",
	"🤢",
	"problematic uncle vibes",
	"wrecked",
	"accidentally the software",
	"absolute cheeks",
	"honestly not our best work",
}

func GetFresh() string {
	return preambles[rand.IntN(len(preambles))]
}
