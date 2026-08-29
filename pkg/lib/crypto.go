package passgen

import "crypto/rand"


type randSourceCrypto struct {
	keyCryptoRandom []byte
	sourceString string
	length uint64
}

func NewSourceCrypto(length uint64, source string) (randSourceCrypto) {
	key := make([]byte, length)
	rand.Read(key)
	if source == "" {
		source = "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm01234567890!"
	}
	return randSourceCrypto{
		keyCryptoRandom: key,
		length: length,
		sourceString: source,
	}
}

func (rsc *randSourceCrypto) Generate() (string) {
	pass := make([]byte, rsc.length)
	for i, b := range rsc.keyCryptoRandom {
		pass[i] = rsc.sourceString[b % byte(len(rsc.sourceString))]
	}
	return string(pass)
}