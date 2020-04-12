package encode

import (
	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/encoding/consts"
	"github.com/go-spatial/geom/encoding/common/baseencoder"
)

type Encoder struct {
	Be *baseencoder.BaseEncoder
}

func (en *Encoder) Err() error {
        return en.Be.Err()
}

func (en *Encoder) Point(pt [2]float64) {
	en.Be.BOM().Write(consts.Point, pt[0], pt[1])
}
func (en *Encoder) MultiPoint(pts [][2]float64) {
	en.Be.BOM().Write(consts.MultiPoint, uint32(len(pts)))

	for _, p := range pts {
		en.Point(p)
	}
}
func (en *Encoder) LineString(ln [][2]float64) {
	en.Be.BOM().Write(consts.LineString, uint32(len(ln)))
	for _, p := range ln {
		en.Be.Write(p[0], p[1])
	}
}

func (en *Encoder) MultiLineString(lns [][][2]float64) {
	en.Be.BOM().Write(consts.MultiLineString, uint32(len(lns)))
	for _, l := range lns {
		en.LineString(l)
	}
}

func (en *Encoder) Polygon(ply [][][2]float64) {
	en.Be.BOM().Write(consts.Polygon, uint32(len(ply)))
	for _, r := range ply {
		// close definition is:
		// •  Verify that the line segments close (z coordinates at start and endpoints must also be the same) and don't cross.
		// gotten from: http://edndoc.esri.com/arcsde/9.0/general_topics/shape_validation.htm
		var needToClose bool
		length := uint32(len(r))

		if length > 0 && (r[0][0] != r[length-1][0] || r[0][1] != r[length-1][1]) {
			// Let's close the ring.
			length += 1
			needToClose = true
		}
		en.Be.Write(length)
		for _, pt := range r {
			en.Be.Write(pt[0], pt[1])
		}
		if needToClose {
			en.Be.Write(r[0][0], r[0][1])
		}
	}
}

func (en *Encoder) MultiPolygon(mply [][][][2]float64) {
	en.Be.BOM().Write(consts.MultiPolygon, uint32(len(mply)))
	for _, p := range mply {
		en.Polygon(p)
	}
}

func (en *Encoder) Collection(geoms []geom.Geometry) {
	if !en.Be.Conti() {
		return
	}
	en.Be.BOM().Write(consts.Collection, uint32(len(geoms)))
	for _, gg := range geoms {
		en.Geometry(gg)
		if !en.Be.Conti() {
			return
		}
	}
}

func (en *Encoder) Geometry(g geom.Geometry) {
	if !en.Be.Conti() {
		return
	}
	switch geo := g.(type) {
	case geom.Pointer:
		en.Point(geo.XY())
	case geom.MultiPointer:
		en.MultiPoint(geo.Points())
	case geom.LineStringer:
		en.LineString(geo.Vertices())
	case geom.MultiLineStringer:
		en.MultiLineString(geo.LineStrings())
	case geom.Polygoner:
		en.Polygon(geo.LinearRings())
	case geom.MultiPolygoner:
		en.MultiPolygon(geo.Polygons())
	case geom.Collectioner:
		en.Collection(geo.Geometries())
	default:
		en.Be.SetError(geom.ErrUnknownGeometry{Geom: g})
	}
}
