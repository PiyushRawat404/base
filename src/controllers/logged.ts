import {Request ,Response} from "express"
import {connectRedis as cR} from "../db/redis"
import axios from "axios"
import { asyncHandler } from "../services/asyncHandler"
import { prisma } from "../lib/prisma"


const handleFakeaApi =asyncHandler(async(req:Request,res:Response)=>{
  const cacheValue =await cR.get("todos")
  if(cacheValue) return res.json(JSON.parse(cacheValue))
  const {data}=await axios.get("https://jsonplaceholder.typicode.com/todos")
  await cR.set("todos",JSON.stringify(data))
  await cR.expire("todos",60)
  return res.json(data);
})

const handleSignupUser=asyncHandler(async(req:Request,res:Response)=>{
  const {fullname,email,password}=req.body

  const lD=await prisma.patients.create({
    data:{
        fullname,
        email,
        password
    }
  })
  return res.json(lD)
})

const handleSigninUser=asyncHandler(async(req:Request,res:Response)=>{
  const{email,password}=req.body
  
  const getUser=await prisma.patients.findMany({
    where:{
      email:email,
      password:password 
    }
  })

  res.send(`logged in as ${getUser}`)
})

export {handleFakeaApi,handleSigninUser,handleSignupUser}