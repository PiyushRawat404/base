import express, { Request, Response } from "express";
import mongoose from 'mongoose';
import userRoutes from "./routes/user.routes"
import axios from "axios"
import redis from "redis"

// const redisClient =redis.createClient();
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

app.get("/api/todos",async(req:Request,res:Response)=>{
  const {data}=await axios.get("https://jsonplaceholder.typicode.com/todos")
  return res.json(data);
})


app.get("/api/v1", (req: Request, res: Response) => {
  res.send("Hello from TypeScript");
});

app.listen(`${Port}`,()=>console.log(`Server running on http://localhost:${Port}`))