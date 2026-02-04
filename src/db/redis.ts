import {createClient} from "redis"

export const connectRedis = createClient({
  url: "redis://127.0.0.1:6379",
});

connectRedis.on("error", (err) => console.log("Redis Client Error", err));

(async () => {
  if (!connectRedis.isOpen) {
    await connectRedis.connect();
    console.log("Redis connected");
  }
})();
