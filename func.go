package tool

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"math/rand"
	"net/url"
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

func CacheSet(key, val string) {
	RunCommand("curl http://161.117.89.225:9097/cache?" + key + "=" + url.QueryEscape(val))
}

func CacheGet(key string) string {
	val, _ := RunCommand("curl http://161.117.89.225:9097/cache?" + key + "=")
	if val == "" {
		return ""
	}

	val, _ = url.QueryUnescape(val)

	return val
}

func RunCommand(s string) (string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	return out.String(), err
}

func InsertSort(arr []string) {
	if len(arr) <= 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

func PasteToPbCopy(str string) {
	RunCommand(fmt.Sprintf("echo '%s'|pbcopy", str))
}
