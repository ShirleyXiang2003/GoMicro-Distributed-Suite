const express = require("express");
const app = express();

// Middleware to parse JSON request bodies
app.use(express.json());

function readJSON(req, res, callback) {
  try {
    // req.body already contains the parsed JSON if you use express.json() middleware
    callback(null, req.body);
  } catch (error) {
    callback(error, null);
  }
}

app.post("/some-endpoint", (req, res) => {
  readJSON(req, res, (err, data) => {
    if (err) {
      return res.status(400).send("Invalid JSON");
    }
    // Do something with the data
    res.send("JSON received and parsed successfully");
  });
});

app.listen(8080, () => {
  console.log("Server is running on port 8080");
});
