package test3

const (
	french  = "French"
	spanish = "Spanish"
    englishHelloPrefix = "Hello, "
    spanishHelloPrefix = "Hola, "
    frenchHelloPrefix = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// The function name starts with a lowercase letter. 
// In Go, public functions start with a capital letter and private ones start with a lowercase. 
// We don't want the internals of our algorithm to be exposed to the world, so we made this function private
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	//default in the switch case will 
	// be branched to if none of the other case statements match.
	default:
		prefix = englishHelloPrefix
	}
	return
}
func main() {
	
}