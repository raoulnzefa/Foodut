<template>
  <div id="ecommerce-checkout-demo">
    <vx-card
      title="Edit New Product"
      subtitle="Be sure to check Product Details when you have finished"
      class="mb-base">
      <form data-vv-scope="product">
        <div class="vx-row">
          <div class="vx-col sm:w-1/2 w-full">
            <div class="vx-row">
              <div class="vx-col sm:w-1/4 w-full">
                <vs-input
                  v-validate="'required|alpha_spaces'"
                  data-vv-as="field"
                  name="Id"
                  label="Id Product:"
                  v-model="idProduct"
                  class="w-full mt-0"
                  readonly="readonly"
                />
                <span
                  v-show="errors.has('product.Name')"
                  class="text-danger"
                  >{{ errors.first("product.productId") }}</span
                >
              </div>
              <div class="vx-col sm:w-3/4 w-full">
                <vs-input
                  v-validate="'required|alpha_spaces'"
                  data-vv-as="field"
                  name="Name"
                  label="Name:"
                  v-model="name"
                  class="w-full mt-0"
                />
                <span
                  v-show="errors.has('product.Name')"
                  class="text-danger"
                  >{{ errors.first("product.Name") }}</span
                >
              </div>
            </div>
            <div class="vx-row">
              <div class="vx-col sm:w-1/3 w-full">
                <vs-input
                  v-validate="'required'"
                  name="Price"
                  label="Price:"
                  v-model="price"
                  class="w-full mt-4"
                />
                <span
                  v-show="errors.has('product.price')"
                  class="text-danger"
                  >{{ errors.first("product.price") }}</span
                >
              </div>
              <div class="vx-col sm:w-1/3 w-full">
                <p class="mt-4 text-sm">Category:</p>
                <v-select
                  id="category"
                  v-model="selectedCategory"
                  :options="category"
                  label="text"
                  :selectable="option => selectedCategory"
                  v-validate="'required'"
                  name="selectedCategory"
                  :clearable="false" 
                  class="w-full mt-1"
                  items="category"
                />
                <span
                  v-show="errors.has('product.category')"
                  class="text-danger"
                  >{{ errors.first("product.category") }}</span
                >
              </div>
              <div class="vx-col sm:w-1/3 w-full">
                <p class="mt-4 text-sm">Stock:</p>
                <vs-input-number style="width: 100px" v-model="stock" />
                <span
                  v-show="errors.has('product.stock')"
                  class="text-danger"
                  >{{ errors.first("product.stock") }}</span
                >
              </div>
            </div>
          </div>
          <div class="vx-col sm:w-1/2 w-full">
            <p class="text-sm">Description:</p>
            <vs-textarea
              rows="5"
              v-validate="'required|alpha_spaces'"
              name="Description"
              v-model="description"
              class="w-full"
              placeholder="Type Your Description"
            />
            <span
              v-show="errors.has('product.description')"
              class="text-danger"
              >{{ errors.first("product.description") }}</span
            >
          </div>
        </div>
      </form>
      <!-- Save & Reset Button -->
      <div class="vx-row">
        <div class="vx-col w-full">
          <div class="mt-8 flex flex-wrap items-center justify-end">
            <vs-button class="ml-auto mt-2" @click="UpdateProduct">Save</vs-button>
            <vs-button class="ml-4 mt-2" type="border" color="warning" @click="reset_data">Cancel</vs-button>
          </div>
        </div>
      </div>
      <template slot="codeContainer">
        &lt;template&gt; &lt;vx-tooltip text=&quot;Tooltip Default&quot;&gt;
        &lt;vs-input-number v-model=&quot;picture_items:1&quot; /&gt;
        &lt;/template&gt;

        <vs-button
          radius
          color="primary"
          type="border"
          icon-pack="feather"
          icon="icon-search"
        ></vs-button>

        &lt;script&gt; export default { data(){ return { switch1:true,
        picture_items:1} } } &lt;/script&gt;
      </template>
    </vx-card>
  </div>
</template>

<script>
import vSelect from 'vue-select'
import { FormWizard, TabContent } from 'vue-form-wizard'
import 'vue-form-wizard/dist/vue-form-wizard.min.css'
import apiProduct from '../../../api/product'
import apiCategory from '../../../api/category'

export default {
  data () {
    return {
      switch1: true,
      idProduct: 0,
      name: "",
      price: 0,
      category: [],
      selectedCategory: {
        text: "",
        value: 0,
        disabled: false,
        divider: false,
        header: ""
      },
      stock: 0,
      description: "",
      picture: [],
      total_picture: 0
    }
  },
  methods: {
    UpdateProduct() {
      this.sellerId = localStorage.getItem('userId')
      console.log("Update Data Product")
      console.log('id: ', this.idProduct)
      console.log('name: ', this.name)
      console.log('price: ', this.price)
      console.log('stock: ', this.stock)
      console.log('selected category: ', this.selectedCategory.value)
      console.log('description: ', this.description)
      apiProduct
        .UpdateProduct(this.idProduct, this.name, this.price, this.stock, this.selectedCategory.value, this.description)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to update product',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to update product',
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
  },
  components: {
    vSelect,
    FormWizard,
    TabContent
  },
  mounted () {
    apiCategory
      .GetAllCategories()
      .then((response) => { 
        this.category = []
        for(let i=0; i<response.length; i++){
          this.category.push({
            text: response[i].productCategory,
            value: response[i].id,
            disabled: false,
            divider: false,
            header: ""
          }) 
        }
        console.log(response)
        console.log(this.category)
        console.log(this.selectedCategory)
      })
      .catch((error) => { console.log('Error get all categories!', error)})
    apiProduct
      .GetProductById(this.$route.params.product_id)
      .then((response) => { 
        this.idProduct = response[0].id
        this.name = response[0].productName
        this.price = response[0].productPrice
        const temp = response[0].categoryId
        apiCategory
          .GetAllCategories()
          .then((response) => { 
            this.category = []
            for(let i=0; i<response.length; i++){
              if(temp == response[i].id){
                this.selectedCategory = {
                  text: response[i].productCategory,
                  value: response[i].id,
                  disabled: false,
                  divider: false,
                  header: ""
                }
              }
              this.category.push({
                text: response[i].productCategory,
                value: response[i].id,
                disabled: false,
                divider: false,
                header: ""
              }) 
            }
            console.log(response)
            console.log(this.category)
            console.log(this.selectedCategory)
          })
          .catch((error) => { console.log('Error get all categories!', error)})
        console.log('kategori: ',this.selectedCategory)
        this. stock = response[0].productStock
        this.description = response[0].description
        // console.log(this.product)
      })
      .catch((error) => { console.log('Error get id product!', error)})
  }
}
</script>
