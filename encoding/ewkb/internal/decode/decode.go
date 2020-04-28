package decode

import (
	"encoding/binary"
	"io"

	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/encoding/common/byteordertype"
	"github.com/go-spatial/geom/encoding/common/errinvalidtype"
	"github.com/go-spatial/geom/encoding/consts"
)

func PointZ(r io.Reader, bom binary.ByteOrder) (pt geom.Point, err error) {
	type g [3]float64
	pt := g{}
	err = binary.Read(r, bom, &pt)
	return pt, err
}

func PointM(r io.Reader, bom binary.ByteOrder) (pt geom.Point, err error) {
	type g [3]float64
	pt := g{}
        err = binary.Read(r, bom, &pt)
        return pt, err
}

func PointZM(r io.Reader, bom binary.ByteOrder) (pt geom.Point, err error) {
	type g [4]float64
        pt := g{}
        err = binary.Read(r, bom, &pt)
        return pt, err
}

func PointS(r io.Reader, bom binary.ByteOrder) (pt geom.Point, err error) {
        var data struct {
		srid uint32
		pt   [2]float64
	}
        err = binary.Read(r, bom, &data)
        return data.pt, err
}

func PointZS(r io.Reader, bom binary.ByteOrder) (pt geom.Point, err error) {
	var data struct {
                srid uint32
                pt   [3]float64
        }
        err = binary.Read(r, bom, &data)
        return data.pt, err
}

func PointMS(r io.Reader, bom binary.ByteOrder) (pt geom.Point, err error) {
	var data struct {
                srid uint32
                pt   [3]float64
        }
        err = binary.Read(r, bom, &data)
        return pt, err
}

func PointZMS(r io.Reader, bom binary.ByteOrder) (pt geom.Point, err error) {
	var data struct {
                srid uint32
                pt   [4]float64
        }
        err = binary.Read(r, bom, &data)
        return data.pt, err
}







func MultiPoint(r io.Reader, bom binary.ByteOrder) (pts geom.MultiPoint, err error) {
	var num, typ uint32 // Number of points
	err = binary.Read(r, bom, &num)
	if err != nil {
		return pts, err
	}

	pts = make([][2]float64, num)
	for i := range pts {

		bom, typ, err = byteordertype.ByteOrderType(r)
		if err != nil {
			return pts, err
		}
		if typ != consts.Point {
			return pts, errinvalidtype.ErrInvalidType{"multipoint", typ}
		}
		err = binary.Read(r, bom, &pts[i])
		if err != nil {
			return pts, err
		}
	}
	return pts, err
}

func LineString(r io.Reader, bom binary.ByteOrder) (ln geom.LineString, err error) {
	var num uint32 // Number of points
	if err = binary.Read(r, bom, &num); err != nil {
		return ln, err
	}
	ln = make([][2]float64, num)
	for i := range ln {
		if err = binary.Read(r, bom, &ln[i]); err != nil {
			return ln, err
		}
	}
	return ln, err
}

func MultiLineString(r io.Reader, bom binary.ByteOrder) (lns geom.MultiLineString, err error) {
	var num uint32
	if err = binary.Read(r, bom, &num); err != nil {
		return lns, err
	}
	lns = make([][][2]float64, num)
	for i := range lns {
		bom, typ, err := byteordertype.ByteOrderType(r)
		if err != nil {
			return lns, err
		}
		if typ != consts.LineString {
			return lns, errinvalidtype.ErrInvalidType{"multilinestring", typ}
		}
		if lns[i], err = LineString(r, bom); err != nil {
			return lns, err
		}
	}
	return lns, err
}

func LinerRing(r io.Reader, bom binary.ByteOrder) (rn [][2]float64, err error) {
	var num uint32 // Number of points
	if err = binary.Read(r, bom, &num); err != nil {
		return rn, err
	}
	rn = make([][2]float64, num)
	for i := range rn {
		if err = binary.Read(r, bom, &rn[i]); err != nil {
			return rn, err
		}
	}
	if num > 1 {
		// Remove the last point if it is the same.
		if rn[0][0] == rn[num-1][0] && rn[0][1] == rn[num-1][1] {
			rn = rn[:num-1]
		}
	}

	return rn, err
}

func Polygon(r io.Reader, bom binary.ByteOrder) (ply geom.Polygon, err error) {
	var num uint32
	if err = binary.Read(r, bom, &num); err != nil {
		return ply, err
	}
	ply = make([][][2]float64, num)
	for i := range ply {
		if ply[i], err = LinerRing(r, bom); err != nil {
			return ply, err
		}
	}
	return ply, err
}

func MultiPolygon(r io.Reader, bom binary.ByteOrder) (plys geom.MultiPolygon, err error) {
	var num uint32
	if err = binary.Read(r, bom, &num); err != nil {
		return plys, err
	}
	plys = make([][][][2]float64, num)
	for i := range plys {
		bom, typ, err := byteordertype.ByteOrderType(r)
		if err != nil {
			return plys, err
		}
		if typ != consts.Polygon {
			return plys, errinvalidtype.ErrInvalidType{"multipolygon", typ}
		}
		if plys[i], err = Polygon(r, bom); err != nil {
			return plys, err
		}
	}
	return plys, err
}

func Collection(r io.Reader, bom binary.ByteOrder) (col geom.Collection, err error) {
	var num uint32
	if err = binary.Read(r, bom, &num); err != nil {
		return col, err
	}
	col = make(geom.Collection, num)
	for i := range col {
		bom, typ, err := byteordertype.ByteOrderType(r)
		if err != nil {
			return col, err
		}
		switch typ {
		case consts.Point:
			col[i], err = Point(r, bom)
		case consts.LineString:
			col[i], err = LineString(r, bom)
		case consts.Polygon:
			col[i], err = Polygon(r, bom)
		case consts.MultiPoint:
			col[i], err = MultiPoint(r, bom)
		case consts.MultiLineString:
			col[i], err = MultiLineString(r, bom)
		case consts.MultiPolygon:
			col[i], err = MultiPolygon(r, bom)
		case consts.Collection:
			col[i], err = Collection(r, bom)
		default:
			err = errinvalidtype.ErrInvalidType{"collection", typ}
		}
		if err != nil {
			return col, err
		}
	}
	return col, err
}
