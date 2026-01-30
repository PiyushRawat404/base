import {Router} from "express"
import { getUser,newUser } from "../controllers/user"
const router=Router()

router.get("/:name",getUser)
router.post("/register",newUser)

export default router

