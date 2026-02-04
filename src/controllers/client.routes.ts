import {Request ,Response} from "express"
import {connectRedis as cR} from "../db/redis"
import axios from "axios"


const handleFakeaApi =async(req:Request,res:Response)=>{

  try{
    const cacheValue =await cR.get("todos")
  if(cacheValue) return res.json(JSON.parse(cacheValue))
  const {data}=await axios.get("https://jsonplaceholder.typicode.com/todos")
  await cR.set("todos",JSON.stringify(data))
  await cR.expire("todos",60)
  return res.json(data);
}catch (error) {
    console.error(error)
    return res.status(500).json({ message: "Something went wrong" })
  }
}

export {handleFakeaApi}