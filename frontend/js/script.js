fetch("http://localhost:8080/rooms")
.then(response =>{
    if(!response.ok){
        throw new Error("Failed to fetch rooms")
    }
    return response.json();
})
.then(rooms =>{
    const container = document.getElementById("rooms");

    rooms.forEach(room =>{
        const div = document.createElement("div");

        div.innerHTML = `
        <h3> ${room.location}</h3>
        <p> Price ${room.price}</p>
        <p> Capacity ${room.capacity}</p>
        <p> ${room.description}</p>
        <hr/> `;

        container.appendChild(div);
    });
})

.catch(error=>{
    console.error(error);
});