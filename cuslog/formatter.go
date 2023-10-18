package cuslog

type Formatter interface {
	// Format Maybe in async goroutine
	// Please write the result to buffer
	Format(entry *Entry) error
}
