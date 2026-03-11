const inputBox = document.getElementById("input-box");
const listContainer = document.getElementById("list-container");

const completedCounter = document.getElementById("completed-counter");
const uncompletedCounter = document.getElementById("uncompleted-counter");


function addTask() {
    const task = inputBox.value.trim();
    if (!task) {
        alert("Enter a task");
        return;
    }
    const li = document.createElement("li");
    li.innerHTML = `
        <label>
            <input type="checkbox">
            <span>${task}</span>
        </label>
        <span class="delete-button">Delete</span>
    `;
    addTaskEvents(li);
    listContainer.appendChild(li);
    inputBox.value = "";
    updateCounters();
    saveTasks();
}



function addTaskEvents(li) {
    const checkbox = li.querySelector("input");
    const deleteBtn = li.querySelector(".delete-button");
    checkbox.addEventListener("click", function () {
         li.classList.toggle("completed", checkbox.checked);
          updateCounters();
          saveTasks();

    });
    deleteBtn.addEventListener("click", function () {
        li.remove();
        updateCounters();
        saveTasks();


    }
);

}
function updateCounters() {
    const totalTasks = document.querySelectorAll("#list-container li").length;
    const completedTasks = document.querySelectorAll(".completed").length;
    const uncompletedTasks = totalTasks - completedTasks;
    completedCounter.textContent = completedTasks;
      uncompletedCounter.textContent = uncompletedTasks;
}
function saveTasks() {
    localStorage.setItem("tasks", listContainer.innerHTML);
}
function loadTasks() {
    const savedTasks = localStorage.getItem("tasks");
    if (savedTasks) {
          listContainer.innerHTML = savedTasks;
        const tasks = listContainer.querySelectorAll("li");
        for (let i = 0; i < tasks.length; i++) {
       addTaskEvents(tasks[i]);

        }
    }
}
loadTasks();
updateCounters();