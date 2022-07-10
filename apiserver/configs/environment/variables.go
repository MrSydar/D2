package environment

var Names = struct {
	MongoUri,
	Database,
	CompanyCollection,
	ItemCollection,
	PlaceCollection string
}{
	MongoUri:          "MONGOURI",
	Database:          "DATABASENAME",
	CompanyCollection: "COMPANYCOLLECTION",
	ItemCollection:    "ITEMCOLLECTION",
	PlaceCollection:   "PLACECOLLECTION",
}
