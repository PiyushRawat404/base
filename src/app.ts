import express from "express";
import http from "http"
import { connectMongoDB } from "./db/mongoose";
import userRoutes from "./routes/user.routes"
import cientRoute from "./routes/client.routes"
import { connectDB } from "./db/postgresql";
import { bridge } from "./middleware/user";

const app=express()
const server=http.createServer(app)
app.use(express.json())
app.use(express.urlencoded({ extended: true }));
const Port=8000

connectMongoDB()
connectDB()
app.use(bridge)
app.use("/",userRoutes)
app.use("/api",cientRoute)

server.listen(`${Port}`,()=>console.log(`Server running on http://localhost:${Port}`))

app.listen(`${Port}`,()=>console.log(`Server running on http://localhost:${Port}`))