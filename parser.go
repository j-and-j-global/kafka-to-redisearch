package main

// Parser is a library, type, or whatever which may be used to
// transform input code from a CMS, which can be matched via the
// provenance field of an incoming message, into something which
// our presentation layers can understand
//
// We generate html with some pretty open tags- it may be that in the
// future this is less appropriate; we may need a more robust/ descriptive
// payload to return on, like some descriptive json. Because this interface
// returns a byte slice, which redis writes pretty much verbatim, we can
// switch to such a standard relatively easily
type Parser interface {
	// Parse takes whatever representation the CMS writes, and transforms
	// it into something our presentation/ front end can handle
	Parse([]byte) ([]byte, error)
}
