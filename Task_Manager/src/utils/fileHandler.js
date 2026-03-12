const fs = require("fs").promises;
const path = require("path");

const filePath = path.join(__dirname,"../../data/output.json");
async function readData() {
    try {
        const data = await fs.readFile(filePath, "utf-8");

        if (!data) {
            return [];
        }

        return JSON.parse(data);

    } catch (error) {
        return [];
    }
}

async function writeData(data) {
    try {
        await fs.writeFile(
            filePath,
            JSON.stringify(data, null, 2),
            "utf-8"
        );
    } catch (error) {
        console.error("Error writing file:", error);
    }
}
module.exports = {readData,writeData};