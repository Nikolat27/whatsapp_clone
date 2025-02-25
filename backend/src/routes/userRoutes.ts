import { Router, Request, Response } from "express";
import connectDB from "../database/database";
import User from "../database/models/User";

const router = Router();

router.get("/", (req: Request, res: Response) => {
    res.send("hi")
})

export default router