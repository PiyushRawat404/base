const dataService = require("../services/dataService");

async function getAllTask(req , res){
    try {
        const tasks = await dataService.getAllTask();
        res.status(200).json({
            success : true,
            data : tasks,
        });
        
    } catch (error) {
        res.status(500).json({
            success : false,
            message : "Failed to fetch tasks"
        });
        
    }
};

async function createTask(req , res){
    try {
        const taskData = req.body;

        const newTask = await dataService.createTask(taskData);
        res.status(201).json({
            success : true,
            data : newTask,
        });
        
        
    } catch (error) {
        res.status(500).json({
            success : false,
            message : "Failed to create tasks"
        });
        
    }
};

async function updateSingleTask(req , res){
   try {
    const taskId = req.params.id;
    const updatedData = req.body;
    const updatedTask = await dataService.updateSingleTask(taskId,updatedData);
    res.status(200).json({
        success : true,
        data : updatedTask,
    });
    
   } catch (error) {

    res.status(500).json({
        sucess : false,
        message : "Failed to update task"
    });
    
   }
};

async function deleteTask(req , res){
    try {
        const taskId = req.params.id;
        await dataService.deleteTask(taskId);
        res.status(200).json({
            success : true,
            message : "Successfully deleted task"
        });
        
    } catch (error) {
        res.status(500).json({
            success : false,
            message : "Failed to delete Task",
        });
        
    }
};







module.exports = {getAllTask,createTask,updateSingleTask,deleteTask};