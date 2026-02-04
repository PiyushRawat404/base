import {Client} from "pg"
import "dotenv/config"

const client =new Client({
    host: "localhost",
    user: "postgres",
    port:5432,
    password:process.env.PG_PASSWORD,
    database: "postgres"
})

export const connectDB =() => {
    client.connect()
    .then(()=>console.log("PostgreSQL connected"))
    .catch((err)=>console.error("Postgres error:", err));
};

export default client
