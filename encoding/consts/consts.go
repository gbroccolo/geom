package consts

//  geometry types
// http://edndoc.esri.com/arcsde/9.1/general_topics/wkb_representation.htm
// https://github.com/postgis/postgis/blob/master/doc/ZMSgeoms.txt
const (
	Point              uint32 = 1
	LineString         uint32 = 2
	Polygon            uint32 = 3
	MultiPoint         uint32 = 4
	MultiLineString    uint32 = 5
	MultiPolygon       uint32 = 6
	Collection         uint32 = 7

	PointZ             uint32 = 2147483649
	LineStringZ        uint32 = 2147483650
	PolygonZ           uint32 = 2147483651
	MultiPointZ        uint32 = 2147483652
	MultiLineStringZ   uint32 = 2147483653
	MultiPolygonZ      uint32 = 2147483654
	CollectionZ        uint32 = 2147483655

	PointM             uint32 = 1073741825
	LineStringM        uint32 = 1073741826
	PolygonM           uint32 = 1073741827
	MultiPointM        uint32 = 1073741828
	MultiLineStringM   uint32 = 1073741829
	MultiPolygonM      uint32 = 1073741830
	CollectionM        uint32 = 1073741831

	PointZM            uint32 = 3221225473
	LineStringZM       uint32 = 3221225474
	PolygonZM          uint32 = 3221225475
	MultiPointZM       uint32 = 3221225476
	MultiLineStringZM  uint32 = 3221225477
	MultiPolygonZM     uint32 = 3221225478
	CollectionZM       uint32 = 3221225479

	PointS             uint32 = 536870913
	LineStringS        uint32 = 536870914
	PolygonS           uint32 = 536870915
	MultiPointS        uint32 = 536870916
	MultiLineStringS   uint32 = 536870917
	MultiPolygonS      uint32 = 536870918
	CollectionS        uint32 = 536870919

	PointZS            uint32 = 2684354561
	LineStringZS       uint32 = 2684354562
	PolygonZS          uint32 = 2684354563
	MultiPointZS       uint32 = 2684354564
	MultiLineStringZS  uint32 = 2684354565
	MultiPolygonZS     uint32 = 2684354566
	CollectionZS       uint32 = 2684354567

	PointMS            uint32 = 1610612737
	LineStringMS       uint32 = 1610612738
	PolygonMS          uint32 = 1610612739
	MultiPointMS       uint32 = 1610612740
	MultiLineStringMS  uint32 = 1610612741
	MultiPolygonMS     uint32 = 1610612742
	CollectionMS       uint32 = 1610612743

	PointZMS           uint32 = 3758096385
	LineStringZMS      uint32 = 3758096386
	PolygonZMS         uint32 = 3758096387
	MultiPointZMS      uint32 = 3758096388
	MultiLineStringZMS uint32 = 3758096389
	MultiPolygonZMS    uint32 = 3758096390
	CollectionZMS      uint32 = 3758096391
)
