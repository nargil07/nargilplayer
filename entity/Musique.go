package entity

type Musique struct{
	Length int
	Path string
	Name string
}


func (m *Musique) SetSize(size int){
	m.Length = size
}

func (m *Musique) SetPath(path string){
	m.Path = path
}

func (m *Musique) SetName(name string){
	m.Name = name
}