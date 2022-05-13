<template>
  <div id="data-list-thumb-view" class="data-list-container mt-6">

    <data-view-sidebar :isSidebarActive="addNewDataSidebar" @closeSidebar="toggleDataSidebar" :data="sidebarData" />

    <vs-table ref="table" multiple v-model="selected" pagination :max-items="itemsPerPage" search :data="products">

      <div slot="header" class="flex flex-wrap-reverse items-center flex-grow justify-between">

        <div class="flex flex-wrap-reverse items-center">

          <!-- ACTION - DROPDOWN -->
          <vs-dropdown vs-trigger-click class="cursor-pointer mr-4 mb-4">

            <div class="p-4 shadow-drop rounded-lg d-theme-dark-bg cursor-pointer flex items-center justify-center text-lg font-medium w-32">
              <span class="mr-2">Actions</span>
              <feather-icon icon="ChevronDownIcon" svgClasses="h-4 w-4" />
            </div>

            <vs-dropdown-menu>

              <vs-dropdown-item>
                <span class="flex items-center">
                  <feather-icon icon="TrashIcon" svgClasses="h-4 w-4" class="mr-2" />
                  <span>Delete</span>
                </span>
              </vs-dropdown-item>
            </vs-dropdown-menu>
          </vs-dropdown>

          <!-- ADD NEW -->
          <div class="p-3 mb-4 mr-4 rounded-lg cursor-pointer flex items-center justify-between text-lg font-medium text-base text-primary border border-solid border-primary" @click="addNewData">
              <feather-icon icon="PlusIcon" svgClasses="h-4 w-4" />
              <span class="ml-2 text-base text-primary">Add New</span>
          </div>
        </div>
      </div>

      <template slot="thead">
        <vs-th>Image</vs-th>
        <vs-th sort-key="name">Name</vs-th>
        <vs-th sort-key="category">Category</vs-th>
        <vs-th sort-key="rate">Rate</vs-th>
        <vs-th sort-key="stock">Stock</vs-th>
        <vs-th sort-key="price">Price</vs-th>
        <vs-th>Action</vs-th>
      </template>

      <template slot-scope="{data}">
        <tbody>
          <vs-tr :data="tr" :key="indextr" v-for="(tr, indextr) in data">

            <vs-td class="img-container">
              <img src="https://i.pinimg.com/564x/d0/a7/f0/d0a7f03c63f1c54887d739892fd75f70.jpg" class="product-img" />
            </vs-td>

            <vs-td>
              <p class="product-name font-medium truncate">{{ tr.productName }}</p>
            </vs-td>

            <vs-td>
              <p class="product-category">{{ tr.categoryId }}</p>
            </vs-td>

            <vs-td>
              <p class="product-rate">
                <img :src="require('../../../../assets/images/raty/star-on-2.png')"/> 
                {{ tr.productRate }}
              </p>
            </vs-td>

            <vs-td>
              <p class="product-stock">{{ tr.productStock }}</p>
            </vs-td>

            <vs-td>
              <p class="product-price">${{ tr.productPrice }}</p>
            </vs-td>

            <vs-td class="whitespace-no-wrap">
              <feather-icon icon="EditIcon" svgClasses="w-5 h-5 hover:text-primary stroke-current" @click.stop="editData(tr.id)" />
              <feather-icon icon="TrashIcon" svgClasses="w-5 h-5 hover:text-danger stroke-current" class="ml-2" @click.stop="deleteData(tr.id)" />
            </vs-td>
          </vs-tr>
        </tbody>
      </template>
    </vs-table>
  </div>
</template>

<script>
import DataViewSidebar from './DataViewSidebar.vue'
import moduleDataList from '@/store/data-list/moduleDataList.js'
import apiProduct from '../../../../api/product'
// import apiUser from '../../../../api/user'

export default {
  components: {
    DataViewSidebar
  },
  data () {
    return {
      selected: [],
      products: [], // Data Product
      itemsPerPage: 5,
      isMounted: false,
      addNewDataSidebar: false,
      sidebarData: {}
    }
  },
  computed: {
    productData() {
      return this.products
    },
    currentPage () {
      if (this.isMounted) {
        return this.$refs.table.currentx
      }
      return 0
    },
    queriedItems () {
      return this.$refs.table ? this.$refs.table.queriedResults.length : this.products.length
    }
  },
  methods: {
    addNewData () {
      this.$router.push({ name : 'seller-add-product'}).catch(() => {})
    },
    deleteData (id) {
      console.log(id)
      apiProduct
        .DeleteProduct(id)
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
          // setTimeout(function(){
          //   window.location.reload(1);
          // }, 1500);
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
    editData (productId) {
      console.log(productId);
      this.$router.push({ path: `/seller/product/edit/${productId}` }).catch(() => {})
    },
    getOrderStatusColor (status) {
      if (status === 'on_hold')   return 'warning'
      if (status === 'delivered') return 'success'
      if (status === 'canceled')  return 'danger'
      return 'primary'
    },
    getPopularityColor (num) {
      if (num > 90)  return 'success'
      if (num > 70)  return 'primary'
      if (num >= 50) return 'warning'
      if (num < 50)  return 'danger'
      return 'primary'
    },
    toggleDataSidebar (val = false) {
      this.addNewDataSidebar = val
    }
  },
  created () {
    if (!moduleDataList.isRegistered) {
      this.$store.registerModule('dataList', moduleDataList)
      moduleDataList.isRegistered = true
    }
    this.$store.dispatch('dataList/fetchDataListItems')
  },
  mounted () {
    // apiUser
    //   .GetStoreByIdWithProduct(localStorage.getItem('userId'))
    //   .then((response) => { 
    //     console.log('product: ', response)
    //     this.products = response.listProduct
    //   })
    //   .catch((error) => { console.log('Error get all data product!', error) })
  }
}
</script>

<style lang="scss">
#data-list-thumb-view {
  .vs-con-table {

    .product-name {
      max-width: 23rem;
    }

    .vs-table--header {
      display: flex;
      flex-wrap: wrap-reverse;
      margin-left: 1.5rem;
      margin-right: 1.5rem;
      > span {
        display: flex;
        flex-grow: 1;
      }

      .vs-table--search{
        padding-top: 0;

        .vs-table--search-input {
          padding: 0.9rem 2.5rem;
          font-size: 1rem;

          &+i {
            left: 1rem;
          }

          &:focus+i {
            left: 1rem;
          }
        }
      }
    }

    .vs-table {
      border-collapse: separate;
      border-spacing: 0 1.3rem;
      padding: 0 1rem;


      tr{
          box-shadow: 0 4px 20px 0 rgba(0,0,0,.05);
          td{
            padding: 10px;
            &:first-child{
              border-top-left-radius: .5rem;
              border-bottom-left-radius: .5rem;
            }
            &:last-child{
              border-top-right-radius: .5rem;
              border-bottom-right-radius: .5rem;
            }
            &.img-container {
              // width: 1rem;
              // background: #fff;

              span {
                display: flex;
                justify-content: flex-start;
              }

              .product-img {
                height: 110px;
              }
            }
          }
          td.td-check{
            padding: 20px !important;
          }
      }
    }

    .vs-table--thead{
      th {
        padding-top: 0;
        padding-bottom: 0;

        .vs-table-text{
          text-transform: uppercase;
          font-weight: 600;
        }
      }
      th.td-check{
        padding: 0 15px !important;
      }
      tr{
        background: none;
        box-shadow: none;
      }
    }

    .vs-table--pagination {
      justify-content: center;
    }
  }
}
</style>
