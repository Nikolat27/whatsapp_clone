import mongoose, { Schema, Document } from "mongoose";

interface IChat extends Document {
    participants: mongoose.Types.ObjectId;
    messages: mongoose.Types.ObjectId;
    lastMessage: mongoose.Types.ObjectId;
    unreadCounts: Record<string, number>;
    createdAt: Date;
    updatedAt: Date;
}

const ChatSchema = new Schema<IChat>({
    participants: [
        {
            type: mongoose.Schema.Types.ObjectId,
            ref: "User",
            required: true,
        },
    ],
    messages: [
        {
            type: mongoose.Schema.Types.ObjectId,
            ref: "Message",
        },
    ],
    lastMessage: {
        type: mongoose.Schema.Types.ObjectId,
        ref: "Message",
    },
    unreadCounts: {
        type: Map,
        of: Number,
        default: {},
    },
    createdAt: {
        type: Date,
        default: Date.now(),
    },
    updatedAt: {
        type: Date,
        default: Date.now(),
    },
});

const Chat = mongoose.model<IChat>("Chat", ChatSchema);
export default Chat;
