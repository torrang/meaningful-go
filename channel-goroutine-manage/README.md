# channel-goroutine-manage
채널을 이용한 고루틴 관리 패턴

## 관리 방식
1. sync.WaitGroup, channel 변수 생성
2. 고루틴(async_task)을 생성할 때마다 2개 변수를 넘기고 WaitGroup.Add(1) 호출
3. 외부 요인(코드에서는 10초 대기 후 취소 함수를 호출하는 고루틴)에 의해 채널이 닫힘
4. 고루틴(async_task)에서 실행 중 채널이 닫혀 해당 case를 타게됨
5. 고루틴(async_task)에서 case 진행 후 return 직전 WaitGroup.Done() 함수를 호출 후 종료
6. 메인 함수에서는 생성한 WaitGroup을 이용해 대기하던 중 WaitGroup의 카운터가 0이 되어 대기 종료

## 특징
* 고루틴에서 인자로 받은 채널은 읽기 전용 채널로 받아서 처리