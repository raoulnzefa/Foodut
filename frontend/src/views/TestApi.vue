<template>
    <div>
        Categories <br>
        <vs-button class="ml-auto mt-2" @click="AddCategories">Add Categories</vs-button><br>
        <vs-button class="ml-auto mt-2" @click="UpdateCategories">Update Categories</vs-button><br>
        Product <br>
        <vs-button class="ml-auto mt-2" @click="AddProduct">Add Product</vs-button><br>
        <vs-button class="ml-auto mt-2" @click="UpdateProduct">Update Product</vs-button><br>
        <vs-button class="ml-auto mt-2" @click="DeleteProduct">Delete Product</vs-button> belum<br>
        <vs-button class="ml-auto mt-2" @click="GetAllProduct">Get All Product</vs-button> ✔<br>
        <vs-button class="ml-auto mt-2" @click="GetProductByName">Get Product By Name</vs-button><br>
        <vs-button class="ml-auto mt-2" @click="GetProductByCategoryId">Get Product By Category Id</vs-button><br>
        <vs-button class="ml-auto mt-2" @click="GetProductByCategoryName">Get Product By Category Name</vs-button><br>
        User <br>
        <vs-button class="ml-auto mt-2" @click="LoginAdmin">Login Admin</vs-button> ✔<br>
        <vs-button class="ml-auto mt-2" @click="LoginCustomer">Login Customer</vs-button> ✔<br>
        <vs-button class="ml-auto mt-2" @click="LoginSeller">Login Seller</vs-button> ✔<br>
        <vs-button class="ml-auto mt-2" @click="Logout">Logout</vs-button> ✔<br>
        <vs-button class="ml-auto mt-2" @click="GetAllUser">Get All User</vs-button> ✔<br>
        <vs-button class="ml-auto mt-2" @click="GetUserById">Get User By Id</vs-button> ✔<br>
        Admin <br>
        <vs-button class="ml-auto mt-2" @click="RegisterAdmin">Register Admin</vs-button> ✔<br>
        Customer <br>
        <vs-button class="ml-auto mt-2" @click="RegisterCustomer">Register Customer</vs-button> ✔<br>
        <vs-button class="ml-auto mt-2" @click="GetAllCustomerWithAssociation">Get All Customer With Association</vs-button><br>
        <vs-button class="ml-auto mt-2" @click="GetAllCustomerComplete">Get All Customer Complete</vs-button><br>
        Seller / Store <br>
        <vs-button class="ml-auto mt-2" @click="RegisterSeller">Register Seller</vs-button> ✔<br>
        <vs-button class="ml-auto mt-2" @click="GetAllStore">Get All Store</vs-button> ✔<br>
        <vs-button class="ml-auto mt-2" @click="GetStoreByIdWithProduct">Get Store By Id With Product</vs-button><br>
        Cart <br>
        <vs-button class="ml-auto mt-2" @click="AddCart">Add Cart</vs-button><br>
        <vs-button class="ml-auto mt-2" @click="UpdateCart">Update Cart</vs-button><br>
        <vs-button class="ml-auto mt-2" @click="DeleteSpecificCart">Delete Specific Cart</vs-button><br>
        <vs-button class="ml-auto mt-2" @click="DeleteCarts">Delete Carts</vs-button><br>
        <vs-button class="ml-auto mt-2" @click="GetCart">Get Cart</vs-button><br>
        Transaction <br>
        <vs-button class="ml-auto mt-2" @click="AddTransaction">Add Transaction</vs-button><br>
        <vs-button class="ml-auto mt-2" @click="GetAllTransaction">Get All Transaction</vs-button><br>
        <vs-button class="ml-auto mt-2" @click="GetAllTransactionDecentralized">Get All Transaction Decentralize</vs-button><br>
        <vs-button class="ml-auto mt-2" @click="GetCustomerTransaction">Get Customer Transaction</vs-button><br>
        <vs-button class="ml-auto mt-2" @click="GetCustomerTransactionDecentralized">Get Customer Transaction Decentralized</vs-button> ✔<br>
        <vs-button class="ml-auto mt-2" @click="GetAllOrder">Get All Orders</vs-button> ✔<br>
    </div>
