package eventbus

import "sync"

type source interface{}

type Bus struct {
	mu       sync.RWMutex
	channels map[string][]chan<- source
}

func (b *Bus) Publish(event string, source interface{}) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if cc, ok := b.channels[event]; ok {
		for _, c := range cc {
			c <- source
		}
	}
}

func (b *Bus) Subscribe(event string, size int) *Subscriber {
	c := make(chan source, size)
	pos := b.addEvent(event, c)
	return &Subscriber{c, b, event, pos}
}

func (b *Bus) addEvent(event string, c chan<- source) int {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.channels[event]; ok {
		b.channels[event] = append(b.channels[event], c)
		return len(b.channels) - 1
	}

	b.channels[event] = []chan<- source{c}
	return 0
}

func (b *Bus) removeEvent(event string, pos int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if cc, ok := b.channels[event]; ok {
		if pos < len(cc) {
			close(cc[pos])
			cc = append(cc[:pos], cc[pos+1:]...)
			b.channels[event] = cc
		}
	}
}

func New() *Bus {
	return &Bus{channels: make(map[string][]chan<- source)}
}
