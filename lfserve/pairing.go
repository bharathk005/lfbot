package lfserve

import "log"

// if pair exists forward the message to that paired chat
// else create a new pair from Free
// else send a wait message

func GetPair(id int64) int64 {
	m = GetMap()
	if m == nil {
		return -1
	}
	v, ok := m.Pairs[id]
	if !ok {
		return -1
	} else {
		return v
	}
}

func NewPair(id int64) (int64, int64) {
	m = GetMap()
	if m == nil {
		log.Printf("map not found")
		return -1, -1
	}
	len := m.Free.size()
	v, ok := m.Pairs[id]
	if !ok {
		if len > 0 {
			k := m.Free.remove()
			if k == id {
				if len == 1 {
					m.Free.add(id)
					return -1, -1
				} else if len > 1 {
					k = m.Free.remove()
					m.Pairs[id] = k
					m.Pairs[k] = id
					return k, -1
				}
			} else {
				m.Pairs[id] = k
				m.Pairs[k] = id
				return k, -1
			}
		} else {
			m.Free.add(id)
			return -1, -1
		}
	} else {
		delete(m.Pairs, v)
		delete(m.Pairs, id)
		if len > 0 {
			k := m.Free.remove()
			m.Pairs[id] = k
			m.Pairs[k] = id
			m.Free.add(v)
			return k, v

		} else {
			m.Free.add(v)
			m.Free.add(id)
			return -1, -1
		}
	}

	return -1, -1

}
