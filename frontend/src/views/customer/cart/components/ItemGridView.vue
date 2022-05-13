<!-- =========================================================================================
  File Name: ItemGridView.vue
  Description: Item Component - Grid VIew
  ----------------------------------------------------------------------------------------
  Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
  Author: Pixinvent
  Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->

<template>
    <!-- <div class="item-grid-view vx-row match-height">
        <div class="vx-col" :class="gridColumnWidth" v-for="item in items" :key="item.objectID"> -->
            <vx-card class="grid-view-item mb-base overflow-hidden" v-on="$listeners">
                <template slot="no-body">

                    <!-- ITEM IMAGE  -->
                    <div class="item-img-container bg-white h-64 flex items-center justify-center mb-4 cursor-pointer" @click="navigate_to_detail_view">
                        <img src="https://i.pinimg.com/564x/d0/a7/f0/d0a7f03c63f1c54887d739892fd75f70.jpg" class="grid-view-img px-4">
                    </div>
                    <div class="item-details px-4">

                        <!-- RATING & PRICE  -->
                        <div class="flex justify-between items-center">
                            <p class="truncate text-sm hover:text-primary cursor-pointer" @click="viewStore(product.sellerId)">{{ product.store.storeName }}</p>
                            <h6 class="font-bold">Rp {{ product.productPrice }}</h6>
                        </div>

                        <!-- TITLE & DESCRIPTION  -->
                        <div class="my-4">
                            <h6 class="truncate font-semibold mb-1 hover:text-primary cursor-pointer" @click="viewProduct(product.id)">{{ product.productName }}</h6>
                            <p class="item-description truncate text-sm">{{ product.description }}</p>
                        </div>
                    </div>

                    <!-- SLOT: ACTION BUTTONS  -->
                    <slot name="action-buttons" />
                </template>
            </vx-card> 
         <!-- </div>
    </div> -->
</template>

<script>
import apiProduct from '../../../../api/product'
import apiUser from '../../../../api/user'

export default{
  props: {
    item: {
      type: Object,
      required: true
    }
  },
  data(){
    return {
      product: {
        id: 0,
        productName: "",
        productPrice: 0,
        productRate: 0,
        productStock: 0,
        description: "",
        sellerId: 0,
        categoryId: 0,
        pictures: [],
        store: {
            userId: 0,
            user: {},
            storeName: "",
            city: ""
        },
      }
    }
  },
  methods: {
    navigate_to_detail_view () {
      this.$router.push({name: 'customer-product-view', params: {item_id: this.item.objectID }})
        .catch(() => {})
    },
    getProduct(){
        apiProduct
            .GetProductById(this.item.productId)
            .then((response) => {  
              this.product.id = response[0].id
              this.product.productName = response[0].productName
              this.product.productPrice = response[0].productPrice
              this.product.productRate = response[0].productRate
              this.product.productStock = response[0].productStock
              this.product.description = response[0].description
              this.product.sellerId = response[0].sellerId
              this.product.categoryId = response[0].categoryId
              this.product.pictures = response[0].pictures
            }).then(()=>{
                apiUser
                  .GetStoreByIdWithProduct(this.product.sellerId)
                  .then((response) => {
                    this.product.store.userId = response.userId
                    this.product.store.storeName = response.storeName                    
                    this.product.store.city = response.city
                  }).catch((error) => { 
                    console.log('Error get data store!', error)
                  })
            })
            .catch((error) => { 
            console.log('Error get product by id!', error) 
            }) 
    },
    viewStore (productId) {
      this.$router.push({ path: `/customer/store/${productId}` }).catch(() => {})
    },
    viewProduct (productId) {
      this.$router.push({ path: `/customer/product/${productId}` }).catch(() => {})
    },
  },
  mounted(){
    this.getProduct()
  }
}
</script>

<style lang="scss" scoped>
.grid-view-item {
    .grid-view-img {
        max-width: 100%;
        max-height: 100%;
        width: auto;
        transition: .35s;
    }

    &:hover {
        transform: translateY(-5px);
        box-shadow: 0px 4px 25px 0px rgba(0,0,0,.25);

        .grid-view-img{
            opacity: 0.9;
        }
    }
}
</style>
