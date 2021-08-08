const { MongoClient } = require('mongodb')

const url = 'mongodb://admin:foobar@localhost'
const client = new MongoClient(url)
const dbName = 'mydatabase'

var express = require('express');
var router = express.Router();

/* GET document listings. */
router.get('/', async function(req, res, next) {
  await client.connect()
  console.log('Connected successfully to server')
  const db = client.db(dbName)
  const collection = db.collection('mycoll')
  const findResult = await collection.find({}).toArray()

  res.send(findResult);
});

module.exports = router;
