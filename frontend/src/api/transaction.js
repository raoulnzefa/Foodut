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
  async GetAllOrders(sellerId){
    const response = await transport.get(`/orders?sellerId=${sellerId}`)
    return response.data.data
  },
  // async AddCart(userId, products[0].productId, products[].quantity) {
  //   const response = await transport
  //     .post('/cart', {
  //       userId: userId,
  //       prproductId: productId, 
  //       quantity: 5
  //     })
  //     console.log('Cart')
  //     console.log(response.data)
  //     if(response.data.statusCode == 201){
  //       return true
  //     }else{
  //       return false
  //     }
  // },
  // async GetCart(){
  //   const response = await transport.get('/cart-a')
  //   return response.data.data
  // }
}