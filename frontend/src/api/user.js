import transport from './transport'

export default{
  async Login (email, password) {
    const response = await transport
      .post('/loginUser', {
        email: email,
        password: password
      })
    console.log(response.data);
    if(response.data.statusCode == 200){
      localStorage.setItem('userId', response.data.data)
      console.log(localStorage.getItem('userId'))
      return true
    }else{
      return false
    }
  },
  async Logout () {
    const response = await transport
      .get('/logout')
      console.log(response.data)
      if(response.data.statusCode == 200){
        return true
      }else{
        return false
      }
  },
  async GetAllUser () {
    const response = await transport.get('/users')
    return response.data.data
  },
  async GetUserById (userId) {
    const response = await transport.get(`/users?user_id=${userId}`)
    return response.data.data
  }
}