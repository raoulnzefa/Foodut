import transport from './transport'

export default{
  async InsertCategories(productCategory){
    const response = await transport
      .post('/categories', {
        productCategory: productCategory
      })
  },
  async InsertProduct(productName, productPrice, productStock, sellerId, productCategory, productPicture){
    const response = await transport
      .post('/products', {
        productName: productName,
        productPrice: productPrice,
        productStock: productStock,
        sellerId: sellerId,
        productCategory: productCategory,
        productPicture: productPicture
      })
      console.log('Product')
      console.log(response.data)
  },
  async GetAllProduct(){
    const response = await transport.get('/products')
    console.log('get all products')
    console.log(response.data.data)
    return response.data.data
  },
  async GetAllProductByStore(){
    const response = await transport.get(`/store/${id}`)
    return response.data.data
  }
}