//Package ewkb is for decoding PostGIS Extended Well Known Binary (EWKB) format for OGC geometry (EWKBGeometry)
// sepcification at https://github.com/postgis/postgis/blob/master/doc/ZMSgeoms.txt
package ewkb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/encoding/consts"
	"github.com/go-spatial/geom/encoding/common/byteordertype"
	"github.com/go-spatial/geom/encoding/ewkb/internal/decode"
	"github.com/go-spatial/geom/encoding/common/baseencoder"
	"github.com/go-spatial/geom/encoding/ewkb/internal/encode"
)

type ErrUnknownGeometryType struct {
	Typ uint32
}

func (e ErrUnknownGeometryType) Error() string {
	return fmt.Sprintf("Unknown Geometry Type %v", e.Typ)
}

//  geometry types
const (
	PointZ             = consts.PointZ
	LineStringZ        = consts.LineStringZ
	PolygonZ           = consts.PolygonZ
	MultiPointZ        = consts.MultiPointZ
	MultiLineStringZ   = consts.MultiLineStringZ
	MultiPolygonZ      = consts.MultiPolygonZ
	CollectionZ        = consts.CollectionZ

	PointM             = consts.PointM
	LineStringM        = consts.LineStringM
	PolygonM           = consts.PolygonM
	MultiPointM        = consts.MultiPointM
	MultiLineStringM   = consts.MultiLineStringM
	MultiPolygonM      = consts.MultiPolygonM
	CollectionM        = consts.CollectionM

	PointZM            = consts.PointZM
	LineStringZM       = consts.LineStringZM
	PolygonZM          = consts.PolygonZM
	MultiPointZM       = consts.MultiPointZM
	MultiLineStringZM  = consts.MultiLineStringZM
	MultiPolygonZM     = consts.MultiPolygonZM
	CollectionZM       = consts.CollectionZM

	PointS             = consts.PointS
	LineStringS        = consts.LineStringS
	PolygonS           = consts.PolygonS
	MultiPointS        = consts.MultiPointS
	MultiLineStringS   = consts.MultiLineStringS
	MultiPolygonS      = consts.MultiPolygonS
	CollectionS        = consts.CollectionS

	PointZS            = consts.PointZS
        LineStringZS       = consts.LineStringZS
        PolygonZS          = consts.PolygonZS
        MultiPointZS       = consts.MultiPointZS
        MultiLineStringZS  = consts.MultiLineStringZS
        MultiPolygonZS     = consts.MultiPolygonZS
        CollectionZS       = consts.CollectionZS

	PointMS            = consts.PointMS
        LineStringMS       = consts.LineStringMS
        PolygonMS          = consts.PolygonMS
        MultiPointMS       = consts.MultiPointMS
        MultiLineStringMS  = consts.MultiLineStringMS
        MultiPolygonMS     = consts.MultiPolygonMS
        CollectionMS       = consts.CollectionMS

	PointZMS           = consts.PointZMS
        LineStringZMS      = consts.LineStringZMS
        PolygonZMS         = consts.PolygonZMS
        MultiPointZMS      = consts.MultiPointZMS
        MultiLineStringZMS = consts.MultiLineStringZMS
        MultiPolygonZMS    = consts.MultiPolygonZMS
        CollectionZMS      = consts.CollectionZMS
)

// DecodeBytes will attempt to decode a geometry encoded as WKB into a geom.Geometry.
func DecodeBytes(b []byte) (geo geom.Geometry, err error) {
	buff := bytes.NewReader(b)
	return Decode(buff)
}

// Decode will attempt to decode a geometry encoded as WKB into a geom.Geometry.
func Decode(r io.Reader) (geo geom.Geometry, err error) {

	bom, typ, err := byteordertype.ByteOrderType(r)
	if err != nil {
		return nil, err
	}
	switch typ {
	case Point:
		pt, err := decode.Point(r, bom)
		return geom.Point(pt), err
	case MultiPoint:
		mpt, err := decode.MultiPoint(r, bom)
		return geom.MultiPoint(mpt), err
	case LineString:
		ln, err := decode.LineString(r, bom)
		return geom.LineString(ln), err
	case MultiLineString:
		mln, err := decode.MultiLineString(r, bom)
		return geom.MultiLineString(mln), err
	case Polygon:
		pl, err := decode.Polygon(r, bom)
		return geom.Polygon(pl), err
	case MultiPolygon:
		mpl, err := decode.MultiPolygon(r, bom)
		return geom.MultiPolygon(mpl), err
	case Collection:
		col, err := decode.Collection(r, bom)
		return col, err
	default:
		return nil, ErrUnknownGeometryType{typ}
	}
}

func EncodeBytes(g geom.Geometry) (bs []byte, err error) {
	buff := new(bytes.Buffer)
	if err = Encode(buff, g); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func Encode(w io.Writer, g geom.Geometry) error {
	return EncodeWithByteOrder(binary.LittleEndian, w, g)
}

func EncodeWithByteOrder(byteOrder binary.ByteOrder, w io.Writer, g geom.Geometry) error {
	en := encode.Encoder{Be: &baseencoder.BaseEncoder{W: w, ByteOrder: byteOrder}}
	en.Geometry(g)
	return en.Err()
}
