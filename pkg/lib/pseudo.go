package passgen

import (
	"math/rand"
	"time"
)

type PseudoGenerate interface{
	 Generate(length uint64) string
}

type randSourcePseudo struct {
	rand *rand.Rand
	sourceString string
}

func NewSourcePseudo(seed int64, source string) (*randSourcePseudo) {
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	if source == "" {
		source = "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm01234567890!"
	}
	r := rand.New(rand.NewSource(seed))
	return &randSourcePseudo{
		rand: r,
		sourceString: source,
	}
}

func (rsp *randSourcePseudo) Generate(length uint64) (string) {
	pass := make([]byte, length)
	for i := range pass {
		pass[i] = rsp.sourceString[rsp.rand.Intn(len(rsp.sourceString))]
	}
	return string(pass)
}