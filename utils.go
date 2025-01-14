/**
 * @Author: DollarKiller
 * @Description: 工具库
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 18:08 2019-09-29
 */
package erguotou

import (
	"bytes"
	"fmt"
	"github.com/dollarkillerx/erguotou/fasthttp"
	"runtime"
	"strconv"
)

func HttpSplice(h1, h2 string) string {
	u := string(h1[len(h1)-1])
	if u == "/" {
		u = h1[:len(h1)-1]
	} else {
		u = h1
	}

	u2 := string(h2[0])
	if u2 == "/" {
		u += h2
	} else {
		u += "/" + h2
	}

	return u
}

type Utils struct{}

func (u *Utils) Get(url string) ([]byte, error) {
	statusCode, body, err := fasthttp.Get(nil, url)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, fmt.Errorf(strconv.Itoa(statusCode))
	}

	return body, nil
}

// 打印堆栈信息
func (u *Utils) PanicTrace(kb int) []byte {
	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine ")
	line := []byte("\n")
	stack := make([]byte, kb<<10) //4KB
	length := runtime.Stack(stack, true)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}
	stack = bytes.TrimRight(stack, "\n")
	return stack
}
