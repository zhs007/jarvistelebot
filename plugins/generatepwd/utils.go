package plugingeneratepwd

import "math/rand"

const letterNormal = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func genPassword(mode string, length int) (str string) {
	for i := 0; i < length; i++ {
		str += string(letterNormal[rand.Intn(len(letterNormal))])
	}

	return
}
