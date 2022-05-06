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
        <div
            :search-client="searchClient"
            index-name="instant_search" id="instant-search-demo">

            <div :hits-per-page.camel="9" />

            <div class="header mb-4">
                <div class="flex md:items-end items-center justify-between flex-wrap">

                    <p class="lg:inline-flex hidden font-semibold filters-label">Filters</p>

                    <div class="flex justify-between items-end flex-grow">
                        <!-- Stats -->
                        <div>
                            <p slot-scope="{  nbHits, processingTimeMS }" class="font-semibold md:block hidden">
                                {{ nbHits }} results found in {{ processingTimeMS }}ms
                            </p>
                        </div>
                    </div>
                </div>
            </div>

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
                        <li><vs-radio v-model="price" vs-value="low">Rp -10000</vs-radio></li>
                        <li><vs-radio v-model="price" vs-value="lowMedium">Rp 10000 - 50000</vs-radio></li>
                        <li><vs-radio v-model="price" vs-value="mediumHigh">Rp 500000 - 100000</vs-radio></li>
                        <vs-radio v-model="price" vs-value="high">Rp >=100000</vs-radio>
                      </ul>
                    </div>
                    <vs-divider />

                    <!-- CATEGORIES -->
                    <h6 class="font-bold mb-4">Category</h6>
                    <div>
                      <ul>
                        <li><vs-radio v-model="category" vs-value="chip">Chip</vs-radio></li>
                        <li><vs-radio v-model="category" vs-value="bread">Bread</vs-radio></li>
                        <li><vs-radio v-model="category" vs-value="beverage">Beverage</vs-radio></li>
                        <li><vs-radio v-model="category" vs-value="candy">Candy</vs-radio></li>
                      </ul>
                    </div>
                    <vs-divider />

                    <!-- Store -->
                    <h6 class="font-bold mb-4">Store</h6>
                      <vs-checkbox class="mb-1 ml-0">Indofood</vs-checkbox>
                      <vs-checkbox class="mb-1 ml-0">Snack Master</vs-checkbox>
                      <vs-checkbox class="mb-1 ml-0">Makaroni</vs-checkbox>
                    <vs-divider />

                    <vs-button class="w-full" color="danger" type="border">Remove Filters</vs-button>
                  </div>
                </vs-sidebar>
              </div>
                <!-- RIGHT COL -->
                <div :class="{'sidebar-spacer-with-margin': clickNotClose}">
                  <div class="mb-8">
                    <vs-input icon-no-border label-placeholder="Search" v-model="searchQuery" class="w-full input-rounded-full" icon="icon-search" icon-pack="feather" />
                    
                    <!-- <input type="text" class="vs-inputx vs-input--input large" style="border: 0px solid rgba(0, 0, 0, 0);" placeholder="Search here"> -->
                    <!-- <vs-input
                      v-validate="'required|email|min:3'"
                      data-vv-validate-on="blur"
                      name="email"
                      icon-no-border
                      icon="icon icon-user"
                      icon-pack="feather"
                      label-placeholder="Email"
                      v-model="email"
                      class="w-full"
                    /> -->
                  </div>
                  <div class="vx-row">
                    <div class="vx-col w-full sm:w-1/3 md:w-1/3 mb-base">
                      <vx-card>
                        <img :src="items.product" alt="product-img" class="w-full mb-3">
                        <div class="vx-row">
                          <div class="vx-col w-1/2">
                            <!-- <div class="vx-row">
                              <div class="vx-col w-1/12">
                                <feather-icon icon="StarIcon" svgClasses="h-5 w-5 text-warning"/>
                              </div>
                              <div class="vx-col">
                                <p>5</p>
                              </div>
                            </div> -->
                            <p class="text-sm">by storename</p>
                          </div>
                          <div class="vx-col w-1/2">
                            <p style="text-align:right">Rp 1000</p>
                          </div>
                        </div>
                        <p class="font-bold"> Product name</p>
                        <p class="text-sm">Description of product</p>
                        <div class="vx-row  mt-3">
                          <div class="vx-col w-1/2">
                            <vs-button class="w-full" color="primary" icon-pack="feather" icon="icon-shopping-cart">Add</vs-button>
                          </div>
                          <div class="vx-col w-1/2">
                            <vs-button class="w-full" color="primary" type="border" icon-pack="feather" icon="icon-shopping-bag">View</vs-button>
                          </div>
                        </div>
                      </vx-card>
                    </div>
                    <div class="vx-col w-full sm:w-1/3 md:w-1/3 mb-base">
                        <vx-card>
                          <img :src="items.product" alt="product-img" class="w-full mb-3">
                          <div class="vx-row">
                            <div class="vx-col w-1/2">
                              <!-- <div class="vx-row">
                                <div class="vx-col w-1/12">
                                  <feather-icon icon="StarIcon" svgClasses="h-5 w-5 text-warning"/>
                                </div>
                                <div class="vx-col">
                                  <p>5</p>
                                </div>
                              </div> -->
                              <p class="text-sm">by storename</p>
                            </div>
                            <div class="vx-col w-1/2">
                              <p style="text-align:right">Rp 1000</p>
                            </div>
                          </div>
                          <p class="font-bold"> Product name</p>
                          <p class="text-sm">Description of product</p>
                          <div class="vx-row  mt-3">
                            <div class="vx-col w-1/2">
                              <vs-button class="w-full" color="primary" icon-pack="feather" icon="icon-shopping-cart">Add</vs-button>
                            </div>
                            <div class="vx-col w-1/2">
                              <vs-button class="w-full" color="primary" type="border" icon-pack="feather" icon="icon-shopping-bag">View</vs-button>
                            </div>
                          </div>
                        </vx-card>
                    </div>
                    <div class="vx-col w-full sm:w-1/3 md:w-1/3 mb-base">
                        <vx-card>
                          <img :src="items.product" alt="product-img" class="w-full mb-3">
                          <div class="vx-row">
                            <div class="vx-col w-1/2">
                              <!-- <div class="vx-row">
                                <div class="vx-col w-1/12">
                                  <feather-icon icon="StarIcon" svgClasses="h-5 w-5 text-warning"/>
                                </div>
                                <div class="vx-col">
                                  <p>5</p>
                                </div>
                              </div> -->
                              <p class="text-sm">by storename</p>
                            </div>
                            <div class="vx-col w-1/2">
                              <p style="text-align:right">Rp 1000</p>
                            </div>
                          </div>
                          <p class="font-bold"> Product name</p>
                          <p class="text-sm">Description of product</p>
                          <div class="vx-row  mt-3">
                            <div class="vx-col w-1/2">
                              <vs-button class="w-full" color="primary" icon-pack="feather" icon="icon-shopping-cart">Add</vs-button>
                            </div>
                            <div class="vx-col w-1/2">
                              <vs-button class="w-full" color="primary" type="border" icon-pack="feather" icon="icon-shopping-bag">View</vs-button>
                            </div>
                          </div>
                        </vx-card>
                    </div>
                  </div>
                  <div class="vx-row">
                    <div class="vx-col w-full sm:w-1/3 md:w-1/3 mb-base">
                      <vx-card>
                        <img :src="items.product" alt="product-img" class="w-full mb-3">
                        <div class="vx-row">
                          <div class="vx-col w-1/2">
                            <!-- <div class="vx-row">
                              <div class="vx-col w-1/12">
                                <feather-icon icon="StarIcon" svgClasses="h-5 w-5 text-warning"/>
                              </div>
                              <div class="vx-col">
                                <p>5</p>
                              </div>
                            </div> -->
                            <p class="text-sm">by storename</p>
                          </div>
                          <div class="vx-col w-1/2">
                            <p style="text-align:right">Rp 1000</p>
                          </div>
                        </div>
                        <p class="font-bold"> Product name</p>
                        <p class="text-sm">Description of product</p>
                        <div class="vx-row  mt-3">
                          <div class="vx-col w-1/2">
                            <vs-button class="w-full" color="primary" icon-pack="feather" icon="icon-shopping-cart">Add</vs-button>
                          </div>
                          <div class="vx-col w-1/2">
                            <vs-button class="w-full" color="primary" type="border" icon-pack="feather" icon="icon-shopping-bag">View</vs-button>
                          </div>
                        </div>
                      </vx-card>
                    </div>
                    <div class="vx-col w-full sm:w-1/3 md:w-1/3 mb-base">
                        <vx-card>
                          <img :src="items.product" alt="product-img" class="w-full mb-3">
                          <div class="vx-row">
                            <div class="vx-col w-1/2">
                              <!-- <div class="vx-row">
                                <div class="vx-col w-1/12">
                                  <feather-icon icon="StarIcon" svgClasses="h-5 w-5 text-warning"/>
                                </div>
                                <div class="vx-col">
                                  <p>5</p>
                                </div>
                              </div> -->
                              <p class="text-sm">by storename</p>
                            </div>
                            <div class="vx-col w-1/2">
                              <p style="text-align:right">Rp 1000</p>
                            </div>
                          </div>
                          <p class="font-bold"> Product name</p>
                          <p class="text-sm">Description of product</p>
                          <div class="vx-row  mt-3">
                            <div class="vx-col w-1/2">
                              <vs-button class="w-full" color="primary" icon-pack="feather" icon="icon-shopping-cart">Add</vs-button>
                            </div>
                            <div class="vx-col w-1/2">
                              <vs-button class="w-full" color="primary" type="border" icon-pack="feather" icon="icon-shopping-bag">View</vs-button>
                            </div>
                          </div>
                        </vx-card>
                    </div>
                    <div class="vx-col w-full sm:w-1/3 md:w-1/3 mb-base">
                        <vx-card>
                          <img :src="items.product" alt="product-img" class="w-full mb-3">
                          <div class="vx-row">
                            <div class="vx-col w-1/2">
                              <!-- <div class="vx-row">
                                <div class="vx-col w-1/12">
                                  <feather-icon icon="StarIcon" svgClasses="h-5 w-5 text-warning"/>
                                </div>
                                <div class="vx-col">
                                  <p>5</p>
                                </div>
                              </div> -->
                              <p class="text-sm">by storename</p>
                            </div>
                            <div class="vx-col w-1/2">
                              <p style="text-align:right">Rp 1000</p>
                            </div>
                          </div>
                          <p class="font-bold"> Product name</p>
                          <p class="text-sm">Description of product</p>
                          <div class="vx-row  mt-3">
                            <div class="vx-col w-1/2">
                              <vs-button class="w-full" color="primary" icon-pack="feather" icon="icon-shopping-cart">Add</vs-button>
                            </div>
                            <div class="vx-col w-1/2">
                              <vs-button class="w-full" color="primary" type="border" icon-pack="feather" icon="icon-shopping-bag">View</vs-button>
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

export default {
  components: {
    ItemGridView: () => import('./components/ItemGridView.vue'),
    ItemListView: () => import('./components/ItemListView.vue')
  },
  data () {
    return {
      items: {
        product: require('@/assets/images/dummy/Doritos.jpg'),
        star: require('@/assets/images/raty/star-on-2.png')
      },
      // Filter Sidebar
      isFilterSidebarActive: true,
      clickNotClose: true,
      currentItemView: 'item-grid-view',
      numericItems: [
        { label: 'All' },
        { label: '<= $10', end: 10 },
        { label: '$10 - $100', start: 10, end: 100 },
        { label: '$100 - $500', start: 100, end: 500 },
        { label: '>= $500', start: 500 }
      ],
      Categories: [
        'hierarchicalCategories.lvl0',
        'hierarchicalCategories.lvl1',
        'hierarchicalCategories.lvl2',
        'hierarchicalCategories.lvl3'
      ]
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

