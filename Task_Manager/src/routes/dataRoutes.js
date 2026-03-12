const express = require("express");
const { getAllTask,createTask,updateSingleTask,deleteTask } = require("../controllers/dataController");
const router = express.Router();

router.get("/health", (req,res) =>{
    res.status(200).json({message : "The connections are working"});
});

//All read 
router.get("/tasks",getAllTask);

//create

router.post("/task",createTask);
//update


router.put("/task/:id",updateSingleTask);


//del

router.delete("/task/:id",deleteTask);

module.exports = router;