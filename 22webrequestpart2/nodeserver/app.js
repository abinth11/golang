const express = require("express");
const port = 8080;
const app = express();
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.get("/get-message", async (req, res) => [
  res
    .status(200)
    .json({ status: "success", message: "Success message from node server" }),
]);

app.listen(port, (err) => {
  if (err) {
    process.exit(err);
  } else {
    console.log("Server is listening on port ", port);
  }
});