</template>

<script>
import apiCategories from '../api/categories'
import apiProduct from '../api/product'
import apiCart from '../api/cart';
import apiTransaction from '../api/transaction'
import apiUser from '../api/user'

export default {
  data () {
    return {}
  },
  methods: {
    //Categories
    AddCategories() {
      const productCategory = "Bakso Frozen"
      apiCategories
        .AddCategories(productCategory)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to create categories',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to create categories',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    UpdateCategories() {
      const oldProductCategory = "frozenfood"
      const newProductCategory = "Permen"
      apiCategories
        .UpdateCategories(oldProductCategory, newProductCategory)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to Update categories',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to Update categories',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    //Product
    AddProduct() {
      const productName = "Chips Chups 2"
      const productPrice = 10000
      const productStock = 99
      const sellerId = 3
      const productCategory = 1
      const productDescription = "telah hadir keripik makaroni jos dengan perpaduan rempah-rempah hadirkan keripik kentang yang enak gurih dan nikmat. Ayo segera dapatkan sekarang juga!"
      const productPicture = [
        "chip1.jpg",
        "chip2.jpg",
        "chip3.jpg"
      ]
      apiProduct
        .AddProduct(productName, productPrice, productStock, sellerId, productCategory, productDescription, productPicture)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to create product',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to create product',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    UpdateProduct() {
      const productId = 3
      const productName = "Chips Chups 111"
      const productPrice = 10000
      const productStock = 33
      const productCategory = 1
      const productDescription = "telah hadir keripik makaroni jos dengan perpaduan rempah-rempah hadirkan keripik kentang yang enak gurih dan nikmat. Ayo segera dapatkan sekarang juga!"
      apiProduct
        .UpdateProduct(productId, productName, productPrice, productStock, productCategory, productDescription)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to update product',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to update product',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    DeleteProduct() {
      const productId = 6
      apiProduct
        .DeleteProduct(productId)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to delete product',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to delete product',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    GetAllProduct() {
      apiProduct
        .GetAllProduct()
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to get all product',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to get all product',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
            console.log('Data Product')
            console.log(response)
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    GetProductByName() {
      const productName = "Chips Chups"
      apiProduct
        .GetProductByName(productName)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to get product',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to get product',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
            console.log('Data Product')
            console.log(response)
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    GetProductByCategoryId() {
      const categoryId = 1
      apiProduct
        .GetProductByCategoryId(categoryId)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to get product',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to get product',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
            console.log('Data Product')
            console.log(response)
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    GetProductByCategoryName() {
      const categoryName = "Keripik"
      apiProduct
        .GetProductByCategoryName(categoryName)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to get product',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to get product',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
            console.log('Data Product')
            console.log(response)
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    //User
    LoginAdmin() {
      const email = "admin@gmail.com"
      const password = "12345678"
      apiUser
        .Login(email, password)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to login',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to login',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    LoginCustomer() {
      const email = "elangel@gmail.com"
      const password = "12345678"
      apiUser
        .Login(email, password)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to login',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to login',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    LoginSeller() {
      const email = "fedly@gmail.com"
      const password = "12345678"
      apiUser
        .Login(email, password)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to login',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to login',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    Logout() {
      apiUser
        .Logout()
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to logout',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to logout',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    GetAllUser() {
      apiUser
        .GetAllUser()
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to get all user',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to get all user',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
            console.log('Data User')
            console.log(response)
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    GetUserById() {
      apiUser
        .GetUserById(localStorage.getItem('userId'))
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to get all user',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to get all user',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
            console.log('Data User')
            console.log(response)
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    //Admin
    RegisterAdmin() {
      const username = "admin4"
      const email = "admin4@gmail.com"
      const name = "admin"
      const password = "12345678"
      apiUser
        .RegisterAdmin(username, email, name, password)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to create admin account',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to create admin account',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    //Customer
    RegisterCustomer() {
      const username = "kimso12"
      const email = "kimso12@gmail.com"
      const name = "kimso kimi"
      const password = "12345678"
      const address = "Jl. Bubur Ayam Enak no. 8"
      apiUser
        .RegisterCustomer(username, email, name, password, address)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to create customer account',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to create customer account',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    GetAllCustomerWithAssociation() {
      apiUser
        .GetAllCustomerWithAssociation()
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to get all store',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to get all store',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
            console.log('Data Store')
            console.log(response)
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    GetAllCustomerComplete() {
      apiUser
        .GetAllCustomerComplete()
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to get all store',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to get all store',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
            console.log('Data Store')
            console.log(response)
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    //Seller
    RegisterSeller() {
      const username = "kimso11"
      const email = "kimso11@gmail.com"
      const name = "kiminonawa"
      const password = "123456781"
      const address = "Jl. Bubur Ayam Enak no. 8"
      const storeName = "Makaroni 123"
      const city = "Jakarta"
      apiUser
        .RegisterSeller(username, email, name, password, address, storeName, city)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to create seller account',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to create seller account',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    GetAllStore() {
      apiUser
        .GetAllStore()
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to get all store',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to get all store',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
            console.log('Data Store')
            console.log(response)
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    GetStoreByIdWithProduct() {
      const userId = 3
      apiUser
        .GetStoreByIdWithProduct(userId)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to get all store',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to get all store',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
            console.log('Data Store')
            console.log(response)
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    //Cart
    AddCart() {
      const userId = 2
      const products = [
        {productId: 3, quantity: 2},
        {productId: 5, quantity: 5}
      ]
      apiCart
        .AddCart(userId, products)       
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to create cart',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to create cart',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    GetCart() {
      const userId = 2
      apiCart
        .GetCart(userId)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to get all cart',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to get all cart',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    UpdateCart() {
      const userId = 2
      const products = [
        {productId: 3, quantity: 12},
        {productId: 5, quantity: 15}
      ]
      apiCart
        .UpdateCart(userId, products)  
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to update cart',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to update cart',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    DeleteSpecificCart() {
      const userId = 2
      const productId = 1
      apiCart
        .DeleteSpecificCart(userId, productId)  
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to delete cart',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to delete cart',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    DeleteCarts() {
      const productsCartUser = [
        {
            "customerId": 2,
            "productId": 2
        },
        {
            "customerId": 2,
            "productId": 3
        }
      ]
      apiCart
        .DeleteCarts(productsCartUser)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to delete all products in cart',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to delete all products in cart',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    //Transaction
    AddTransaction() {
      const customerId = 2
      const extraAddress = "Semarang"
      const paymentOption = "ATM"
      apiTransaction
        .AddTransaction(customerId, extraAddress, paymentOption)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to create transaction',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to create transaction',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    GetAllTransaction() {
      apiTransaction
        .GetAllTransaction()
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to get all transaction',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to get all transaction',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
            console.log('Data Transaction')
            console.log(response)
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    GetAllTransactionDecentralized() {
      apiTransaction
        .GetAllTransactionDecentralized()
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to get all transaction decentralized',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to get all transaction decentralized',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
            console.log('Data Transaction')
            console.log(response)
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    GetCustomerTransaction() {
      const customerId = 2
      apiTransaction
        .GetCustomerTransaction(customerId)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to get customer transaction',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to get customer transaction',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
            console.log('Data Transaction')
            console.log(response)
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    GetCustomerTransactionDecentralized() {
      const customerId = 2
      apiTransaction
        .GetCustomerTransactionDecentralized(customerId)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to get customer transaction decentralized',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to get customer transaction decentralized',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
            console.log('Data Transaction')
            console.log(response)
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    GetAllOrder() {
      const sellerId = 3
      apiTransaction
        .GetAllOrders(sellerId)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to get all order',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to get all order',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
            console.log('Data Order')
            console.log(response)
          }
        })
        .catch((error) => {          
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    }
  },
  components: {}
}
</script>
