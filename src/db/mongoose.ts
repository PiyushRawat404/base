import mongoose from 'mongoose';

export const connectMongoDB=()=>mongoose.connect("mongodb://localhost:27017/healthcheckup")
.then(() => console.log('Connected to MongoDB'))
.catch((error) => {
    console.error('Error connecting to MongoDB:', error.message);
});



