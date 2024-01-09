# goroutine-closure-for-loop
for loop 안에서 고루틴 사용 방법 (Go 1.21 이하 기준)

## 문제점
Go에서는 for loop 안에서 다음과 같이 고루틴을 사용할 수 있는데 결과에서 보다시피 의도한대로 표시하지 않는다.

이는 Go 1.21 이하 버전에서는 다음과 같이 실행되기 때문이다.

1. for loop가 처리되면서 i 변수가 1 증가
2. i 값을 프린트하는 클로저 고루틴 호출
3. 고루틴이 실제로 실행되기 전 for loop가 모두 실행됨
4. 고루틴이 현재 설정된 i 변수 값인 6을 표시 (이미 for loop가 모두 실행되었기 때문에)


```go
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i <= 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(1 * time.Second)
}

// 결과
// 6
// 6
// 6
// 6
// 6
// 6
```

## 해결 방안
Go 1.21 이하 버전에서 해결하기 위해서는 다음과 같은 코드 중 한가지 방식을 사용하면 된다.

어떤 방식을 사용하던 고루틴으로 실행하기 때문에 결과가 랜덤으로 표시되며 변수를 할당하는 방식이기 때문에 메모리를 약간 더 사용한다.

### 1. 변수 재할당
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i <= 5; i++ {
        j := i
		go func() {
			fmt.Println(j)
		}()
	}

	time.Sleep(1 * time.Second)
}
```

### 2. Pass by value
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i <= 5; i++ {
		go func(v int) {
			fmt.Println(v)
		}(i)
	}

	time.Sleep(1 * time.Second)
}
```
