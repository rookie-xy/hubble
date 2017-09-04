package event

// 适配各种不同的消息
type Event interface {
    MakeHeader()
    MakeBody()
    MakeFooter()
    Get()
    Set()
}
