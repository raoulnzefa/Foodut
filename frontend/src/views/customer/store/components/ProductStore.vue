<template>
  <div id="data-list-thumb-view" class="data-list-container mt-6">
    <data-view-sidebar :isSidebarActive="addNewDataSidebar" @closeSidebar="toggleDataSidebar" :data="sidebarData" />
    <vs-table ref="table" multiple v-model="table" pagination :max-items="itemsPerPage" search :data="listProduct">
      <template slot="thead">
        <vs-th>Image</vs-th>
        <vs-th sort-key="name">Name</vs-th>
        <vs-th sort-key="category">Category</vs-th>
        <vs-th sort-key="rate">Rate</vs-th>
        <vs-th sort-key="stock">Stock</vs-th>
        <vs-th sort-key="price">Price</vs-th>
      </template>

      <template >
        <tbody>
          <vs-tr :key="index" v-for="(product, index) in listProduct">

            <vs-td class="img-container">
              <img src="https://i.pinimg.com/564x/d0/a7/f0/d0a7f03c63f1c54887d739892fd75f70.jpg" class="product-img" />
            </vs-td>

            <vs-td>
              <p class="product-name font-medium truncate">{{ product.productName }}</p>
            </vs-td>

            <vs-td>
              <p class="product-category">{{ product.categoryId }}</p>
            </vs-td>

            <vs-td>
              <p class="product-rate">
                <img :src="require('../../../../assets/images/raty/star-on-2.png')"/> 
                {{ product.productRate }}
              </p>
            </vs-td>

            <vs-td>
              <p class="product-stock">{{ product.productStock }}</p>
            </vs-td>

            <vs-td>
              <p class="product-price">${{ product.productPrice }}</p>
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

export default {
  props: [
    "listProduct"
  ],
  components: {
    DataViewSidebar
  },
  data () {
    return {
      table: [],
      itemsPerPage: 5,
      isMounted: false,
      addNewDataSidebar: false,
      sidebarData: {}
    }
  },
  methods: {
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
