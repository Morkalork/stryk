package core

type LevelMap struct {
	Id        int                   `json:id`
	Count     int                   `json:count`
	Colors    []string              `json:colors`
	Bricks    []Brick               `json:bricks`
	Seconds   int                   `json:seconds`
	Listeners map[string][]chan int `json:"-"`
}

func (m *LevelMap) AddListener(name string, ch chan int) {
	if m.Listeners == nil {
		m.Listeners = make(map[string][]chan int)
	}

	_, ok := m.Listeners[name]
	if ok {
		m.Listeners[name] = append(m.Listeners[name], ch)
	} else {
		m.Listeners[name] = []chan int{ch}
	}
}

func (m *LevelMap) RemoveFinisher(name string, ch chan int) {
	_, ok := m.Listeners[name]
	if ok {
		for i := range m.Listeners[name] {
			if m.Listeners[name][i] == ch {
				m.Listeners[name] = append(m.Listeners[name][:i], m.Listeners[name][i+1:]...)
			}
		}
	}
}

func (m *LevelMap) Emit(name string, finishedLevel int) {
	_, ok := m.Listeners[name]
	if ok {
		for _, handler := range m.Listeners[name] {
			go func(handler chan int) {
				handler <- finishedLevel
			}(handler)
		}
	}
}
