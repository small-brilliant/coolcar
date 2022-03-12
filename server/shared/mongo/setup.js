db.account.createIndex({
	open_id: 1,
},{
	unique: true,
})

db.trip.createIndex({
    "trip.accountid":1,//这个1表示从小到大的意思
    "trip.status":1,
},{
    unique: true,//以上面的字段不能重复基础创建索引
    partialFilterExpression:{
        "trip.status":1,//trip.status为1的情况下只能有一个
    }
})