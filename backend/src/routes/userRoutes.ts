import { App, Request, Response } from "@tinyhttp/app";
import memjs from "memjs";

const memcached = memjs.Client.create("127.0.0.1:11211");

const router = new App();

router.get("/login", async (req: Request, res: Response) => {
    await memcached.set("username", "password");

        

});

export default router;
