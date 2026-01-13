const API_URL = "http://localhost:8080/rooms";

export async function fetchRooms() {
    const res = await fetch(API_URL);
    if (!res.ok) throw new Error("Failed to fetch rooms");
    return res.json();
}

export async function createRoom(roomData) {
    const res = await fetch(API_URL, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(roomData),
    });

    if (!res.ok) throw new Error("Failed to create room");
    return res.json();
}
