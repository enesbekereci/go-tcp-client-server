package main

type User struct {
	name          string
	remoteaddress string
	starttime     string
	state         bool
	filepath      string
	logqueue      LogQueue
}

type MessageType int

const (
	Red   = 0
	Green = 1
	Blue  = 2
	White = 3
	Black = 4
)
const (
	Typing     MessageType = 0
	Screenshot MessageType = 1
	Info       MessageType = 2
	Other      MessageType = 3
)
const queuelimit int = 10

type NetMessage struct {
	//mt   MessageType
	//data string
}

type LogMessage struct {
	date    int
	message string
	next_m  *LogMessage
	prev_m  *LogMessage
}

type LogQueue struct {
	start   *LogMessage
	end     *LogMessage
	current *LogMessage
	count   int
}

func (l *LogQueue) Add(m *LogMessage) {
	if l.start == nil {
		l.start = m
		l.end = m
		l.current = m
	} else {
		l.end.next_m = m
		m.prev_m = l.end
		l.end = m
		l.count--
	}
	if l.count == 0 {
		l.start = l.start.next_m
		l.start.prev_m = nil
		l.count++
	}
}

func (l *LogQueue) Get() *LogMessage {
	r := l.current
	if l.current != nil {
		l.current = l.current.next_m
	}
	return r
}

func (l *LogQueue) ResetCurrent() {
	l.current = l.start
}
