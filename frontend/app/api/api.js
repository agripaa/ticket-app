import axios from "axios";

const headers = {
    accept: 'application/json',
    'Content-Type': 'application/json'
}

const api = axios.create({
    baseURL: `http://localhost:8080/api/ticket`,
    headers
})

export default api;