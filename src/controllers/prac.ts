import {Request ,Response} from "express"
import { User } from "../model/user"
// import client from "../db/postgresql"
import { prisma } from '../lib/prisma'
import { asyncHandler } from "../services/asyncHandler"



const getUser=async(req:Request,res:Response)=>{
    const name=req.params.name;
    await User.find({name})
    
    res.send(name)
}

const newUser=async(req:Request,res:Response)=>{
    const { fullname, phoneno, gender,age, problem,email,password } = req.body;
    
   const user=await prisma.patients.create({
    data:{
      fullname,
      phoneno,
      gender,
      age,
      problem,   
      email,
      password,

    } 
    })
     return res.status(201).json({
      msg: "User created successfully",
      user,
    });
  }
const sqlUser=async(req:Request,res:Response)=>{
  const result:unknown =await prisma.patients.findFirst({
    where:{
      fullname:"dl"
    }
  })
  res.json("hello"+result)
}


const handlegetuser=asyncHandler(async(req:Request,res:Response)=>{
  const allUsers= await prisma.patients.findMany({
    where:{
      age: 10
    }
  })
  res.json(allUsers)
})

const handlePostUser=async(req:Request,res:Response)=>{
  try{

    const team = req.body
    const known=await prisma.doctor.create({
     data:{
      name :team.name,
      specialist :team.specialist,
      shift :team.shift,
      email :team.email,
      password :team.password
     } 
    })
    
    return res.json({
      msg:"created",
      known
    })
  }
  catch(error){
    console.log("error :"+ error)
      res.status(501).json({msg:"data error"})
  }finally{
    console.log("rows created")
  }
}

export {getUser,newUser,sqlUser,handlegetuser,handlePostUser}