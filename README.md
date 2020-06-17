# Go coding skill 1


Implement `AddProductToPost`, `RemoveProductFromPost` function to make these tests pass.

## Red state(Error)
```
$ go test ./models -v -test.run TestPost$
=== RUN   TestPost
=== RUN   TestPost/AddProductToPost
    TestPost/AddProductToPost: post_test.go:36:
                Error Trace:    post_test.go:36
                Error:          Not equal:
                                expected: 0
                                actual  : 3
                Test:           TestPost/AddProductToPost
--- FAIL: TestPost (0.01s)
    --- FAIL: TestPost/AddProductToPost (0.00s)
...
```

## Green state(Tests Pass)
```
go test ./models -v -test.run TestPost$
=== RUN   TestPost
=== RUN   TestPost/AddProductToPost
=== RUN   TestPost/RemoveProductFromPost
--- PASS: TestPost (0.01s)
    --- PASS: TestPost/AddProductToPost (0.00s)
    --- PASS: TestPost/RemoveProductFromPost (0.00s)
PASS
ok      github.com/Blur-Consulting/prismpop-api/models  8.452s
```

## Code to implement:
`models/post.go` 66 line
```golang
func RemoveProductFromPost(ctx context.Context, postId int64, productIds []int64) error {
	return nil
}

func AddProductToPost(ctx context.Context, postId int64, productIds []int64) ([]Product, error) {
	return nil, nil
}
```


# Go coding skill 2

`longFunc` function执行需要很长时间。
请将代码修改 为持续执行2秒以上时，强制终止此application。


```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	elapsedTime := longFunc()
	fmt.Println(elapsedTime)

}
func longFunc() time.Duration {
	start := time.Now()
	duration := time.Duration(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(5))
	time.Sleep(time.Second * duration)
	return time.Now().Sub(start)
}
```

# Architecture Design

请根据以下要求设计出能够显示page view count的架构。
要求:
1. count必须实时显示。 不得以小时为单位/日为单位统计并展示。
2. 在特定时间内,多次点击也仅为一次。
3. 当用户突然蜂拥而至时，60秒以内应该可以做到scale up。

