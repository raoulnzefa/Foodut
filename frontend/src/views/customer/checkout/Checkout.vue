<!-- =========================================================================================
  File Name: ECommerceCheckout.vue
  Description: eCommerce Checkout page
  ----------------------------------------------------------------------------------------
  Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
  Author: Pixinvent
  Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->

<template>
    <div id="ecommerce-checkout-demo">
        <form-wizard
            ref="checkoutWizard"
            color="rgba(var(--vs-primary), 1)"
            :title="null"
            :subtitle="null"
            :hide-buttons="true">

            <!-- tab 1 content -->
            <tab-content title="Cart" icon="feather icon-shopping-cart" class="mb-5">

                <!-- IF CART HAVE ITEMS -->
                <div class="vx-row" v-if="carts.length">

                    <!-- LEFT COL -->
                    <div class="vx-col lg:w-2/3 w-full relative">
                        <div class="items-list-view" v-for="(item, index) in carts" :key="item.objectID">
                            <item-list-view :item="item" class="mb-base">

                                <!-- SLOT: ITEM META -->
                                <template slot="item-meta">
                                    <h6
                                      class="item-name font-semibold mb-1 cursor-pointer hover:text-primary"
                                      @click="$router.push({name: 'ecommerce-item-detail-view', params: {item_id: item.objectID }}).catch(() => {})">{{ item.product.productName }}</h6>
                                    <p class="text-sm mb-2">By <span class="font-semibold cursor-pointer hover:text-primary">{{ item.store.storeName }}</span></p>
                                    <p class="text-success text-sm">In Stock: {{ item.product.productStock }}</p>

                                    <p class="mt-4 font-bold text-sm">Quantity</p>
                                    <vs-input-number min="1" max="10" :value="item.quantity" @input="updateItemQuantity($event, index)" class="inline-flex" />
                                </template>

                                <!-- SLOT: ACTION BUTTONS -->
                                <template slot="action-buttons">

                                    <!-- PRIMARY BUTTON: REMOVE -->
                                    <div class="item-view-primary-action-btn p-3 rounded-lg flex flex-grow items-center justify-center cursor-pointer mb-3" @click="deleteSpesifikCart(item.productId)">
                                        <feather-icon icon="XIcon" svgClasses="h-4 w-4" />
                                        <span class="text-sm font-semibold ml-2">REMOVE</span>
                                    </div>
                                </template>
                            </item-list-view>
                        </div>
                    </div>

                    <!-- RIGHT COL -->
                    <div class="vx-col lg:w-1/3 w-full">
                        <vx-card>
                            <p class="font-semibold mb-3">Price Details</p>
                            <div class="flex justify-between mb-2" v-for="item in carts" :key="item.objectID">
                                <span class="text-grey">{{ item.product.productName }}</span>
                                <span>Rp {{ item.product.productPrice * item.quantity }}</span>
                            </div>
                            <vs-divider />

                            <div class="flex justify-between font-semibold mb-3">
                                <span>Total</span>
                                <span>Rp {{ totalCart }}</span>
                            </div>

                            <vs-button class="w-full" @click="$refs.checkoutWizard.nextTab()">PLACE ORDER</vs-button>
                        </vx-card>
                    </div>
                </div>

                <!-- IF NO ITEMS IN CART -->
                <vx-card title="You don't have any items in your cart." v-else>
                    <vs-button @click="$router.push('/customer/browse').catch(() => {})">Browse Shop</vs-button>
                </vx-card>

            </tab-content>

            <!-- tab 2 content -->
            <tab-content title="Address" icon="feather icon-home" class="mb-5">
                <div class="vx-row">
                    <!-- LEFT COL: NEW ADDRESS -->
                    <div class="vx-col lg:w-2/3 w-full">
                        <vx-card title="Add New Address" subtitle="Be sure to check &quot;Deliver to this address&quot; when you have finished" class="mb-base">
                          <form data-vv-scope="add-new-address">
                            <!-- Address -->
                            <div>
                              <label class="text-sm">Address</label>
                              <vs-textarea v-model="addressInput"/> 
                            </div>
                            <vs-button class="mt-6 ml-auto flex" @click="savedAddress()">SAVE AND DELIVER HERE</vs-button>
                          </form>
                        </vx-card>
                    </div>

                    <!-- RIGHT COL: CONTINUE WITH SAVED ADDRESS -->
                    <div class="vx-col lg:w-1/3 w-full">
                        <vx-card :title="user.name">
                            <div>
                                <p>{{ user.address }}</p>
                            </div>
                            <vs-divider />
                            <vs-button class="w-full" @click="$refs.checkoutWizard.nextTab()">DELIVER TO THIS ADDRESS</vs-button>
                        </vx-card>
                    </div>

                </div>
            </tab-content>

            <!-- tab 3 content -->
            <tab-content title="Payment" icon="feather icon-credit-card" class="mb-5">
                <div class="vx-row">
                    <!-- LEFT COL: PAYMENT OPTIONS -->
                    <div class="vx-col lg:w-2/3 w-full">
                        <vx-card title="Payment Options" subtitle="Be sure to click on correct payment option" class="mb-base">
                            <div>
                                <ul>
                                  <li>
                                    <!-- CVV BLOCK -->
                                    <form data-vv-scope="cvv-form">
                                        <div class="flex items-center flex-wrap">
                                            <vs-input
                                                v-validate="'required'"
                                                name="payment"
                                                v-model="payment"
                                                class="w-2/3 mr-3 ml-2" />
                                            <vs-button @click="makePayment">CONTINUE</vs-button>
                                        </div>
                                        <span v-show="errors.has('cvv-form.cvv')" class="text-danger ml-24">{{ errors.first('cvv-form.cvv') }}</span>
                                    </form>
                                  </li>
                                </ul>
                            </div>
                        </vx-card>
                    </div>

                    <!-- RIGHT COL: PRICE -->
                    <div class="vx-col lg:w-1/3 w-full">
                        <vx-card title="Price Details">
                            <div class="flex justify-between mb-2" v-for="item in carts" :key="item.objectID">
                                <span class="text-grey">{{ item.product.productName }}</span>
                                <span>Rp {{ item.product.productPrice * item.quantity }}</span>
                            </div>
                            <vs-divider />

                            <div class="flex justify-between font-semibold mb-3">
                                <span>Total</span>
                                <span>Rp {{ totalCart }}</span>
                            </div>
                        </vx-card>
                    </div>
                </div>
            </tab-content>

        </form-wizard>
    </div>
