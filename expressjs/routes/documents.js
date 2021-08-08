var express = require('express');
var router = express.Router();

/* GET document listings. */
router.get('/', function(req, res, next) {
  res.send('sdfsdf respond with a resource');
});

module.exports = router;
