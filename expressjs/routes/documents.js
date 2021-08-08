
const { MongoClient } = require('mongodb')

const url = process.env.DB_URI
const client = new MongoClient(url)
const dbName = process.env.DB_NAME

var express = require('express');
var router = express.Router();

/* GET document listings. */
router.get('/', async function(req, res, next) {
  await client.connect()
  const db = client.db(dbName)
  const collection = db.collection(process.env.DB_COLL)
  const findResult = await collection.find({}).toArray()
  await client.close()

  res.send(findResult);
});

module.exports = router;
