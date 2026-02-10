import {Router} from "express"
import { handleFakeaApi } from "../controllers/client.controllers"
const router=Router()

router.get("/todos",handleFakeaApi)

export default router