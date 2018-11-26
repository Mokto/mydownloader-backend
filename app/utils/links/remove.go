package links


func Remove(id string) error {
	links := GetAll()
	for i, link := range links {
		if link.ID == id {
			links = append(links[:i], links[i+1:]...)
			break
		}
	}
	return Save(links)
}
