import {fetchRooms} from "./api.js";

const roomsContainer = document.getElementById("rooms");

async function loadRooms(){
    try{
        const rooms = await fetchRooms();

        if(rooms.length == 0){
            roomsContainer.innerHTML = "<p>No rooms available. <p>";
            return;
        }

        roomsContainer.innerHTML = rooms
        .map(
            (room) => `
            <div class= "room-card">
            <h3> ${room.location}</h3>
            <p>  ${room.description}</p>
            <p> Price: ${room.price}</p>
            <p> Capacity: ${room.capacity}</p>
            </div>`

        )
        .join("");
    }catch (error){
        roomsContainer.innerHTML = "<p>Error loading rooms </p>";
    }
}

loadRooms();