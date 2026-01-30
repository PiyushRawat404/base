import mongoose, { Schema } from "mongoose";

const userSchema= new Schema({
    clientNo:{
        type:Number,
        required:true,
        unique:true,
    },
    fullName:{
        type:String,
        required:true,
    },
    phoneNo:{
        type:Number,
        required:true,
    },
    gender:{
        type:String,
        required:true,
    },
    problem:{
        type:String,
        required:true,
    }
})
export const User= mongoose.model("User",userSchema)