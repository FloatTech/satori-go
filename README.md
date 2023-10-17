# satori-go
[satori](https://satori.js.org/zh-CN/) protocol in golang

## Usage
```go
client := NewClient("12345678", "i am token")

client.Listen(func(event *Event) {
	fmt.Println(event)
})

client.CreateMessage("87654321", "hello world!")
```

## Thanks
- [chronocat](https://github.com/chrononeko/chronocat) - 神秘猫猫
- [zerobot](https://github.com/wdvxdr1123/ZeroBot) - 一个基于onebot协议的机器人Go开发框架