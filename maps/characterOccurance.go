package maps

func CountCharacterOccurance(characters string) map[string]int {
	count := map[string]int{}
	if len(characters) == 0 {
		return count
	}
	for _, character := range characters {
		//fmt.Println("character:", character)
		key := string(character)

		if count[key] == 0 {
			count[key] = 1
		} else {
			count[key] = count[key] + 1
		}

	}
	return count
}
