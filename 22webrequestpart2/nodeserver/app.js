const express = require("express");
const fs = require("fs");
const path = require("path");

const port = 8080;
const app = express();

app.use(express.json());
app.use(express.urlencoded({ extended: false }));

app.get("/get-message", (req, res) => [
  res
    .status(200)
    .json({ status: "success", message: "Success message from node server" }),
]);

app.post("/save-user", (req, res) => {
  try {
    const userData = req.body;

    const dataFilePath = path.join(__dirname, "data.json");

    fs.readFile(dataFilePath, "utf8", (err, data) => {
      if (err && err.code !== "ENOENT") {
        return res.status(500).json({ error: "Failed to read data file" });
      }

      let jsonData = [];

      if (data) {
        try {
          jsonData = JSON.parse(data);
        } catch (parseError) {
          return res.status(500).json({ error: "Failed to parse data file" });
        }
      }

      jsonData.push(userData);

      fs.writeFile(
        dataFilePath,
        JSON.stringify(jsonData, null, 2),
        (writeError) => {
          if (writeError) {
            return res.status(500).json({ error: "Failed to save user data" });
          }

          res.status(200).json({ message: "User data saved successfully" });
        }
      );
    });
  } catch (error) {
    console.log(error);
    return res.status(500).json({ error: "Failed to save new user" });
  }
});

app.post("/handle-post-form", (req, res) => {
  console.log(req.body);
  res.status(200).json("received the data from client");
});

app.listen(port, (err) => {
  if (err) {
    process.exit(err);
  } else {
    console.log("Server is listening on port ", port);
  }
});
