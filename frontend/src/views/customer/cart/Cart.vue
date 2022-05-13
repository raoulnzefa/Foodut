<!-- =========================================================================================
  File Name: ECommerceWishList.vue
  Description: eCommerce Wish List Page
  ----------------------------------------------------------------------------------------
  Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
  Author: Pixinvent
  Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->

<template>
  <div id="ecommerce-wishlist-demo">
    <tab-content title="Cart" icon="feather icon-shopping-cart" class="mb-5">
      <!-- IF ITEMS IN CART -->
      
        <div class="items-grid-view vx-row match-height" v-if="carts.length">
            <div class="vx-col lg:w-1/4 md:w-1/3 sm:w-1/2 w-full" v-for="item in carts" :key="item.objectID">
                <item-grid-view :item="item">
                    <!-- SLOT: ACTION BUTTONS -->
                    <template slot="action-buttons">
                        <div class="flex flex-wrap">
                            <!-- SECONDARY BUTTON: MOVE TO CART -->
                            <div
                                class="item-view-secondary-action-btn bg-primary p-3 flex flex-grow items-center justify-center text-white cursor-pointer"
                                @click="deleteSpesifikCart(item.productId)">
                                <feather-icon icon="XIcon" svgClasses="h-4 w-4" />
                                <span class="text-sm font-semibold ml-2">REMOVE</span>
                            </div>
                        </div>
                    </template>
                </item-grid-view>
            </div>
        </div>
      
      <!-- IF NO ITEMS IN CART -->
      <vx-card title="You don't have any items in your wish list." v-else>
          <vs-button @click="$router.push('/customer/browse').catch(() => {})">Browse Shop</vs-button>
      </vx-card>
    </tab-content>
  </div>
</template>

<script>
const ItemGridView = () => import('./components/ItemGridView.vue')
import apiCart from '../../../api/cart'

export default {
  data () {
    return {
      carts: []
    }
  },
  components: {
    ItemGridView
  },
  computed: {
    isInCart () {
      return (itemId) => this.$store.getters['eCommerce/isInCart'](itemId)
    },
    isInWishList () {
      return (itemId) => this.$store.getters['eCommerce/isInWishList'](itemId)
    }
  },
  methods: {
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
    }
    // removeItemFromWishList (item) {
    //   this.$store.dispatch('eCommerce/toggleItemInWishList', item)
    // },
    // cartButtonClicked (item) {
    //   if (this.isInCart(item.objectID)) this.$router.push('/apps/eCommerce/checkout').catch(() => {})
    //   else {
    //     this.additemInCart(item)
    //     this.removeItemFromWishList(item)
    //   }
    // },
    // additemInCart (item) {
    //   this.$store.dispatch('eCommerce/additemInCart', item)
    // },
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
}
</script>

<style lang="scss" scoped>
#ecommerce-wishlist-demo {
    .item-view-primary-action-btn {
        color: #2c2c2c !important;
        background-color: #f6f6f6;
        min-width: 50%;
    }

    .item-view-secondary-action-btn {
        min-width: 50%;
    }
}
</style>
