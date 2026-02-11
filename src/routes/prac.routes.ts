import {Router} from "express"
import { getUser,handlegetuser,handlePostUser,newUser,sqlUser } from "../controllers/prac"
const router=Router()

router.get("/name",getUser)
router.post("/register",newUser)
router.get("/sql/result",sqlUser)
router.get("/users",handlegetuser)
router.post("/doctor",handlePostUser)

export default router

