package goutil

import (
	"math/rand"
	"strings"
	"time"
)

const (
	AlphaCharset       = "abcdefghijklmnopqrstuvwxyz"
	AlphaUCharset      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NumCharset         = "0123456789"
	AlphaNumCharset    = AlphaCharset + NumCharset
	AlphaUNumCharset   = AlphaUCharset + NumCharset
	AlphaAllNumCharset = AlphaCharset + AlphaUNumCharset
	HexadecimalCharset = NumCharset + "abcdef"
	SymbolCharset      = "~!@#$%^&*()_-+=[{}]|;:,<.>?"
)

type RandomString struct {
	charset    string
	randomizer *rand.Rand
}

func NewRandomString(charset string) *RandomString {
	useCharset := AlphaNumCharset

	charset = strings.TrimSpace(charset)
	if len(charset) > 0 {
		useCharset = charset
	}

	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	return &RandomString{charset: useCharset, randomizer: randomizer}
}

func (r *RandomString) SetCharset(charset string) {
	r.charset = charset
}

func (r *RandomString) WithCharset(charset string) *RandomString {
	r.SetCharset(charset)

	return r
}

func (r *RandomString) SetRandomizer(randomizer *rand.Rand) {
	r.randomizer = randomizer
}

func (r *RandomString) WithRandomizer(randomizer *rand.Rand) *RandomString {
	r.SetRandomizer(randomizer)

	return r
}

func (r *RandomString) Generate(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = r.charset[r.randomizer.Intn(len(r.charset))]
	}

	return string(b)
}

func (r *RandomString) GenerateRange(min, max int) string {
	if min < 1 {
		min = 1
	}

	if max < min {
		max = min
	}

	max = max - min

	length := min + rand.Intn(max)

	return r.Generate(length)
}
