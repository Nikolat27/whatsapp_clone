import express from "express";
import userRoutes from "./routes/userRoutes";

const app = express();
const port = 8000;

app.listen(port, "api.localhost", () => {
    console.log("Server started successfully!");
});

app.use("/users", userRoutes);
