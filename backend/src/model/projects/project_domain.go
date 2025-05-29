package projects

type Project struct {
	iD      string
	name    string
	idAsana string
}

func (p *Project) GetID() string {
	return p.iD
}
func (p *Project) GetName() string {
	return p.name
}
func (p *Project) GetIdAsana() string {
	return p.idAsana
}
func (p *Project) SetID(id string) {
	p.iD = id
}
