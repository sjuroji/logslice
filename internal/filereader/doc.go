// Package filereader provides a thin abstraction over file and stream
// reading for logslice.
//
// It exposes a Reader type that wraps bufio.Scanner with a larger
// default buffer suitable for wide log lines, and supports both
// file-backed and io.Reader-backed sources.
//
// Typical usage:
//
//	r, err := filereader.New("/var/log/app.log")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer r.Close()
//
//	for r.Scan() {
//		fmt.Println(r.Text())
//	}
//	if err := r.Err(); err != nil {
//		log.Fatal(err)
//	}
package filereader
