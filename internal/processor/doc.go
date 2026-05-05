// Package processor provides the core pipeline that connects log filtering
// with output writing.
//
// A Processor reads lines sequentially from any io.Reader (a file, stdin,
// a byte buffer, etc.), passes each line through a [logfilter.Filter], and
// forwards matching lines to an [output.Writer].
//
// Basic usage:
//
//	f := logfilter.New(logfilter.Options{Pattern: "ERROR"})
//	w := output.New(os.Stdout, output.Options{LineNumbers: true})
//	p := processor.New(f, w, processor.Options{MaxLines: 100})
//
//	n, err := p.Process(file)
//
// MaxLines in Options caps the number of matching lines that are written.
// Setting it to zero disables the cap.
package processor
