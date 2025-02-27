import { App, Request, Response } from "@tinyhttp/app";
import bcrypt from "bcrypt";
import User from "../database/models/User";
import memjs from "memjs";
import fs from "fs";
import path from "path";

const memcached = memjs.Client.create("127.0.0.1:11211");
const router = new App();

function generateSessionIdCode(length: number): string {
    const characters =
        "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    let code = "";

    for (let i = 0; i <= length; i++) {
        const randomIndex = Math.floor(Math.random() * characters.length);
        code += characters[randomIndex];
    }
    return code;
}

function hashPassword(password: string): Promise<string> {
    return bcrypt.hash(password, 10);
}

function verifyPassword(
    plainPassword: string,
    hashedPassword: string
): Promise<boolean> {
    return bcrypt.compare(plainPassword, hashedPassword);
}

async function getHashedPassword(username: string) {
    const user = await User.findOne({
        username: username,
    }).select("password");
    return user?.password;
}

async function checkFieldExistence(field: string, value: string) {
    const user = await User.findOne({ [field]: value });
    return user ? true : false;
}

async function createUser(
    email: string,
    username: string,
    hashedPassword: string
) {
    await User.create({
        email,
        username,
        password: hashedPassword,
    });
}

async function authenticateUser(username: string, plainPassword: string) {
    const doesUsernameExist = await checkFieldExistence("username", username);
    if (!doesUsernameExist) return false;

    // Verify password
    const hashedPassword = await getHashedPassword(username);
    if (!hashedPassword) return false;

    const isPasswordCorrect = await verifyPassword(
        plainPassword,
        hashedPassword
    );

    return isPasswordCorrect;
}

async function getUserId(username: string): Promise<number> {
    const user = await User.findOne({
        username,
    }).select("id");
    return user?.id;
}

router.post("/register", async (req: Request, res: Response) => {
    const { email, username, password } = req.body;
    const hashedPassword = await hashPassword(password);

    // Validating
    const doesEmailExist = await checkFieldExistence("email", email);
    const doesUsernameExist = await checkFieldExistence("username", username);

    if (doesEmailExist || doesUsernameExist) {
        res.status(400).send("Either username or email is used!");
        return;
    }
    await createUser(email, username, hashedPassword);

    res.send("User added successfully!");
});

router.post("/login", async (req: Request, res: Response) => {
    const { username, password, userSessionId } = req.body;

    // Check session id exists
    if (userSessionId) {
        const doesSessionIdExist = await memcached.get(userSessionId);
        if (doesSessionIdExist.value) {
            return res.status(400).send("User is already loged in!");
        }
    }

    const isUserAuthenticated = await authenticateUser(username, password);
    if (!isUserAuthenticated) {
        return res
            .status(401)
            .send("Either Password or Username is incorrect!");
    }

    const userId: number = await getUserId(username);
    const sessionId: string = generateSessionIdCode(12);

    await memcached.set(sessionId, userId.toString(), { expires: 3600 * 24 });
    return res.status(200).send({
        user_session_id: sessionId,
    });
});

router.delete("/logout", async (req: Request, res: Response) => {
    const { userSessionId } = req.body;

    if (!userSessionId) {
        return res.status(400).send("User is already logged out");
    }

    await memcached.delete(userSessionId);
    return res.status(200).send("User loged out successfully!");
});

router.put("/update-profile", async (req: Request, res: Response) => {
    const { profilePicture } = req.body;
    console.log(profilePicture);
    return;
});

export default router;
