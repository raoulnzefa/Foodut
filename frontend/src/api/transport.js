import axios from 'axios'

const transport = axios.create({
  baseURL: 'http://localhost:1234',
  withCredentials: true
})

export default transport