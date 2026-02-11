import { Request,Response,NextFunction } from "express"

type asyncFn=(
    req:Request,
    res:Response,
    next:NextFunction
)=>Promise<unknown>



export const asyncHandler=(fn:asyncFn)=>{
    return async(req:Request,res:Response,next:NextFunction)=>{
        try{
            const result=await fn(req,res,next);
            return result
        }catch(error){
            // throw new Error("internal server error")
        return res.json("error" +error)
        }
    }
}