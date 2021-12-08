# zdpgo_rabbitmq
简化Golang操作RabbitMQ的组件库

## 一、快速入门

### 1.1 发布和定义

#### 1.1.1 生产者
```go
package main

import (
	"github.com/zhangdapeng520/zdpgo_rabbitmq"
)

func main(){
	mq:=zdpgo_rabbitmq.NewDefaultRabbitMQ()
	mq.Send("msg", "你好，我是张大鹏！")
}
```

#### 1.1.2 消费者
```go
package main

import (
	"time"

	"github.com/zhangdapeng520/zdpgo_rabbitmq"
)

func main(){
	mq:=zdpgo_rabbitmq.NewDefaultRabbitMQ()
	mq.Receive("msg")
	time.Sleep(3*time.Second)
}
```
