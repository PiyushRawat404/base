import express, { Request, Response } from "express";
import mongoose from 'mongoose';
import userRoutes from "./routes/user.routes"
import cientRoute from "./routes/client.routes"
import axios from "axios"
import { ClientRequest } from "node:http";
// import client from "./client.js"

const app=express()
app.use(express.json())
app.use(express.urlencoded({ extended: true }));
const Port=8000

mongoose.connect("mongodb://localhost:27017/healthcheckup")
.then(() => console.log('Connected to MongoDB'))
.catch((error) => {
    console.error('Error connecting to MongoDB:', error.message);
});

app.use("/",userRoutes)
app.use("/api",cientRoute)

// app.get("/api/todos",async(req:Request,res:Response)=>{

//   const cacheValue =await client.get("todos")
//   if(cacheValue) return res.json(cacheValue)
//   const {data}=await axios.get("https://jsonplaceholder.typicode.com/todos")
//   await client.set("todos",data)
//   await client.expire("todos",60)
//   return res.json(data);
// })


app.get("/api/v1", (req: Request, res: Response) => {
  res.send("Hello from TypeScript");
});

app.listen(`${Port}`,()=>console.log(`Server running on http://localhost:${Port}`))