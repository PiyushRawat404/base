import {Router} from "express"
import { handleFakeaApi } from "../controllers/client.routes"
const router=Router()

router.get("/todos",handleFakeaApi)

export default router