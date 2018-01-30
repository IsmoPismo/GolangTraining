package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	bs, _ := ioutil.ReadAll(res.Body)
	// sc := bufio.NewScanner(res.Body) => takes a Reader Interface =>
	// func NewScanner(r io.Reader) *Scanner {
	// 	return &Scanner{
	// 		r:            r,
	// 		split:        ScanLines,
	// 		maxTokenSize: MaxScanTokenSize,
	// 	}
	// }
	str := string(bs)
	fmt.Println(str)

}

//  type Response struct {
// 	Status     string // e.g. "200 OK"
// 	StatusCode int    // e.g. 200
// 	Proto      string // e.g. "HTTP/1.0"
// 	ProtoMajor int    // e.g. 1
// 	ProtoMinor int    // e.g. 0
//
// 	// Header maps header keys to values. If the response had multiple
// 	// headers with the same key, they may be concatenated, with comma
// 	// delimiters.  (Section 4.2 of RFC 2616 requires that multiple headers
// 	// be semantically equivalent to a comma-delimited sequence.) When
// 	// Header values are duplicated by other fields in this struct (e.g.,
// 	// ContentLength, TransferEncoding, Trailer), the field values are
// 	// authoritative.
// 	//
// 	// Keys in the map are canonicalized (see CanonicalHeaderKey).
// 	Header Header
//
// 	// Body represents the response body.
// 	//
// 	// The http Client and Transport guarantee that Body is always
// 	// non-nil, even on responses without a body or responses with
// 	// a zero-length body. It is the caller's responsibility to
// 	// close Body. The default HTTP client's Transport does not
// 	// attempt to reuse HTTP/1.0 or HTTP/1.1 TCP connections
// 	// ("keep-alive") unless the Body is read to completion and is
// 	// closed.
// 	//
// 	// The Body is automatically dechunked if the server replied
// 	// with a "chunked" Transfer-Encoding.
// 	Body io.ReadCloser

// // Reader is the interface that wraps the basic Read method.
// //
// // Read reads up to len(p) bytes into p. It returns the number of bytes
// // read (0 <= n <= len(p)) and any error encountered. Even if Read
// // returns n < len(p), it may use all of p as scratch space during the call.
// // If some data is available but not len(p) bytes, Read conventionally
// // returns what is available instead of waiting for more.
// //
// // When Read encounters an error or end-of-file condition after
// // successfully reading n > 0 bytes, it returns the number of
// // bytes read. It may return the (non-nil) error from the same call
// // or return the error (and n == 0) from a subsequent call.
// // An instance of this general case is that a Reader returning
// // a non-zero number of bytes at the end of the input stream may
// // return either err == EOF or err == nil. The next Read should
// // return 0, EOF.
// //
// // Callers should always process the n > 0 bytes returned before
// // considering the error err. Doing so correctly handles I/O errors
// // that happen after reading some bytes and also both of the
// // allowed EOF behaviors.
// //
// // Implementations of Read are discouraged from returning a
// // zero byte count with a nil error, except when len(p) == 0.
// // Callers should treat a return of 0 and nil as indicating that
// // nothing happened; in particular it does not indicate EOF.
// //
// // Implementations must not retain p.
// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }
