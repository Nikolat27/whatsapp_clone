import { App } from "@tinyhttp/app";
import userRoutes from "./routes/userRoutes";

const app = new App();
const port = 8000;

app.use("/users", userRoutes);

app.listen(port, () => {
    console.log("Server is started successfully!");
});
