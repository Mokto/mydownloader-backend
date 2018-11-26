package downloads


func Remove(id string) error {
	downloads := GetAll()
	for i, download := range downloads {
		if download.ID == id {
			downloads = append(downloads[:i], downloads[i+1:]...)
			break
		}
	}
	return Save(downloads)
}
