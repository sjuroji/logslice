// Package highlight applies ANSI terminal colour escapes to substrings of
// log lines that match a regular expression.
//
// # Usage
//
//	h, err := highlight.New("ERROR", highlight.ColorRed)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(h.Apply(line))
//
// A nil *Highlighter is safe to use; Apply returns the line unchanged,
// making it convenient to disable highlighting without extra nil checks
// in calling code.
//
// Supported colour names: "yellow" (default), "red", "cyan".
package highlight
