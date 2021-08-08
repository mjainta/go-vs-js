const app = require("../app.js");
const supertest = require("supertest");
const { MongoClient } = require('mongodb')

const url = process.env.DB_URI
const client = new MongoClient(url)
const dbName = process.env.DB_NAME

beforeEach(async () => {
  await client.connect()
  const db = client.db(dbName)
  const collection = db.collection(process.env.DB_COLL)
  await collection.deleteMany({})
  await collection.insertMany([
    {
      "_id": "0",
      "name": "name0",
    },
    {
      "_id": "1",
      "name": "name1",
    },
    {
      "_id": "2",
      "name": "name2",
    },
  ])
});

afterEach(async () => {
  await client.connect()
  const db = client.db(dbName)
  const collection = db.collection(process.env.DB_COLL)
  await collection.deleteMany({})
});

it('Testing to see if Jest works', () => {
  expect(1).toBe(1)
})

test("GET /documents", async () => {
  await supertest(app).get("/documents")
    .expect(200)
    .then((response) => {
      // Check type and length
      expect(Array.isArray(response.body)).toBeTruthy();
      expect(response.body.length).toEqual(3);

      // Check data
      expect(response.body[0]._id).toBe("0");
      expect(response.body[0].name).toBe("name0");
      expect(response.body[1]._id).toBe("1");
      expect(response.body[1].name).toBe("name1");
      expect(response.body[2]._id).toBe("2");
      expect(response.body[2].name).toBe("name2");
    });
});