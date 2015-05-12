// START 1 OMIT
error
Stringer
Handler
ResponseWriter
sort.Interface
Reader, Writer, ReadWriter,
Marshaler, Umarshaler
TB (tests, benchmarks)
// END 1 OMIT

// START 2 OMIT
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader // HL
	Writer // HL
}
// END 2 OMIT