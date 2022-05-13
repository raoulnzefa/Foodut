<!-- =========================================================================================
  File Name: ListStore.vue
  Description: List Store page
  ----------------------------------------------------------------------------------------
  Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
  Author: Pixinvent
  Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->

<template>
  <div id="page-user-store">
    <div class="vx-card p-3">
      <!-- AgGrid Table -->
      <!--https://www.ag-grid.com/vue-data-grid/row-selection/-->
      <ag-grid-vue
        ref="agGridTable"
        :components="components"
        :gridOptions="gridOptions"
        class="ag-theme-material w-100 my-4 ag-grid-table"
        :columnDefs="columnDefs"
        :defaultColDef="defaultColDef"
        :rowData="usersData"
        rowSelection="multiple"
        colResizeDefault="shift"
        :animateRows="true"
        :floatingFilter="true"
        :pagination="true"
        :paginationPageSize="paginationPageSize"
        :suppressPaginationPanel="true"
        @row-clicked="onRowClicked"
        :enableRtl="$vs.rtl">
      </ag-grid-vue>

      <vs-pagination
        :total="totalPages"
        :max="7"
        v-model="currentPage" />
    </div>
  </div>
</template>

<script>
import { AgGridVue } from 'ag-grid-vue'
import '@/assets/scss/vuexy/extraComponents/agGridStyleOverride.scss'
import vSelect from 'vue-select'

// Store Module
import moduleUserManagement from '@/store/user-management/moduleUserManagement.js'

// Cell Renderer
import CellRendererVerified from './components/CellRendererVerified.vue'
import CellRendererActions from './components/CellRendererActions.vue'

//Api User
import apiUser from '../../../api/user'


export default {
  components: {
    AgGridVue,
    vSelect,

    // Cell Renderer
    CellRendererVerified,
    CellRendererActions
  },
  data () {
    return {
      // AgGrid
      gridApi: null,
      gridOptions: {},
      defaultColDef: {
        sortable: true,
        resizable: true,
        suppressMenu: true
      },
      columnDefs: [
        {
          headerName: 'ID',
          field: 'id',
          width: 125,
          filter: true
        },
        {
          headerName: 'Store',
          field: 'storeName',
          filter: true,
          width: 230
        },
        {
          headerName: 'Username',
          field: 'username',
          filter: true,
          width: 230
        },
        {
          headerName: 'Email',
          field: 'email',
          filter: true,
          width: 250
        },
        {
          headerName: 'Name',
          field: 'name',
          filter: true,
          width: 200
        },
        {
          headerName: 'City',
          field: 'city',
          filter: true,
          width: 170
        }
      ],
      onRowClicked: event => { 
        console.log('A row store was clicked, Here is detail information about what row was clicked => ' , event)
        console.log(event.data.id)
        this.$router.push(`/admin/store/${event.data.id}`).catch(() => {})
      },
      //Data Seller
      rowStores: [],
      rowSelection: null,
      // Cell Renderer Components
      components: {
        CellRendererVerified,
        CellRendererActions
      }
    }
  },
  watch: {
    roleFilter (obj) {
      this.setColumnFilter('role', obj.value)
    },
    statusFilter (obj) {
      this.setColumnFilter('status', obj.value)
    }
  },
  computed: {
    usersData () {
      return this.rowStores
    },
    paginationPageSize () {
      if (this.gridApi) return this.gridApi.paginationGetPageSize()
      else return 10
    },
    totalPages () {
      if (this.gridApi) return this.gridApi.paginationGetTotalPages()
      else return 0
    },
    currentPage: {
      get () {
        if (this.gridApi) return this.gridApi.paginationGetCurrentPage() + 1
        else return 1
      },
      set (val) {
        this.gridApi.paginationGoToPage(val - 1)
      }
    }
  },
  methods: {
    onSelectionChanged() {
      const selectedRows = this.gridApi.getSelectedRows();
      console.log(selectedRows)
    },
    setColumnFilter (column, val) {
      const filter = this.gridApi.getFilterInstance(column)
      let modelObj = null

      if (val !== 'all') {
        modelObj = { type: 'equals', filter: val }
      }

      filter.setModel(modelObj)
      this.gridApi.onFilterChanged()
    },
    resetColFilters () {
      // Reset Grid Filter
      this.gridApi.setFilterModel(null)
      this.gridApi.onFilterChanged()

      // Reset Filter Options
      this.roleFilter = this.statusFilter = this.isVerifiedFilter = this.departmentFilter = { label: 'All', value: 'all' }

      this.$refs.filterCard.removeRefreshAnimation()
    },
    updateSearchQuery (val) {
      this.gridApi.setQuickFilter(val)
    },
    viewStore (params) {
      console.log("Triggere view store")
      console.log(params)
    }
  },
  mounted () {
    apiUser
      .GetAllStore()
      .then((response) => { this.rowStores = response })
      .catch((error) => { console.log('Error get all data store!', error)})
  },
  created () {
    if (!moduleUserManagement.isRegistered) {
      this.$store.registerModule('userManagement', moduleUserManagement)
      moduleUserManagement.isRegistered = true
    }
    this.$store.dispatch('userManagement/fetchUsers').catch(err => { console.error(err) })
  }
}

</script>

<style lang="scss">
#page-user-list {
  .user-list-filters {
    .vs__actions {
      position: absolute;
      right: 0;
      top: 50%;
      transform: translateY(-58%);
    }
  }
}
</style>
