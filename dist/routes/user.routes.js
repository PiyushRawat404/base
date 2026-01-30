"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = require("express");
const user_1 = require("../controllers/user");
const router = (0, express_1.Router)();
router.get("/:name", user_1.getUser);
router.post("/register", user_1.newUser);
exports.default = router;
//# sourceMappingURL=user.routes.js.map