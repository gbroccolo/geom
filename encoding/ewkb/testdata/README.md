# Quick Description of the testdata file.

## Desc
Desc is a quick description of the test case.
This Also, starts the test block.

## Expected

The expected "extended" Geometry Collection:

Geometries are described as follows:
`xxx,yyy` for a point.
`xxx,yyy,zzz` for a 3DZ point.
`xxx,yyy,zzz,mmm` for a 4D point.
`xxx,yyy,mmm` for a 3DM point.

Information about the SRID can be specified:
`iiii xxx,yyy` for a point defined in SRID=iiii.
`iiii xxx,yyy,zzz` for a 3DZ point defined in SRID=iiii.
etc.

There are then the other extended geometries, i.e.:
`( xxx,yyy,zzz xxx,yyy,zzz )` for MultiPoint.
`(( ))` for a collection.
`{{ }}` for a MultiPolygon
`{ }` for a Polygon
`[[ ]]` for a MutliLineString
`[ ]` for a LineString

Also in this case, the information about the SRID can be added:
`iiii ( xxx,yyy,zzz xxx,yyy,zzz )` for MultiPoint.
etc.

### 3DZ Point
```
expected:
34,45,56
```

### 3DZ MultiPoint
```
expected:
(1,2,3 34,45,56)
```

### 3DM LineString
```
expected:
[ 1,2,1234  34,45,2345 ]
```

### MultiLineString in WGS84 projection
```
4326 [[ [1,2 34,45] [12,23 1,4] ]]
```

### Collection in WGS84 projection
```
4326 (( 12,23 (1,2 3,4) [1,2 3,4] [[ [1,2 3,4] [1,4 5,6] ]] ))
```

## Bytes 

Are described using hex numbers. Each hex number is expected to be two digits: from 00-FF; or from 00-ff.

Extended WKB protocol used by PostGIS is documented here: https://github.com/postgis/postgis/blob/master/doc/ZMSgeoms.txt

## BOM

This is the Endian value, it can be eighter “little” or “big”.
