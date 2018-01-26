package metier

import "github.com/nargil07/nargilplayer/entity"

type MetierMusiques struct {
	musiques []*entity.Musique
	size     int
}

func (m *MetierMusiques) AddMusique(musique *entity.Musique) {
	m.musiques = append(m.musiques, musique)
	m.size++
}

func (m *MetierMusiques) GetMusiques() []*entity.Musique {
	return m.musiques
}
func (m *MetierMusiques) GetSize() int {
	return m.size
}
