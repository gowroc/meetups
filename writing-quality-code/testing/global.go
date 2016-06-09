const port = 1000 // Bad.
var port = 1000 // Better.
const defaultPort = 1000 // Best.
type ServerOpts struct {
    Port int // Make it default somewhere.
}