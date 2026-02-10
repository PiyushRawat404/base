import { Request,Response } from "express"

export const asyncHandler=(fn:Function)=>{
    return async(req:Request,res:Response,next:any)=>{
        try{
            const result=await fn(req,res,next);
            return result
        }catch(error){
            // throw new Error("internal server error")
        return res.json("error" +error)
        }
    }
}