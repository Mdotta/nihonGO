package game

type Option func(*flashcardGameModel)

func WithReverseMode() Option {
	return func(m *flashcardGameModel) {
		// swap keys and values in the pool
		newPool := make(map[string]string)
		for k, v := range m.pool {
			newPool[v] = k
		}
		m.pool = newPool
	}
}
