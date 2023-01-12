const mongoose = require("mongoose");
const config = require("config");
const db = config.get("mongoURI");

const connectDB = async () => {
  try {
    mongoose.set("strictQuery", true);
    await mongoose.connect(db);

    console.log("MongoDB is connected");
  } catch (err) {
    console.log("MongoDB connection has failed \n");
    console.log(`Error: ${err.message}`);
    process.exit(1);
  }
};

module.exports = connectDB;
