"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.newUser = exports.getUser = void 0;
const user_1 = require("../model/user");
const getUser = async (req, res) => {
    const name = req.params.name;
    await user_1.User.find({ name });
    res.send(name);
};
exports.getUser = getUser;
const newUser = async (req, res) => {
    try {
        const { name, num, gender, problem } = req.body;
        if (!name || !num || !gender || !problem) {
            return res.status(400).json({ msg: "All fields are required" });
        }
        const user = await user_1.User.create({
            clientNo: Date.now(),
            fullName: name,
            phoneNo: num,
            gender: gender,
            problem: problem
        });
        return res.status(201).json({
            msg: "User created successfully",
            user,
        });
    }
    catch (error) {
        console.error(error);
        return res.status(500).json({ msg: "Server error" });
    }
};
exports.newUser = newUser;
//# sourceMappingURL=user.js.map