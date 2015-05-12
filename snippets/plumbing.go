// START 1 OMIT
type Reader interface {
	Read(p []byte) (n int, err error)
}
// END 1 OMIT
// START 2 OMIT
// json
func NewDecoder(r io.Reader) *Decoder // HL

// gzip
func NewReader(r io.Reader) (*Reader, error) // HL
// END 2 OMIT

// START 3 OMIT
response, _ := client.Do(request)
raw, _ := gzip.NewReader(response.Body)
json.NewDecoder(raw).Decode(&data)
// END 3 OMIT




// file instead of request
file, _ := os.Open(filename)
raw, _ := gzip.NewReader(file)
err := json.NewDecoder(raw).Decode(&v)

// sources of readers
*os.File implements io.Reader

files, http requests, byte buffers, network connections

