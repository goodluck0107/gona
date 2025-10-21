package channel

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/goodluck0107/gona/utils"
)

func TestChannelWrite(t *testing.T) {
	rand.Seed(time.Now().Unix())
	data := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	doWrite(data)
	fmt.Println("sendOver")
}
func channelWrite(data []byte) (int, error) {
	totalLength := len(data)
	sendLength := rand.Intn(totalLength) + 1
	fmt.Println("channelWrite sendLength:", sendLength, " totalLength:", totalLength)
	fmt.Println("writeData:", data[:sendLength])
	return sendLength, nil
}
func doWrite(data []byte) (err error) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			err = errors.New(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	var goal int = len(data)
	var hasWriteLength int = 0
	for {
		i, err1 := channelWrite(data)
		if err1 != nil {
			err = err1
			return
		}
		if i > 0 {
			data = data[i:]
			hasWriteLength = hasWriteLength + i
		}
		if hasWriteLength >= goal {
			return
		}
	}
}
