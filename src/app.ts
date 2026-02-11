import express from "express";
// import http from "http"
import { connectMongoDB } from "./db/mongoose";
import pracRoutes from "./routes/prac.routes"
import loggedRoute from "./routes/logged.routes"
import { connectDB } from "./db/postgresql";
// import { bridge } from "./middleware/user";

const app=express()
// const server=http.createServer(app)
app.use(express.json())
app.use(express.urlencoded({ extended: true }));
const Port=8000


connectMongoDB()
connectDB()

// app.use(bridge)
app.use("/prac",pracRoutes)
app.use("/",loggedRoute)

// server.listen(`${Port}`,()=>console.log(`Server running on http://localhost:${Port}`))

app.listen(`${Port}`,()=>console.log(`Server running on http://localhost:${Port}`))