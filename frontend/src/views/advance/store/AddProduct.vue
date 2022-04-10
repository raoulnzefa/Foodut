<template>
  <div id="ecommerce-checkout-demo">
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
                  v-model="Name"
                  class="w-full mt-0"
                />
                <span
                  v-show="errors.has('add-new-product.Name')"
                  class="text-danger"
                  >{{ errors.first("add-new-product.Name") }}</span
                >
              </div>
            </div>
            <div class="vx-row">
              <div class="vx-col sm:w-1/3 w-full">
                <vs-input
                  v-validate="'required'"
                  name="Price"
                  label="Price:"
                  v-model="Price"
                  class="w-full mt-5"
                />
                <span
                  v-show="errors.has('add-new-product.price')"
                  class="text-danger"
                  >{{ errors.first("add-new-product.price") }}</span
                >
              </div>
              <div class="vx-col sm:w-1/3 w-full">
                <vs-input
                  v-validate="'required'"
                  name="Category"
                  label="Category:"
                  v-model="Category"
                  class="w-full mt-5"
                />
                <span
                  v-show="errors.has('add-new-product.category')"
                  class="text-danger"
                  >{{ errors.first("add-new-product.category") }}</span
                >
              </div>
              <div class="vx-col sm:w-1/3 w-full">
                <p class="mt-4 text-sm">Stock:</p>
                <vs-input-number style="width: 100px" v-model="stock" />
                <span
                  v-show="errors.has('add-new-product.stock')"
                  class="text-danger"
                  >{{ errors.first("add-new-product.stock") }}</span
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
              v-model="Description"
              class="w-full"
              placeholder="Type Your Description"
            />
            <span
              v-show="errors.has('add-new-product.description')"
              class="text-danger"
              >{{ errors.first("add-new-product.description") }}</span
            >
          </div>
        </div>
        <div class="vx-row">
          <div class="vx-col">
            <p class="mt-4 text-sm">Picture Items:</p>
            <vs-input-number style="width: 100px" v-model="picture_items" />
            <span
              v-show="errors.has('add-new-product.items')"
              class="text-danger"
              >{{ errors.first("add-new-product.items") }}</span
            >
          </div>
          <div class="mt-10">
            <vs-button
              v-on:click="getPictureItems()"
              radius
              color="primary"
              type="border"
              icon-pack="feather"
              icon="icon-archive"
            ></vs-button>
          </div>
        </div>
        <div class="vx-row">
          <div class="vx-col w-full">
            <div class="upload-img mt-5" v-if="!dataImg">
              <input type="file" class="hidden" ref="uploadImgInput" @change="updateCurrImg" accept="image/*">
              <vs-button @click="$refs.uploadImgInput.click()">Upload Image</vs-button>
            </div>
          </div>
        </div>
      </form>
      <!-- Save & Reset Button -->
      <div class="vx-row">
        <div class="vx-col w-full">
          <div class="mt-8 flex flex-wrap items-center justify-end">
            <!-- :disabled="!validateForm" -->
            <vs-button class="ml-auto mt-2" @click="CreateProduct">Create</vs-button>
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
import { FormWizard, TabContent } from 'vue-form-wizard'
import 'vue-form-wizard/dist/vue-form-wizard.min.css'
import apiProduct from '../../../api/product'

export default {
  data () {
    return {
      switch1: true,
      stock:10,
      picture_items: 1
    }
  },
  methods: {
    Createproduct() {
      apiProduct
        .AddProduct(
          // this.productName, this.productPrice, this.productStock, this.sellerId, this.productCategory, this.productPicture
          
          )
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to create product',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
            return
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
    },
    getPictureItems () {
      alert(this.picture_items)
    }
  },
  components: {
    FormWizard,
    TabContent
  }
}
</script>
