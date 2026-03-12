
const fileHandler = require("../utils/fileHandler");



async function getAllTask(){
    const task = await fileHandler.readData();
    return task;

};

async function createTask(taskData){
    const task = await fileHandler.readData();
    
    const newTask = {
        id : Date.now().toString(),
        title : taskData.title,
        completed : false
    };
    task.push(newTask);
    await fileHandler.writeData(task);
    return newTask;

};

async function updateSingleTask(taskId,updatedData){
    const task = await fileHandler.readData();
    const findTask = task.find(t => t.id === taskId);

    if(!findTask){
        console.log("Error");
        throw new Error("Task not found");
    }
     Object.assign(findTask,updatedData);
    await fileHandler.writeData(task);
    return task;

    
};

async function deleteTask(taskId){
    const task = await fileHandler.readData();
    const filteredTask = [];
    for(let i =0 ;i <task.length;i++){
        if(task[i].id !== taskId){
            filteredTask.push(task[i]);
        }
    }
    await fileHandler.writeData(filteredTask);
    return true;

};

module.exports = {getAllTask,createTask,updateSingleTask,deleteTask};