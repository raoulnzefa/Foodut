import transport from './transport'

export default{
  async Login(email, password) {
    const response = await transport
      .post('/loginUser', {
        email: email,
        password: password
      })
    console.log(response.data);
    if(response.data.statusCode == 200){
      localStorage.setItem('userId', response.data.data)
      return true
    }else{
      return false
    }
  },
  async Logout() {
    const response = await transport
      .get('/logout')
      console.log(response.data)
      if(response.data.statusCode == 200){
        localStorage.setItem('userId', "")
        return true
      }else{
        return false
      }
  },
  async GetAllUser() {
    const response = await transport.get('/users')
    return response.data.data
  },
  async GetUserById(userId) {
    const response = await transport.get(`/users/${userId}`)
    return response.data.data
  },
  async UpdateGeneralUser(userId, profilePhoto, username, name) {
    const response = await transport
      .put(`/users/${userId}`, {
        profilePhoto: profilePhoto,
        username: username,
        name: name
      })
    console.log('Update General User')
    console.log(response.data)
    if(response.data.statusCode == 200){
      return true
    }else{
      return false
    }
  },
  async UpdatePasswordUser(userId, newPassword) {
    const response = await transport
      .put(`/users/${userId}`, {
        password: newPassword
      })
    console.log('Update Password')
    console.log(response.data)
    if(response.data.statusCode == 200){
      return true
    }else{
      return false
    }
  },
  async RegisterAdmin(username, email, name, password) {
    const response = await transport
      .post('/admin', {
        username: username,
        email: email,
        name: name,
        password: password
      })
      console.log('Admin')
      console.log(response.data)
      if(response.data.statusCode == 201){
        return true
      }else{
        return false
      }
  },
  async RegisterCustomer(username, email, name, password, address) {
    const response = await transport
      .post('/customer', {
        username: username,
        email: email,
        name: name,
        password: password,
        address: address
      })
      console.log('Customer')
      console.log(response.data)
      if(response.data.statusCode == 201){
        return true
      }else{
        return false
      }
  },
  async GetCustomerWithAssociation(userId) {
    const response = await transport.get(`/customers?user_id=${userId}`)
    return response.data.data
  },
  async GetAllCustomerWithAssociation() {
    const response = await transport.get('/customers')
    return response.data.data
  },
  async GetAllCustomerComplete() {
    const response = await transport.get('/customers-complete')
    return response.data.data
  },
  async RegisterSeller(username, email, name, password, address, storeName, city) {
    const response = await transport
    .post('/seller', {
      username: username,
      email: email,
      name: name,
      password: password,
      address: address,
      storeName: storeName,
      city: city
    })
    console.log('Seller')
    console.log(response.data)
    if(response.data.statusCode == 201){
      return true
    }else{
      return false
    }
  },
  async GetAllStore() {
    const response = await transport.get('/sellers')
    return response.data.data
  },
  async GetStoreByIdWithProduct(userId) {
    const response = await transport.get(`/store?userId=${userId}`)
    return response.data.data
  }
}