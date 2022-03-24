import axios from 'axios'

export default{
  async Login () {
    axios
      .post('http://localhost:1234/loginUser', {
        email: 'fedlysell@gmail.com',
        password: 'fedly'
      })
      .then(response => {
        console.log('Succes login')
        console.log(response)
        console.log('---')
      })
      .catch(error => {
        console.log('Error login')
        console.log(error)
        console.log('---')
      })
  },
  async Logout () {
    axios
      .get('http://localhost:1234/logout')
      .then(response => {
        console.log('Succes logout')
        console.log(response)
        console.log('---')
      })
      .catch(error => {
        console.log('Error logout')
        console.log(error)
        console.log('---')
      })
  },
  async GetAllUser () {
    axios
      .get('http://localhost:1234/users')
      .then(response => {
        console.log('Succes get all users')
        console.log(response)
        console.log('---')
      })
      .catch(error => {
        console.log('Error get all users')
        console.log(error)
        console.log('---')
      })
  }
}