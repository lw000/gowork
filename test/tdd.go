package main

type Server struct {
	Name string
}

type Servers []Server

func ListServer() Servers {
	return []Server{
		{Name: "Server1"},
		{Name: "Server2"},
		{Name: "Server3"},
		{Name: "Server4"},
		{Name: "Server5"},
		{Name: "Server6"},
	}
}

func (s Servers) Filter(name string) Servers {
	filtered := make(Servers, 0)

	return filtered
}

func (s Servers) Check(name string) bool {

	return true
}

func (s Servers) Add(name string) {
	// v := Server{Name: name}
	// s = append(s, &v)
}

func (s Servers) Remove(name string) {

}

func (s Servers) Len() int {
	return len(s)
}
