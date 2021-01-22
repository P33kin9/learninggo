package memo_test

import (
	memo "cf/ch9/memo1"
	"cf/ch9/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe! Test fails.
func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}

/*
//!+output
$ go test -v cf/ch9/memo1
http://goproxy.io, 524.234633ms, 33704 bytes
http://gopl.io, 4.409090448s, 4154 bytes
http://www.so.com, 268.550228ms, 132173 bytes
http://www.baidu.com, 129.720646ms, 290612 bytes
https://www.sohu.com, 66.744682ms, 215702 bytes
https://www.sogou.com, 119.219428ms, 15146 bytes
--- PASS: Test (5.52s)
=== RUN   TestConcurrent
https://www.sohu.com, 13.853726ms, 215702 bytes
https://www.sogou.com, 80.10402ms, 15146 bytes
http://www.baidu.com, 145.908217ms, 290617 bytes
http://goproxy.io, 190.338001ms, 33704 bytes
http://www.so.com, 257.23351ms, 132169 bytes
http://gopl.io, 562.861557ms, 4154 bytes
--- PASS: TestConcurrent (0.56s)
PASS
ok  	cf/ch9/memo1	6.093s
//!-output
*/
/*
//!+race
$ got -run=TestConcurrent -race -v
=== RUN   TestConcurrent
WARNING: DATA RACE
	...
Write at 0x00c000076d80 by goroutine 13:
  runtime.mapassign_faststr()
      /usr/local/Cellar/go/1.15.6/libexec/src/runtime/map_faststr.go:202 +0x0
  cf/ch9/memo1.(*Memo).Get()
      /Users/p33king/learngo/src/Tgopl/ch9/memo1/memo.go:28 +0x1cd
  cf/ch9/memotest.Concurrent.func1()
      /Users/p33king/learngo/src/Tgopl/ch9/memotest/memotest.go:86 +0xe5

Previous write at 0x00c000076d80 by goroutine 29:
  runtime.mapassign_faststr()
      /usr/local/Cellar/go/1.15.6/libexec/src/runtime/map_faststr.go:202 +0x0
  cf/ch9/memo1.(*Memo).Get()
      /Users/p33king/learngo/src/Tgopl/ch9/memo1/memo.go:28 +0x1cd
  cf/ch9/memotest.Concurrent.func1()
      /Users/p33king/learngo/src/Tgopl/ch9/memotest/memotest.go:86 +0xe5
      ...
FAIL
exit status 1
FAIL	cf/ch9/memo1	4.374s
//!-race
*/
