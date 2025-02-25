import mongoose, { Schema, Document } from "mongoose";

interface IUser extends Document {
    email: string;
    name: string;
    password: string;
    aboutMe?: string; // Optional
    createdAt: Date;
}

const UserSchema = new Schema<IUser>({
    email: { type: String, required: true, unique: true },
    name: { type: String, required: true },
    password: { type: String, required: true },
    aboutMe: { type: String, default: '' },
    createdAt: { type: Date, default: Date.now },
})


const User = mongoose.model<IUser>("User", UserSchema)
export default User