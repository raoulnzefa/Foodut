import transport from './transport'

export default{
  async AddCategories(productCategory){
    const response = await transport
      .post('/categories', {
        productCategory: productCategory
      })
      console.log('Categories')
      console.log(response.data)
      if(response.data.statusCode == 201){
        return true
      }else{
        return false
      }
  },
  async AddProduct(productName, productPrice, productStock, sellerId, productCategory, productDescription, productPicture){
    const response = await transport
      .post('/products', {
        productName: productName,
        productPrice: productPrice,
        productStock: productStock,
        sellerId: sellerId,
        productCategory: productCategory,
        productDescription: productDescription,
        productPicture: productPicture
      })
      console.log('Product')
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
  //belum
  // async GetAllProductByStore(){
  //   const response = await transport.get(`/store/?userId=${id}`)
  //   return response.data.data
  // }
}