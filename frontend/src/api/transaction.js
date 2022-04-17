import transport from './transport'

export default{
  async AddTransaction(customerId, extraAddress, paymentOption) {
    const response = await transport
      .post('/transactions', {
        customerId: customerId,
        extraAddress: extraAddress,
        paymentOption: paymentOption
      })
    console.log('Transaction')
    console.log(response.data)
    if(response.data.statusCode == 201){
      return true
    }else{
      return false
    }
  },
  async GetAllTransaction(){
    const response = await transport.get('/transactions')
    return response.data.data
  },
  async GetAllTransactionDecentralized(){
    const response = await transport.get('/transactions-d')
    return response.data.data
  },
  async GetCustomerTransaction(customerId){
    const response = await transport.get(`/transactions-c?customerId=${customerId}`)
    return response.data.data
  },
  async GetCustomerTransactionDecentralized(customerId){
    const response = await transport.get(`/transactions-cd?customerId=${customerId}`)
    return response.data.data
  },
  async GetAllOrders(sellerId){
    const response = await transport.get(`/orders?sellerId=${sellerId}`)
    return response.data.data
  }
}