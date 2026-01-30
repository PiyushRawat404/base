"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const mongoose_1 = __importDefault(require("mongoose"));
const user_routes_1 = __importDefault(require("./routes/user.routes"));
const axios_1 = __importDefault(require("axios"));
const redis_1 = __importDefault(require("redis"));
const redisClient = redis_1.default.createClient();
const app = (0, express_1.default)();
app.use(express_1.default.json());
app.use(express_1.default.urlencoded({ extended: true }));
const Port = 8000;
mongoose_1.default.connect("mongodb://localhost:27017/healthcheckup")
    .then(() => console.log('Connected to MongoDB'))
    .catch((error) => {
    console.error('Error connecting to MongoDB:', error.message);
});
app.use("/", user_routes_1.default);
app.get("/api/todos", async (req, res) => {
    const { data } = await axios_1.default.get("https://jsonplaceholder.typicode.com/todos");
    return res.json(data);
});
app.get("/api/v1", (req, res) => {
    res.send("Hello from TypeScript");
});
app.listen(`${Port}`, () => console.log(`Server running on http://localhost:${Port}`));
//# sourceMappingURL=app.js.map