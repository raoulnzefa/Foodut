<!-- =========================================================================================
  File Name: ECommerceShop.vue
  Description: eCommerce Shop Page
  ----------------------------------------------------------------------------------------
  Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
    Author: Pixinvent
  Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->

<template>
  <div>
      <div index-name="instant_search" id="instant-search-demo">
          <div :hits-per-page.camel="9" />
          <div id="content-container" class="relative clearfix">
            <div>
              <vs-sidebar
                class="items-no-padding vs-sidebar-rounded background-absolute"
                parent="#content-container"
                :click-not-close="clickNotClose"
                :hidden-background="clickNotClose"
                v-model="isFilterSidebarActive">

                <div class="p-6 filter-container">
                  <!-- PRICE -->
                  <h6 class="font-bold mb-3">Price Range</h6>
                  <div>
                    <ul class="demo-aligment">
                      <li><vs-radio vs-value="low">{{ numericItems[0].label }}</vs-radio></li>
                      <li><vs-radio vs-value="lowMedium">{{ numericItems[1].label }}</vs-radio></li>
                      <li><vs-radio vs-value="mediumHigh">{{ numericItems[2].label }}</vs-radio></li>
                      <li><vs-radio vs-value="high">{{ numericItems[3].label }}</vs-radio></li>
                    </ul>
                  </div>
                  <vs-divider />

                  <!-- CATEGORIES -->
                  <h6 class="font-bold mb-4">Category</h6>
                  <div>
                    <ul>
                      <li v-for='category in categories' :key='category.id'><vs-radio class="mb-1 ml-0" vs-value="category">{{ category.productCategory }}</vs-radio></li>
                    </ul>
                  </div>
                  <vs-divider />

                  <!-- Store -->
                  <h6 class="font-bold mb-4">Store</h6>
                    <vs-checkbox class="mb-1 ml-0" v-for='store in stores' :key='store.id'>{{ store.storeName }}</vs-checkbox>
                  <vs-divider />

                  <vs-button class="w-full" color="danger" type="border">Remove Filters</vs-button>
                </div>
              </vs-sidebar>
            </div>
            <!-- RIGHT COL -->
            <div :class="{'sidebar-spacer-with-margin': clickNotClose}">
              <div class="mb-8">
                <vs-input icon-no-border label-placeholder="Search" class="w-full input-rounded-full" icon="icon-search" icon-pack="feather" />
              </div>
              <div class="vx-row">
                <div class="vx-col w-full sm:w-1/3 md:w-1/3 mb-base" v-for='product in products' :key='product.id'>
                  <vx-card>
                    <img src='https://i.pinimg.com/564x/d0/a7/f0/d0a7f03c63f1c54887d739892fd75f70.jpg' alt="product-img" class="w-full mb-3">
                    <div class="vx-row">
                      <div class="vx-col w-1/2">
                        <p class="text-sm">{{ product.sellerId }}</p>
                      </div>
                      <div class="vx-col w-1/2">
                        <p style="text-align:right">Rp {{ product.productPrice }}</p>
                      </div>
                    </div>
                    <p class="truncate font-semibold mb-1 hover:text-primary cursor-pointer" @click="viewProduct(product.id)">{{ product.productName }}</p>
                    <p class="item-description truncate text-sm">{{ product.description }}</p>
                    <div class="vx-row  mt-3">
                      <div class="vx-col w-full">
                        <vs-button class="w-full"  @click="popupAdd=true" color="primary" type="filled" icon-pack="feather" icon="icon-shopping-cart">Add</vs-button>
                        <vs-popup class="popup" background-color="rgba(25,25,25,0.1)" title="Failed to Add Product!" :active.sync="popupAdd">
                          <p>Please login as customer to add product in your cart. If you don't have customer account, please register first!</p>
                          <vs-divider />
                          <div class="vx-row mt-3">
                            <div class="vx-col w-1/2">
                              <vs-button class="w-full"  @click="login" color="primary" type="filled">Login</vs-button>
                            </div>
                            <div class="vx-col w-1/2">
                              <vs-button class="w-full"  @click="register" color="primary" type="border">Register</vs-button>
                            </div>
                          </div>
                        </vs-popup>
                      </div>
                    </div>
                  </vx-card>
                </div>
              </div>
            </div>
          </div>
      </div>
  </div>
</template>

<script>
import apiUser from '../../../api/user'
import apiCategory from '../../../api/category'
import apiProduct from '../../../api/product'

