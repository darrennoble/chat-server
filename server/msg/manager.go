package msg

import (
	"sync"
)

const defaultChannelCapacity = 10

type Callback func(*Msg)

type Manager struct {
	users           map[string]map[string]bool //map[channel]map[user]true
	msgs            map[string]*MsgList
	ChannelCapacity int
	callbacks       []Callback
	lock            sync.Mutex
}

func (m *Manager) AddChannel(channel string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.addChan(channel)
}

func (m *Manager) addChan(channel string) {
	if _, ok := m.msgs[channel]; !ok {
		m.users[channel] = map[string]bool{}
		m.msgs[channel] = NewMsgList(m.getChannelCapacity())
	}
}

func (m *Manager) AddUser(user, channel string) bool {
	m.lock.Lock()
	defer m.lock.Unlock()

	usrMap, ok := m.users[channel]
	if !ok {
		m.addChan(channel)
		usrMap = m.users[channel]
	}

	_, ok = usrMap[user]
	if ok {
		return false
	}

	usrMap[user] = true

	return true
}

func (m *Manager) RemoveUser(user, channel string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if usrMap, ok := m.users[channel]; ok {
		delete(usrMap, user)
	}
}

func (m Manager) UserInChannel(user, channel string) bool {
	m.lock.Lock()
	defer m.lock.Unlock()

	usrMap, ok := m.users[channel]
	if !ok {
		return false
	}

	_, ok = usrMap[user]
	return ok
}

func (m Manager) getChannelCapacity() int {
	if m.ChannelCapacity <= 0 {
		return defaultChannelCapacity
	}

	return m.ChannelCapacity
}

func (m *Manager) RegisterCallback(c Callback) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.callbacks = append(m.callbacks, c)
}

func (m *Manager) SendMsg(channel string, msg *Msg) {
	m.lock.Lock()
	defer m.lock.Unlock()

	ml, ok := m.msgs[channel]
	if !ok {
		return
	}

	ml.Add(msg)

	for _, c := range m.callbacks {
		go c(msg)
	}
}

func (m *Manager) GetHistory(channel string) []*Msg {
	m.lock.Lock()
	defer m.lock.Unlock()

	ml, ok := m.msgs[channel]
	if !ok {
		return nil
	}

	cap := m.getChannelCapacity()
	msgs := make([]*Msg, 0, cap)

	for i := 0; i < cap; i++ {
		m := ml.Get(i)
		if m == nil {
			break
		}
		msgs = append(msgs, m)
	}

	return msgs
}
