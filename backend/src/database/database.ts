import mongoose from "mongoose";

const connectDB = async () => {
    mongoose.connect("mongodb://localhost:27017/whatsapp_db");
};
export default connectDB;
