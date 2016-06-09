package main

import (
	"bufio"
	"io"
)

func write() {
	_, err = fd.Write(p0[a:b])
	if err != nil {
		return err
	}
	_, err = fd.Write(p1[c:d])
	if err != nil {
		return err
	}
	_, err = fd.Write(p2[e:f])
	if err != nil {
		return err
	}
	// and so on
}

type errWriter struct {
	w   io.Writer
	err error
}

func (ew *errWriter) write(buf []byte) {
	if ew.err != nil {
		return
	}
	_, ew.err = ew.w.Write(buf)
}

func write2() {
	ew := &errWriter{w: fd}
	ew.write(p0[a:b])
	ew.write(p1[c:d])
	ew.write(p2[e:f])
	if ew.err != nil {
		return ew.err
	}
}

func scan() {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		token := scanner.Text()
		// process token
	}
	if err := scanner.Err(); err != nil {
		// process the error
	}
}
