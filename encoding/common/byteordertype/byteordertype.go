package byteordertype

import (
        "encoding/binary"
        "io"
)

type ErrBadBOM byte

func (e ErrBadBOM) Error() string { return "decode: bad byte order marker" }

func ByteOrderType(r io.Reader) (byteOrder binary.ByteOrder, typ uint32, err error) {
        var bom = make([]byte, 1, 1)
        // the bom is the first byte
        if _, err = r.Read(bom); err != nil {
                return byteOrder, typ, err
        }

        // the bom should be either 0 or 1
        switch bom[0] {
        case 0:
                byteOrder = binary.BigEndian
        case 1:
                byteOrder = binary.LittleEndian
        default:
                return byteOrder, typ, ErrBadBOM(bom[0])
        }

        // Reading the type which is 4 bytes
        err = binary.Read(r, byteOrder, &typ)
        return byteOrder, typ, err
}
