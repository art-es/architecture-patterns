package eventbus

type Subscriber struct {
	c     <-chan source
	b     *Bus
	event string
	pos   int
}

func (s *Subscriber) Receive() (interface{}, bool) {
	msg, ok := <-s.c
	return msg, ok
}

func (s *Subscriber) Unsubscribe() {
	s.b.removeEvent(s.event, s.pos)
}
