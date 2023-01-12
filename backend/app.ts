import { Express, Request, Response } from "express";

import dotenv from "dotenv";

dotenv.config();

const express = require("express");
const connectDB = require("./config/db");

const app: Express = express();
const port = process.env.PORT || 8082;

connectDB();

app.get("/", (req: Request, res: Response) => {
  res.send("Test test");
});

app.listen(port, () => {
  console.log(`⚡️[server]: Server is running at http://localhost:${port}`);
});
