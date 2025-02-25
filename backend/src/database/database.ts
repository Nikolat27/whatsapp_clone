import mongoose from "mongoose";



const connectDB = async () => {
    
        mongoose.connect("mongodb://127.0.0.1:27017/mydatabase");
   
};

function a() {


}

export default connectDB