<template>
  <div id="ecommerce-checkout-demo">
    <add-categories /><br>
    <vx-card
      title="Add New Product"
      subtitle="Be sure to check Product Details when you have finished"
      class="mb-base">
      <form data-vv-scope="add-new-product">
        <div class="vx-row">
          <div class="vx-col sm:w-1/2 w-full">
            <div class="vx-row">
              <div class="vx-col w-full">
                <vs-input
                  v-validate="'required|alpha_spaces'"
                  data-vv-as="field"
                  name="Name"
                  label="Name:"
                  v-model="name"
                  class="w-full mt-0"
                />
                <span
                  v-show="errors.has('add-new-product.productName')"
                  class="text-danger"
                  >{{ errors.first("add-new-product.productName") }}</span
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
                  v-show="errors.has('add-new-product.productPrice')"
                  class="text-danger"
                  >{{ errors.first("add-new-product.productPrice") }}</span
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
                  v-show="errors.has('add-new-product.productCategory')"
                  class="text-danger"
                  >{{ errors.first("add-new-product.productCategory") }}</span>
              </div>
              <div class="vx-col sm:w-1/3 w-full">
                <p class="mt-4 text-sm">Stock:</p>
                <vs-input-number style="width: 100px" v-model="stock" />
                <span
                  v-show="errors.has('add-new-product.productStock')"
                  class="text-danger"
                  >{{ errors.first("add-new-product.productStock") }}</span
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
              v-show="errors.has('add-new-product.productDescription')"
              class="text-danger"
              >{{ errors.first("add-new-product.productDescription") }}</span
            >
          </div>
        </div>
        <div class="vx-row">
          <div class="vx-col sm:w-1/5 w-full">
            <p class="mt-4 text-sm">Picture:</p>
            <div class="upload-img mt-5">
              <input type="file" class="hidden" ref="uploadImgInput" @change="updateNameFilePicture" accept="image/*" multiple="multiple">
              <vs-button @click="$refs.uploadImgInput.click()">Upload Image</vs-button>
            </div>
          </div>
          <div class="vx-col sm:w-1/6 w-full">
            <p class="mt-4 text-sm">Total Picture</p>
            <vs-input style="padding-top:18px; width:50px" v-validate="'required'" name="totalPicture" v-model="total_picture" readonly="readonly"/>
          </div>
        </div>
      </form>
      <!-- Save & Reset Button -->
      <div class="vx-row">
        <div class="vx-col w-full">
          <div class="mt-8 flex flex-wrap items-center justify-end">
            <!-- :disabled="!validateForm" -->
            <vs-button class="ml-auto mt-2" @click="createProduct">Create</vs-button>
            <vs-button class="ml-4 mt-2" type="border" color="warning">Cancel</vs-button>
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
import AddCategories from './categories/AddCategories.vue'
import apiProduct from '../../../api/product'
import apiCategory from '../../../api/category'

export default {
  components: {
    vSelect,
    FormWizard,
    TabContent,
    AddCategories
  },
  data () {
    return {
      switch1: true,
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
      total_picture: 0,
    }
  },
  methods: {
    updateNameFilePicture(){
      // Read selected files
      const files = this.$refs.uploadImgInput.files;
      const totalfiles = this.$refs.uploadImgInput.files.length;
      console.log(files, totalfiles)
      this.total_picture = totalfiles
    },
    createProduct() {
      this.sellerId = localStorage.getItem('userId')
      const totalPicture = this.$refs.uploadImgInput.files;
      const pictureArray = []
      for (let index = 0; index < totalPicture.length; index++) {
        pictureArray.push(totalPicture[index].name)
      }
      console.log("Data Product Baru")
      console.log(this.name, this.price, this.stock, this.sellerId,  this.selectedCategory.value, this.description, pictureArray)
      apiProduct
        .AddProduct(this.name, this.price, this.stock, this.sellerId,  this.selectedCategory.value, this.description, pictureArray)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to create product',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to create product',
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
  }
}
</script>
