import transport from './transport'

export default{
  async AddProduct(productName, productPrice, productStock, sellerId, productCategory, productDescription, productPicture){
    const response = await transport
      .post('/products', {
        productName: productName,
        productPrice: parseInt(productPrice),
        productStock: parseInt(productStock),
        sellerId: parseInt(sellerId),
        productCategory: parseInt(productCategory),
        productDescription: productDescription,
        productPicture: productPicture
      })
      console.log('Add Product')
      console.log(response.data)
      if(response.data.statusCode == 201){
        return true
      }else{
        return false
      }
  },
  async UpdateProduct(productId, productName, productPrice, productStock, productCategory, productDescription){
    const response = await transport
      .put(`/products/${productId}`, {
        productName: productName,
        productPrice: parseInt(productPrice),
        productStock: parseInt(productStock),
        productCategory: parseInt(productCategory),
        productDescription: productDescription
      })
      console.log('Update Product')
      console.log(response.data)
      if(response.data.statusCode == 200){
        return true
      }else{
        return false
      }
  },
  async DeleteProduct(productId){
    const response = await transport
      .delete(`/products/${productId}`)
      console.log('Delete Product')
      console.log(response.data)
      if(response.data.statusCode == 201){
        return true
      }else{
        return false
      }
  },
  async GetAllProduct(){
    const response = await transport.get('/products')
    return response.data.data
  },
  async GetProductById(productId){
    const response = await transport.get(`/products?id=${productId}`)
    return response.data.data
  },
  async GetProductByName(productName){
    const response = await transport.get(`/products-by-name?product_name=${productName}`)
    return response.data.data
  },
  async GetProductByCategoryId(categoryId){
    const response = await transport.get(`/products-by-category-id?category_id=${categoryId}`)
    return response.data.data
  },
  async GetProductByCategoryName(categoryName){
    const response = await transport.get(`/products-by-category-name?category_name=${categoryName}`)
    return response.data.data
  }
}