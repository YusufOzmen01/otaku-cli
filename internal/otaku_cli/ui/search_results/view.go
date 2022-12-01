package search_results

func (m UI) View() string {
	if m.loading {
		return "Loading details, hang on..."
	}

	//data := "Anime Search Result \n\n"
	//
	//for i, result := range m.Results {
	//	if m.cursor == i {
	//		data += ">"
	//	}
	//
	//	data += fmt.Sprintf(" %s (%s)\n", result.AnimeTitle, result.AnimeUrl)
	//}
	//
	//return data

	return "\n" + m.list.View()
}