export default {
  data () {
    return {
      products: [],
      // Filter Sidebar
      isFilterSidebarActive: true,
      clickNotClose: true,
      currentItemView: 'item-grid-view',
      numericItems: [
        { label: '<= Rp 10000', end: 10000 },
        { label: 'Rp 10000 - 50000', start: 10000, end: 50000 },
        { label: 'Rp 50000 - 100000', start: 50000, end: 100000 },
        { label: '>= Rp 100000', start: 100000 }
      ],
      categories: [],
      stores: [],
      popupAdd: false,
      popupView: false
    }
  },
  computed: {
    toValue () {
      return (value, range) => [
        value.min !== null ? value.min : range.min,
        value.max !== null ? value.max : range.max
      ]
    },

    // GRID VIEW
    isInCart () {
      return (itemId) => this.$store.getters['eCommerce/isInCart'](itemId)
    },
    isInWishList () {
      return (itemId) => this.$store.getters['eCommerce/isInWishList'](itemId)
    },
    windowWidth () { return this.$store.state.windowWidth }
  },
  watch: {
    windowWidth () {
      this.setSidebarWidth()
    }
  },
  methods: {
    setSidebarWidth () {
      if (this.windowWidth < 992) {
        this.isFilterSidebarActive = this.clickNotClose = false
      } else {
        this.isFilterSidebarActive = this.clickNotClose = true
      }
    },
    viewProduct (productId) {
      this.$router.push({ path: `/admin/product/${productId}` }).catch(() => {})
    },
    login () {
      this.popupAdd = false
      this.popupView = false
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
            return
          }
          this.$vs.loading.close()
          this.$router.push({ name : 'login-page'}).catch(() => {})
        })
        .catch((error) => {
          this.$vs.loading.close()
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },
    register () {
      this.popupAdd = false
      this.popupView = false
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
            return
          }
          this.$vs.loading.close()
          this.$router.push({ name : 'register-page'}).catch(() => {})
        })
        .catch((error) => {
          this.$vs.loading.close()
          this.$vs.notify({
            title: 'Error',
            text: error.message,
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          })
        })
    },

    // GRID VIEW - ACTIONS
    toggleFilterSidebar () {
      if (this.clickNotClose) return
      this.isFilterSidebarActive = !this.isFilterSidebarActive
    },
    toggleItemInWishList (item) {
      this.$store.dispatch('eCommerce/toggleItemInWishList', item)
    },
    additemInCart (item) {
      this.$store.dispatch('eCommerce/additemInCart', item)
    },
    cartButtonClicked (item) {
      this.isInCart(item.objectID) ? this.$router.push('/customer/checkout').catch(() => {}) : this.additemInCart(item)
    }
  },
  mounted () {
    apiProduct
      .GetAllProduct()
      .then((response) => { 
        this.products = response 
        for(let i=0; i<this.products.length; i++){
          apiUser
            .GetStoreByIdWithProduct(this.products[i].sellerId)
            .then((response) => { this.products[i].sellerId = response.storeName })
            .catch((error) => { console.log('Error get storename!', error)})
        }  
      })
      .catch((error) => { console.log('Error get all product!', error)})
    apiCategory
      .GetAllCategories()
      .then((response) => { this.categories = response })
      .catch((error) => { console.log('Error get all categories!', error)})
    apiUser
      .GetAllStore()
      .then((response) => { this.stores = response })
      .catch((error) => { console.log('Error get all store!', error)})
  },
  created () {
    this.setSidebarWidth()
  }
}

</script>


<style lang="scss">
#instant-search-demo {
  .header {
    .algolia-filters-label {
      width: calc(260px + 2.4rem);
    }
  }

  #content-container {

    .vs-sidebar {
      position: relative;
      float: left;
    }
  }

  .search-input-right-aligned-icon {
    padding: 1rem 1.5rem;
  }

  .price-slider {
    min-width: unset;
  }

  .item-view-primary-action-btn {
    color: #2c2c2c !important;
    background-color: #f6f6f6;
    min-width: 50%;
  }

  .item-view-secondary-action-btn {
    min-width: 50%;
  }
}

.theme-dark {
  #instant-search-demo {
    #content-container {
      .vs-sidebar {
        background-color: #10163a;
      }
    }
  }
}

@media (min-width: 992px) {
  .vs-sidebar-rounded {
    .vs-sidebar {
      border-radius: .5rem;
    }

    .vs-sidebar--items {
      border-radius: .5rem;
    }
  }
}

@media (max-width: 992px) {
  #content-container {
    .vs-sidebar {
      position: absolute !important;
      float: none !important;
    }
  }
}

</style>

