import {Request ,Response} from "express"
import { User } from "../model/user"
import client from "../db/postgresql"
import { prisma } from '../lib/prisma'
import { error } from "node:console"
import { asyncHandler } from "../utility/asyncHandler"
const getUser=async(req:Request,res:Response)=>{
    const name=req.params.name;
    await User.find({name})
    
    res.send(name)
}

const newUser=async(req:Request,res:Response)=>{
    try {
    const { fullname, phoneno, gender,age, problem } = req.body;

    if (!fullname || !phoneno || !gender || !problem || !age) {
      return res.status(400).json({ msg: "All fields are required" });
    }
    
   const user=await prisma.patients.create({
    data:{
      fullname,
      phoneno,
      gender,
      age,
      problem,     
    } 
    })
     return res.status(201).json({
      msg: "User created successfully",
      user,
    });
  } catch (error) {
    console.error(error);
    return res.status(500).json({ msg: "Server error" });
  }
};

const sqlUser=async(req:Request,res:Response)=>{
  // const result =await prisma.patients.findFirst({fullname:"pkfi"})
  // res.send(result.rows)
  // res.json("hello")
  console.log("hello")
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