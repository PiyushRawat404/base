import {Router} from "express"
import { handleFakeaApi, handleSigninUser, handleSignupUser } from "../controllers/logged"
const router=Router()

router.get("/redis",handleFakeaApi)
router.post("/signup",handleSignupUser)
router.get("/signin",handleSigninUser)


export default router