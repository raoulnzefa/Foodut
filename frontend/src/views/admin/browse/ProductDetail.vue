<!-- =========================================================================================
  File Name: ECommerceItemDetail.vue
  Description: eCommerce Item Detail page
  ----------------------------------------------------------------------------------------
  Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
  Author: Pixinvent
  Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->

<template>
  <div id="item-detail-page">
    <vs-alert color="danger" title="Error Fetching Product Data" :active.sync="error_occured">
      <span>{{ error_msg }}. </span>
      <span>
        <span>Check </span><router-link :to="{name:'ecommerce-shop'}" class="text-inherit underline">All Items</router-link>
      </span>
    </vs-alert>

    <vx-card>
      <template slot="no-body">
        <div class="item-content">
          <!-- Item Main Info -->
          <div class="product-details p-6">
            <div class="vx-row mt-6">
              <div class="vx-col md:w-2/5 w-full flex items-center justify-center">
                <div class="product-img-container w-3/5 mx-auto mb-10 md:mb-0">
                  <img src="https://i.pinimg.com/564x/d0/a7/f0/d0a7f03c63f1c54887d739892fd75f70.jpg" class="responsive">
                </div>
              </div>

              <!-- Item Content -->
              <div class="vx-col md:w-3/5 w-full">
                <h3>{{ product.productName }}</h3>
                <p class="my-2">
                  <span class="mr-2">by</span>
                  <span>{{ product.sellerId }}</span>
                </p>
                <p class="flex items-center flex-wrap">
                  <span class="text-2xl leading-none font-medium text-primary mr-4 mt-2">Rp {{ product.productPrice }}</span>
                  <span class="pl-4 mr-2 mt-2 border border-solid d-theme-border-grey-light border-t-0 border-b-0 border-r-0"><star-rating :show-rating="false" :star-size="20" read-only /></span>
                  <span class="cursor-pointer leading-none mt-2">{{ product.productRate }} ratings</span>
                </p>
                <vs-divider />
                <p>{{ product.description }}</p>
                <vs-divider />
                <!-- Quantity -->
                <div class="vx-row">
                  <div class="vx-col w-full">
                    <p class="my-2">
                      <span class="text-success">Available</span>
                      <span class="mx-2">-</span>
                      <span>{{ product.productStock }} pcs</span>
                    </p>
                  </div>
                  <div class="vx-col w-full mt-5">
                    <div class="flex flex-wrap items-start mb-4">
                      <!-- Add To Cart Button -->
                      <vs-button
                        class="mr-4 mb-4"
                        icon-pack="feather"
                        icon="icon-shopping-cart"
                        @click="popupAdd=true">
                        ADD TO CART
                      </vs-button>
                      <vs-popup class="popup" background-color="rgba(25,25,25,0.25)" title="Failed to Add Product!" :active.sync="popupAdd">
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
                      <!-- View In Cart Button -->
                      <vs-button
                        key="border"
                        class="mb-4"
                        type="border"
                        icon-pack="feather"
                        icon="icon-shopping-cart"
                        @click="popupView=true">
                        VIEW IN CART
                      </vs-button>
                      <vs-popup class="popup" background-color="rgba(25,25,25,0.25)"  title="Failed to View Cart!" :active.sync="popupView">
                        <p>Please login as customer to view cart. If you don't have customer account, please register first!</p>
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
                </div>
              </div>
            </div>
          </div>
        </div>
      </template>
    </vx-card>
    <vx-card>
      <template slot="no-body">
        <!-- Related Products -->
        <div class="related-products text-center">
          <div class="related-headings mb-8 text-center">
            <h2 class="uppercase mt-10">Related Products</h2>
            <p>People also search for this items</p>
          </div>
          <swiper :options="swiperOption" :dir="$vs.rtl ? 'rtl' : 'ltr'" :key="$vs.rtl" class="related-product-swiper px-12 py-6">
            <swiper-slide v-for='product in products' :key='product.id' class="p-6 rounded cursor-pointer">
              <div @click="viewProduct(product.id)">
                <!-- Item Heading -->
                <div class="item-heading mb-4">
                  <p class="text-lg font-semibold truncate">{{ product.productName }}</p>
                  <p class="text-sm">
                    <span class="mr-2">by</span>
                    <span>{{ product.sellerId }}</span>
                  </p>
                </div>
                <!-- Item Image -->
                <div class="img-container w-full mx-auto my-base">
                  <img class="responsive" src="https://i.pinimg.com/564x/d0/a7/f0/d0a7f03c63f1c54887d739892fd75f70.jpg">
                </div>
                <!-- Item Meta -->
                <div class="item-meta">
                  <star-rating :show-rating="false" :rating="product.rating" :star-size="14" class="justify-center" read-only />
                  <p class="text-lg font-medium text-primary">${{ product.productPrice }}</p>
                </div>
              </div>
            </swiper-slide>
            <div class="swiper-button-prev" slot="button-prev"></div>
            <div class="swiper-button-next" slot="button-next"></div>
          </swiper>
        </div>

      </template>
    </vx-card>
  </div>
