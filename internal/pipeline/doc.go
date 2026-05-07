// Package pipeline assembles the filereader, logfilter, processor, and output
// components into a cohesive unit that can process a single log file end-to-end.
//
// # Usage
//
//	s := stats.New()
//	pl := pipeline.New(cfg, os.Stdout, s)
//	if err := pl.Run("/var/log/app.log"); err != nil {
//		log.Fatal(err)
//	}
//
// The Pipeline type is intentionally stateless between Run calls so that the
// same instance can be reused across multiple files within a single invocation.
package pipeline
