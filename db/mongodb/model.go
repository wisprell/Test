package mongodb

type MongoCount struct {
	Type  string `bson:"type"`  //开始时间
	Count int    `bson:"count"` //结束时间
}
