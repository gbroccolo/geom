package baseencoder

import (
	"encoding/binary"
        "errors"
        "io"
)

// base struct to be embedded
type BaseEncoder struct {
        // W is the writer to which the binary data will be written to.
        W io.Writer
        // ByteOrder is the Byte Order Marker, it defaults to binary.LittleEndian
        ByteOrder binary.ByteOrder
        err       error
}

var EncoderIsNilErr = errors.New("Encoder can not be nil")

func (en *BaseEncoder) Conti() bool { return !(en == nil || en.err != nil) }

func (en *BaseEncoder) Write(data ...interface{}) *BaseEncoder {
        if !en.Conti() {
                return en
        }
        if en.ByteOrder == nil {
                en.ByteOrder = binary.LittleEndian
        }
        for i := range data {
                en.err = binary.Write(en.W, en.ByteOrder, data[i])
                if en.err != nil {
                        return en
                }
        }
        return en
}

func (en *BaseEncoder) Err() error {
        if en == nil {
                return EncoderIsNilErr
        }
        return en.err
}

func (en *BaseEncoder) BOM() *BaseEncoder {
        if !en.Conti() {
                return en
        }
        if en.ByteOrder != nil && en.ByteOrder == binary.BigEndian {
                en.Write(byte(0))
                return en
        }
        en.Write(byte(1))
        return en
}

func (en *BaseEncoder) SetError(err error) {
        en.err = err
}