</template>

<script>
import 'swiper/dist/css/swiper.min.css'
import { swiper, swiperSlide } from 'vue-awesome-swiper'
import StarRating from 'vue-star-rating'
import apiProduct from '../../../api/product'
import apiUser from '../../../api/user'

export default{
  components: {
    swiper,
    swiperSlide,
    StarRating
  },
  data () {
    return {
      product: {},
      products: [],
      popupAdd: false,
      popupView: false,
      error_occured: false,
      error_msg: '',

      // Related Products Swiper
      swiperOption: {
        slidesPerView: 5,
        spaceBetween: 55,
        breakpoints: {
          1600: {
            slidesPerView: 4,
            spaceBetween: 55
          },
          1300: {
            slidesPerView: 3,
            spaceBetween: 55
          },
          900: {
            slidesPerView: 2,
            spaceBetween: 55
          },
          640: {
            slidesPerView: 1,
            spaceBetween: 55
          }
        },
        navigation: {
          nextEl: '.swiper-button-next',
          prevEl: '.swiper-button-prev'
        }
      },

      // Below is data which is common in any item
      // Algolia's dataSet don't provide this kind of data. So, here's dummy data for demo
      available_item_colors: ['#7367F0', '#28C76F', '#EA5455', '#FF9F43', '#1E1E1E'],
      opt_color: '#7367F0',
      is_hearted: false,

      related_items: [
        {
          'name'       : 'Apple - Apple Watch Series 1 42mm Space Gray Aluminum Case Black Sport Band - Space Gray Aluminum',
          'brand'      : 'Apple',
          'price'      : 229,
          'image'      : 'https://pixinvent.com/demo/vuexy-vuejs-admin-dashboard-template/products/01.png',
          'rating'     : 4,
          'objectID'   : '5546604'
        },
        {
          'name'       : 'Beats by Dr. Dre - Powerbeats2 Wireless Earbud Headphones - Black/Red',
          'brand'      : 'Beats by Dr. Dre',
          'price'      : 199.99,
          'image'      : 'https://pixinvent.com/demo/vuexy-vuejs-admin-dashboard-template/products/08.png',
          'rating'     : 4,
          'objectID'   : '5565002'
        },
        {
          'name'       : 'Amazon - Fire TV Stick with Alexa Voice Remote - Black',
          'brand'      : 'Amazon',
          'price'      : 39.99,
          'image'      : 'https://pixinvent.com/demo/vuexy-vuejs-admin-dashboard-template/products/03.png',
          'rating'     : 4,
          'objectID'   : '5477500'
        },
        {
          'name'       : 'Apple - Apple Watch Nike+ 42mm Silver Aluminum Case Silver/Volt Nike Sport Band - Silver Aluminum',
          'brand'      : 'Apple',
          'price'      : 399,
          'image'      : 'https://pixinvent.com/demo/vuexy-vuejs-admin-dashboard-template/products/07.png',
          'rating'     : 4,
          'objectID'   : '5547700'
        },
        {
          'name'       : 'Google - Chromecast Ultra - Black',
          'brand'      : 'Google',
          'price'      : 69.99,
          'image'      : 'https://pixinvent.com/demo/vuexy-vuejs-admin-dashboard-template/products/05.png',
          'rating'     : 4,
          'objectID'   : '5578628'
        },
        {
          'name'       : 'Beats by Dr. Dre - Beats EP Headphones - White',
          'brand'      : 'Beats by Dr. Dre',
          'price'      : 129.99,
          'image'      : 'https://pixinvent.com/demo/vuexy-vuejs-admin-dashboard-template/products/02.png',
          'rating'     : 4,
          'objectID'   : '5577781'
        },
        {
          'name'       : 'LG - 40" Class (39.5" Diag.) - LED - 1080p - HDTV - Black',
          'brand'      : 'LG',
          'price'      : 279.99,
          'image'      : 'https://pixinvent.com/demo/vuexy-vuejs-admin-dashboard-template/products/09.png',
          'rating'     : 4,
          'objectID'   : '5613404'
        }
      ]
    }
  },
  computed: {
    item_qty () {
      // const item = this.$store.getters['eCommerce/getCartItem'](this.item_data.objectID)
      // return Object.keys(item).length ? item.quantity : 1
      return 1
    },
    itemColor () {
      return (obj) => {
        const style_obj = {}

        obj.style.forEach(p => { style_obj[p] = obj.color })

        return style_obj
      }
    },
    isInWishList () {
      return 0 // (itemId) => this.$store.getters['eCommerce/isInWishList'](itemId)
    },
    isInCart () {
      return 0 // (itemId) => this.$store.getters['eCommerce/isInCart'](itemId)
    }
  },
  methods: {
    // toggleItemInWishList (item) {
    //   this.$store.dispatch('eCommerce/toggleItemInWishList', item)
    // },
    // toggleItemInCart (item) {
    //   this.$store.dispatch('eCommerce/toggleItemInCart', item)
    // }
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
    viewProduct (productId) {
      this.$router.push({ path: `/admin/product/${productId}` }).catch(() => {})
    }
  },
  mounted () {
    apiProduct
      .GetProductById(this.$route.params.product_id)
      .then((response) => { 
        this.product = response[0] 
        apiUser
          .GetStoreByIdWithProduct(this.product.sellerId)
          .then((response) => { 
            this.product.sellerId = response.storeName 
            this.products = response.listProduct
            for(let i=0; i<this.products.length; i++){
              this.products[i].sellerId = response.storeName
            }
          })
          .catch((error) => { console.log('Error get storename!', error)})
        console.log(this.product)
      })
      .catch((error) => { console.log('Error get id product!', error)})
  }
}
</script>

<style lang="scss">

@import "@/assets/scss/vuexy/_variables.scss";

#item-detail-page {
  .color-radio {
    height: 2.28rem;
    width: 2.28rem;

    > div {
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
    }
  }

  .item-features {
    background-color: #f7f7f7;

    .theme-dark & {
      background-color: $theme-dark-secondary-bg;
    }
  }

  .product-sales-meta-list {
    .vs-list--icon {
      padding-left: 0;
    }
  }

  .related-product-swiper {
      // padding-right: 2rem;
      // padding-left: 2rem;

    .swiper-wrapper {
      padding-bottom: 2rem;

      > .swiper-slide {
        background-color: #f7f7f7;
        box-shadow: 0 4px 18px 0 rgba(0,0,0,0.1), 0 5px 12px 0 rgba(0,0,0,0.08) !important;

        .theme-dark & {
          background-color: $theme-light-dark-bg;
        }
      }
    }

    .swiper-button-next,
    .swiper-button-prev {
      transform: scale(.5);
      filter: hue-rotate(40deg);
    }

    .swiper-button-next {
      right: 0
    }

    .swiper-button-prev {
      left: 0;
    }
  }
}
</style>
