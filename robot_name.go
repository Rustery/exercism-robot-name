package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

// Robot main structure.
type Robot struct {
	name string
}

const abc = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var usedCount int
var possibleNames []string

func init() {
	for _, v := range abc {
		for _, vv := range abc {
			for i := 0; i < 1000; i++ {
				possibleNames = append(possibleNames, fmt.Sprintf("%c%c%03d", v, vv, i))
			}
		}
	}
	rand.Shuffle(len(possibleNames), func(i, j int) {
		possibleNames[i], possibleNames[j] = possibleNames[j], possibleNames[i]
	})
}

func (r *Robot) Name() (string, error) {
	if len(r.name) > 0 {
		return r.name, nil
	}
	for i := usedCount; i < len(possibleNames); i++ {
		r.name = possibleNames[i]
		usedCount++
		return r.name, nil
	}
	return "", errors.New("out of namespace range")
}

func (r *Robot) Reset() {
	r.name = ""
}
