package errinvalidtype

import (
        "fmt"
)

type ErrInvalidType struct {
        // In which collection type was the invalide type found.
        Primary string
        Type    uint32
}

func (e ErrInvalidType) Error() string {
        return fmt.Sprintf("decode: invalid type for %v", e.Primary)
}
