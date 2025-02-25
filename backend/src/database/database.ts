import mongoose from "mongoose";

const connectDB = async () => {
    try {
        mongoose.connect("mongodb://127.0.0.1:27017/mydatabase");
    } catch (error) {
        console.error("Error connecting to database: ", error);
        process.exit(1);
    }
};

export default connectDB