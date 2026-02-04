import {Request ,Response} from "express"
import { User } from "../model/user"
import client from "../db/postgresql"


const getUser=async(req:Request,res:Response)=>{
    const name=req.params.name;
    await User.find({name})
    
    res.send(name)
}

const newUser=async(req:Request,res:Response)=>{
    try {
    const { name, num, gender, problem } = req.body;

    if (!name || !num || !gender || !problem) {
      return res.status(400).json({ msg: "All fields are required" });
    }
    
   const user=await User.create({
        clientNo:Date.now(),
        fullName:name,
        phoneNo:num,
        gender:gender,
        problem:problem
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
  const result =await client.query("Select * from public.patient")
  // res.send(result.rows)
  // res.json("hello")
  console.log("hello")
}

export {getUser,newUser,sqlUser}