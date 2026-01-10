const API_BASE_URL = "http://localhost:8000";


export async function fetchRooms(){
  const response = await fetch(`${API_BASE_URL}/rooms`);
  if(!response.ok){
    throw new Error("Failed to fetch rooms");
  }
  return response.json();
}