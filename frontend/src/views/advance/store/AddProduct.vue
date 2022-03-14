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
      <vx-card title="Add New Product" subtitle="Be sure to check &quot;Product Details&quot; when you have finished" class="mb-base">
          <form data-vv-scope="add-new-address">
              <div class="vx-row">
                  <div class="vx-col sm:w-1/2 w-full">
                      <vs-input
                          v-validate="'required|alpha_spaces'"
                          data-vv-as="field"
                          name="Name"
                          label="Name:"
                          v-model="Name"
                          class="w-full mt-0" />
                      <span v-show="errors.has('add-new-product.Name')" class="text-danger">{{ errors.first('add-new-product.Name') }}</span>
                  </div>
                  <div class="vx-col sm:w-1/2 w-full">
                      <vs-input
                          v-validate="'required|alpha_spaces'"
                          name="Description"
                          label="Description:"
                          v-model="Description"
                          class="w-full mt-0" />
                      <span v-show="errors.has('add-new-product.description')" class="text-danger">{{ errors.first('add-new-product.description') }}</span>
                  </div>
              </div>
              <div class="vx-row">
                  <div class="vx-col sm:w-1/2 w-full">
                      <vs-input
                          v-validate="'required'"
                          name="Price"
                          label="Price:"
                          v-model="Price"
                          class="w-full mt-5" />
                      <span v-show="errors.has('add-new-product.price')" class="text-danger">{{ errors.first('add-new-product.price') }}</span>
                  </div>
                  <div class="vx-col sm:w-1/2 w-full">
                      <vs-input
                          name="Shipping"
                          label="Shipping:"
                          v-model="Shipping"
                          class="w-full mt-5" />
                  </div>
              </div>
              <div class="vx-row">
                  <div class="vx-col sm:w-1/2 w-full">
                      <vs-input
                          v-validate="'required'"
                          name="Weight"
                          label="Weight:"
                          v-model="Weight"
                          class="w-full mt-5" />
                      <span v-show="errors.has('add-new-product.weight')" class="text-danger">{{ errors.first('add-new-product.weight') }}</span>
                  </div>
                  <div class="vx-col sm:w-1/2 w-full">
                      <vs-input
                          v-validate="'required'"
                          name="Stock"
                          label="Stock:"
                          v-model="Stock"
                          class="w-full mt-5" />
                      <span v-show="errors.has('add-new-product.stock')" class="text-danger">{{ errors.first('add-new-product.stock') }}</span>
                  </div>
              </div>
              <vs-button class="mt-6 ml-auto flex" @click.prevent="submitNewAddressForm">SAVE</vs-button>
          </form>
      </vx-card>
    </div>
</template>

<script>
import { FormWizard, TabContent } from 'vue-form-wizard'
import 'vue-form-wizard/dist/vue-form-wizard.min.css'
const ItemListView = () => import('./components/ItemListView.vue')

export default {
  data () {
    return {

      // TAB 2
      fullName: '',
      mobileNum: '',
      pincode: '',
      houseNum: '',
      area: '',
      landmark: '',
      city: '',
      state: '',
      addressType: 1,
      addressTypeOptions: [
        { text: 'Home', value: 1 },
        { text: 'Office', value: 2 }
      ],

      // TAB 3
      paymentMethod: 'debit-card',
      cvv: ''
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
    makePayment () {
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
