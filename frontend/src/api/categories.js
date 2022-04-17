import transport from './transport'

export default{
  async AddCategories(productCategory){
    const response = await transport
      .post('/categories', {
        productCategory: productCategory
      })
      console.log('Add Categories')
      console.log(response.data)
      if(response.data.statusCode == 201){
        return true
      }else{
        return false
      }
  },
  async UpdateCategories(oldProductCategory, newProductCategory){
    const response = await transport
    .put(`/categories/${oldProductCategory}`, {
      productCategory: newProductCategory
    })
    console.log('Update Categories')
    console.log(response.data)
    if(response.data.statusCode == 201){
      return true
    }else{
      return false
    }
  }
}