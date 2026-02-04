import express from "express";
import { connectMongoDB } from "./db/mongoose";
import userRoutes from "./routes/user.routes"
import cientRoute from "./routes/client.routes"
import { connectDB } from "./db/postgresql";

const app=express()
app.use(express.json())
app.use(express.urlencoded({ extended: true }));
const Port=8000

connectMongoDB()
connectDB()

app.use("/",userRoutes)
app.use("/api",cientRoute)

app.listen(`${Port}`,()=>console.log(`Server running on http://localhost:${Port}`))