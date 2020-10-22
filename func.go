package tool

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateRandom(n int) string {

	// 随机种子
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func MD5(str string) string {

	return strings.ToUpper(Md5(str))
}

func Md5(str string) string {

	h := md5.New()
	h.Write([]byte(str))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func RunBash(s string) (string, error) {
	return RunCommand(s)
}

func RunCommand(s string) (string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	return out.String(), err
}
