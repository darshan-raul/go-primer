package maps

type Dictionary map[string]string

func Search(dictionary Dictionary, searchstr string) string{
	return dictionary[searchstr]
}