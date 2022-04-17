import transport from './transport'

export default{
  async AddCart(userId, products) {
    const response = await transport
      .post('/cart', {
        userId: userId,
        products: products
      })
    console.log('Add Cart')
    console.log(response.data)
    if(response.data.statusCode == 201){
      return true
    }else{
      return false
    }
  },
  async GetCart(userId) {
    const response = await transport.get(`/cart-a?customerId=${userId}`)
    console.log('Cart')
    console.log(response.data)
    return response.data.data
  },
  async UpdateCart(userId, products) {
    const response = await transport
      .put('/cart', {
        userId: userId,
        products: products
      })
    console.log('Update Cart')
    console.log(response.data)
    if(response.data.statusCode == 201){
      return true
    }else{
      return false
    }
  },
  async DeleteSpecificCart(userId, productId) {
    const response = await transport
      .delete('/cart-specific', {
        data: {
          userId: userId,
          productId: productId,
        }
      })
    console.log('Delete Spesific Cart')
    console.log(response.data)
    if(response.data.statusCode == 201){
      return true
    }else{
      return false
    }
  },
  async DeleteCarts(productsCartUser) {
    const response = await transport.delete('/cart', { data: productsCartUser })
    console.log('Delete All Products in Cart')
    console.log(response.data)
    if(response.data.statusCode == 201){
      return true
    }else{
      return false
    }
  }
}