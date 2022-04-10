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
  async AddAdmin(username, email, name, password) {
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
  async AddCustomer(username, email, name, password, address) {
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
  async GetAllCustomerWithAssociation() {
    const response = await transport.get('/customers')
    return response.data.data
  },
  async GetAllCustomerComplete() {
    const response = await transport.get('/customers-complete')
    return response.data.data
  },
  async AddSeller(username, email, name, password, address, storeName, city) {
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