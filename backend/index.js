const express = require("express");
const bodyParser = require("body-parser");
const app = express();
const port = 5000;

app.use(bodyParser.json());
app.use(
  bodyParser.urlencoded({
    extended: true,
  })
);

// Routes
app.get("/", (request, response) => {
  response.json({ message: "Backend API" });
});

app.listen(port, () => {
  console.log(`App running on port ${port}`);
});
