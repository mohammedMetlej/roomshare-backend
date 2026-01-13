import { createRoom } from "./api.js";

const form = document.getElementById("roomForm");

if (form) {
    form.addEventListener("submit", async (e) => {
        e.preventDefault();

        const roomData = {
            owner_id: Number(document.getElementById("owner_id").value),
            location: document.getElementById("location").value,
            price: Number(document.getElementById("price").value),
            capacity: Number(document.getElementById("capacity").value),
            description: document.getElementById("description").value
        };

        try {
            await createRoom(roomData);
            alert("Room created successfully");
            form.reset();
        } catch (err) {
            alert(err.message || "Failed to create room");
        }
    });
}
