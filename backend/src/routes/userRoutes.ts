import { App, Request, Response } from "@tinyhttp/app";
import memjs from "memjs";
import bcrypt from "bcrypt";

// const memcached = memjs.Client.create("127.0.0.1:11211");

const router = new App();

async function hashPassword(password: string) {
    return await bcrypt.hash(password, 10);
}

async function verifyPassword(plainPassword: string, hashedPassword: string) {
    return await bcrypt.compare(plainPassword, hashedPassword);
}

router.get("/login", async (req: Request, res: Response) => {
    const password = req.query.password;
    if (typeof password !== "string") return res.status(400);
    let hashedPassword = await hashPassword(password);
    res.send(`Hashed Password: ${hashedPassword}`);

    let verify: boolean = await verifyPassword(password, hashedPassword);
    console.log("Verify: ", verify)
    res.send(verify);
});

export default router;
