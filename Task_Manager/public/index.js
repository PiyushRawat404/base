const inputBox = document.getElementById("input-box");
const listContainer = document.getElementById("list-container");

const completedCounter = document.getElementById("completed-counter");
const uncompletedCounter = document.getElementById("uncompleted-counter");

const API_URL = "/api";

async function loadTasks() {
    try {

        const response = await fetch(`${API_URL}/tasks`);
        const result = await response.json();

        const tasks = result.data;

        listContainer.innerHTML = "";

        tasks.forEach(task => renderTask(task));

        updateCounters();

    } catch (error) {

        console.error("Error loading tasks:", error);

    }
}

async function addTask() {

    const title = inputBox.value.trim();

    if (!title) {
        alert("Enter a task");
        return;
    }

    try {

        const response = await fetch(`${API_URL}/task`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ title })
        });

        const result = await response.json();
        const newTask = result.data || result;

        renderTask(newTask);

        inputBox.value = "";

        updateCounters();

    } catch (error) {

        console.error("Error creating task:", error);

    }
}

function renderTask(task) {

    const li = document.createElement("li");

    li.innerHTML = `
        <label>
            <input type="checkbox" ${task.completed ? "checked" : ""}>
            <span>${task.title}</span>
        </label>
        <span class="delete-button">🗑</span>
        <hr/>
    `;

    const checkbox = li.querySelector("input");
    const deleteBtn = li.querySelector(".delete-button");

    if (task.completed) {

        li.classList.add("completed");

        listContainer.appendChild(li);

    } else {

        const firstCompleted = listContainer.querySelector(".completed");

        if (firstCompleted) {
            listContainer.insertBefore(li, firstCompleted);
        } else {
            listContainer.prepend(li);
        }

    }

    checkbox.addEventListener("change", async () => {

        try {

            await fetch(`${API_URL}/task/${task.id}`, {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    completed: checkbox.checked
                })
            });

            li.classList.toggle("completed", checkbox.checked);

            if (checkbox.checked) {

                listContainer.appendChild(li);

            } else {

                const firstCompleted = listContainer.querySelector(".completed");

                if (firstCompleted) {
                    listContainer.insertBefore(li, firstCompleted);
                } else {
                    listContainer.prepend(li);
                }

            }

            updateCounters();

        } catch (error) {

            console.error("Error updating task:", error);

        }

    });

    deleteBtn.addEventListener("click", async () => {

        try {

            await fetch(`${API_URL}/task/${task.id}`, {
                method: "DELETE"
            });

            li.remove();

            updateCounters();

        } catch (error) {

            console.error("Error deleting task:", error);

        }

    });

}

function updateCounters() {

    const totalTasks = document.querySelectorAll("#list-container li").length;
    const completedTasks = document.querySelectorAll(".completed").length;

    const uncompletedTasks = totalTasks - completedTasks;

    completedCounter.textContent = completedTasks;
    uncompletedCounter.textContent = uncompletedTasks;

}

inputBox.addEventListener("keydown", function (e) {

    if (e.key === "Enter") {
        addTask();
    }

});

loadTasks();