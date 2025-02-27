import { App } from "@tinyhttp/app";
import { cors } from "@tinyhttp/cors";
import userRoutes from "./routes/userRoutes";
import { json } from "milliparsec";
import connectDB from "./database/database";

const app = new App();
const port = 8000;

app.use(cors({ origin: "http://localhost:5000" }));
app.use(json());
app.options("*", cors());
app.use("/users", userRoutes);

connectDB().then(() => app.listen(port));
