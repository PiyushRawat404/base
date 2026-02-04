import {Router} from "express"
import { getUser,newUser,sqlUser } from "../controllers/user"
const router=Router()

router.get("/:name",getUser)
router.post("/register",newUser)
router.get("/sql/result",sqlUser)

export default router

