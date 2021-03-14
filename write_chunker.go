package rasterm

import "io"

type XfrmFunc func([]byte) ([]byte, error)

type WriteChunker struct {
	chunk  []byte
	writer io.Writer
	ix     int
	Xfrm   XfrmFunc
}

/*
	Writer interface which buffers/flushes in chunks that are
	`chunkSize` bytes long.  Optional `Xfrm` function in struct
	allows for additional []byte processing befor sending to the
	underlying io.Writer.
*/
func NewWriteChunker(iWri io.Writer, chunkSize int) WriteChunker {

	if chunkSize < 1 {
		panic("invalid chunk size")
	}

	return WriteChunker{
		chunk:  make([]byte, chunkSize),
		writer: iWri,
	}
}

func (pC *WriteChunker) Flush() (E error) {

	tmp := pC.chunk[:pC.ix]
	if pC.Xfrm != nil {
		if tmp, E = pC.Xfrm(tmp); E != nil {
			return
		}
	}

	_, E = pC.writer.Write(tmp)
	pC.ix = 0
	return
}

func (pC *WriteChunker) Write(src []byte) (int, error) {

	chunkSize := len(pC.chunk)

	for _, bt := range src {

		pC.chunk[pC.ix] = bt
		pC.ix++

		if pC.ix >= chunkSize {
			if e := pC.Flush(); e != nil {
				return 0, e
			}
		}
	}

	return len(src), nil
}