</template>

<script>
import { FormWizard, TabContent } from 'vue-form-wizard'
import 'vue-form-wizard/dist/vue-form-wizard.min.css'
const ItemListView = () => import('./components/ItemListView.vue')
import apiCart from '../../../api/cart'
import apiProduct from '../../../api/product'
import apiUser from '../../../api/user'
import apiTransaction from '../../../api/transaction'

export default {
  data () {
    return {
      carts: [],
      user: { 
        address: "",
      },
      totalCart : 0,
      payment: ""
    }
  },
  computed: {
    cartItems () {
      return this.$store.state.eCommerce.cartItems.slice().reverse()
    },
    isInWishList () {
      return (itemId) => this.$store.getters['eCommerce/isInWishList'](itemId)
    }
  },
  methods: {
    savedAddress(){
      if(this.addressInput == null || this.addressInput == ""){
        this.$vs.notify({
          title: 'Error',
          text: 'Please enter valid details',
          color: 'warning',
          iconPack: 'feather',
          icon: 'icon-alert-circle'
        })
      }else{
        this.user.address = this.addressInput 
        console.log("hei keubah" , this.user.address)
      }
    },
    calculateTotalPrice () {
      let tempTotal = 0;
      for(let i=0; i<this.carts.length; i++){
        console.log(this.carts[i].product.productPrice , this.carts[i].quantity)
        tempTotal +=  parseInt( this.carts[i].product.productPrice) *  parseInt(this.carts[i].quantity) 
      }
      this.totalCart = tempTotal ;
    },
    getDetailCart(){
      this.totalCart = 0
      for(let i=0; i<this.carts.length; i++){
        //customerId, quantity, productId
        this.carts[i].product = {
          id: 0,
          productName: "Dummy",
          productPrice: 0,
          productRate: 0,
          productStock: 0,
          description: "",
          sellerId: 0,
          categoryId: 0,
          pictures: [],
        }
        this.carts[i].store = {
          userId: 0,
          user: {},
          storeName: "",
          city: ""
        }
        apiProduct
          .GetProductById(this.carts[i].productId)
          .then((response) => {  
            this.carts[i].id = response[0].id
            this.carts[i].product.productName = response[0].productName
            this.carts[i].product.productPrice = response[0].productPrice
            this.carts[i].product.productRate = response[0].productRate
            this.carts[i].product.productStock = response[0].productStock
            this.carts[i].product.description = response[0].description
            this.carts[i].product.sellerId = response[0].sellerId
            this.carts[i].product.categoryId = response[0].categoryId
            this.carts[i].product.pictures = response[0].pictures
          }).then(()=>{
              apiUser
                .GetStoreByIdWithProduct(this.carts[i].product.sellerId)
                .then((response) => {
                  this.carts[i].store.userId = response.userId
                  this.carts[i].store.storeName = response.storeName                    
                  this.carts[i].store.city = response.city
                }).catch((error) => { 
                  console.log('Error get data store!', error)
                })
          }).then(() =>{
            this.calculateTotalPrice()
          })
          .catch((error) => { 
            console.log('Error get product by id!', error) 
          }) 
      }

    },
    makePayment(){
      if(this.payment != null && this.payment != ""){
        apiTransaction
        .AddTransaction(parseInt(localStorage.getItem('userId')), this.user.address, this.payment)
        .then(() => {
          this.$vs.notify({
            title: 'Success',
            text: 'Payment received successfully',
            color: 'success',
            iconPack: 'feather',
            icon: 'icon-check'
          })
        })
        .catch(() => {          
          this.$vs.notify({
            title: 'Error',
            text: 'Please enter payment option',
            color: 'warning',
            iconPack: 'feather',
            icon: 'icon-alert-circle'
          })
        })
      }else{
        this.$vs.notify({
          title: 'Error',
          text: 'Please enter payment option',
          color: 'warning',
          iconPack: 'feather',
          icon: 'icon-alert-circle'
        })
      }
    },
    deleteSpesifikCart(productId){
      apiCart
        .DeleteSpecificCart(parseInt(localStorage.getItem('userId')), productId)
        .then((response) => {
          console.log('spesifik cart: ',response)
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
    // TAB 1
    removeItemFromCart (item) {
      this.$store.dispatch('eCommerce/toggleItemInCart', item)
    },
    wishListButtonClicked (item) {
      // Toggle in Wish List
      if (this.isInWishList(item.objectID)) this.$router.push('/apps/eCommerce/wish-list').catch(() => {})
      else {
        this.toggleItemInWishList(item)
        this.removeItemFromCart(item)
      }
    },
    toggleItemInWishList (item) {
      this.$store.dispatch('eCommerce/toggleItemInWishList', item)
    },
    updateItemQuantity (event, index) {
      const itemIndex = Math.abs(index + 1 - this.cartItems.length)
      this.$store.dispatch('eCommerce/updateItemQuantity', { quantity: event, index: itemIndex })
    },

    // TAB 2
    submitNewAddressForm () {
      return new Promise(() => {
        this.$validator.validateAll('add-new-address').then(result => {
          if (result) {
            // if form have no errors
            this.$refs.checkoutWizard.nextTab()
          } else {
            this.$vs.notify({
              title: 'Error',
              text: 'Please enter valid details',
              color: 'warning',
              iconPack: 'feather',
              icon: 'icon-alert-circle'
            })
          }
        })
      })
    },

    // TAB 3
    makePayment1 () {
      return new Promise(() => {
        this.$validator.validateAll('cvv-form').then(result => {
          if (result) {
            // if form have no errors
            this.$vs.notify({
              title: 'Success',
              text: 'Payment received successfully',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
          } else {
            this.$vs.notify({
              title: 'Error',
              text: 'Please enter valid details',
              color: 'warning',
              iconPack: 'feather',
              icon: 'icon-alert-circle'
            })
          }
        })
      })
    }
  },
  components: {
    ItemListView,
    FormWizard,
    TabContent
  },
  mounted(){
    apiCart
      .GetCart(localStorage.getItem('userId'))
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
          this.carts = response
          for(let i=0; i<response.length; i++){
            this.carts[i].productId = response[i].productId
          }
        }
        console.log('cart checkout: ', this.carts)
        this.getDetailCart()
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
    apiUser
      .GetCustomerWithAssociation(localStorage.getItem('userId'))
      .then((response) => { 
        console.log('user: ', response) 
        this.user.name = response[0].user.name
        this.user.address = response[0].address
      })
      .catch((error) => {          
        console.log('Error get data store!', error)
      })
  }
}
</script>

<style lang="scss" scoped>
#ecommerce-checkout-demo {
    .item-view-primary-action-btn {
        color: #2c2c2c !important;
        background-color: #f6f6f6;
    }

    .item-name {
        display: -webkit-box;
        -webkit-box-orient: vertical;
        overflow: hidden;
        -webkit-line-clamp: 2;
    }

    .vue-form-wizard {
        padding-bottom: 0;

        ::v-deep .wizard-header {
            padding: 0;
        }

        ::v-deep .wizard-tab-content {
            padding-right: 0;
            padding-left: 0;
            padding-bottom: 0;

            .wizard-tab-container{
              margin-bottom: 0 !important;
            }
        }
    }
}
</style>
