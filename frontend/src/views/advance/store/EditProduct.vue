<template>
    <div id="ecommerce-checkout-demo">
      <vx-card title="Edit New Product" subtitle="Be sure to check &quot;Product Details&quot; when you have finished" class="mb-base">
          <form data-vv-scope="edit-new-product">
              <div class="vx-row">
                  <div class="vx-col sm:w-1/3 w-full">
                      <vs-input
                          v-validate="'required|alpha_spaces'"
                          data-vv-as="field"
                          name="Name"
                          label="Name:"
                          v-model="Name"
                          class="w-full mt-0" />
                      <span v-show="errors.has('edit-new-product.Name')" class="text-danger">{{ errors.first('edit-new-product.Name') }}</span>
                  </div>
                  <div class="vx-col sm:w-2/3 w-full">
                      <p class="text-sm">Description:</p> 
                      <vs-textarea
                          style="height:125px"
                          v-validate="'required|alpha_spaces'"
                          name="Description"
                          v-model="Description"
                          class="w-full"
                          placeholder="Type Your Description"/>
                      <span v-show="errors.has('edit-new-product.description')" class="text-danger">{{ errors.first('edit-new-product.description') }}</span>
                  </div>
              </div>
              <div class="vx-row">
                  <div class="vx-col sm:w-1/6 w-full">
                      <vs-input
                          v-validate="'required'"
                          name="Price"
                          label="Price:"
                          v-model="Price"
                          class="w-full mt-5" />
                      <span v-show="errors.has('edit-new-product.price')" class="text-danger">{{ errors.first('edit-new-product.price') }}</span>
                  </div>
                  <div class="vx-col sm:w-1/6 w-full">
                      <vs-input
                          name="Shipping"
                          label="Shipping:"
                          v-model="Shipping"
                          class="w-full mt-5" />
                      <span v-show="errors.has('edit-new-product.shipping')" class="text-danger">{{ errors.first('edit-new-product.shipping') }}</span>
                  </div>
              </div>
              <div class="vx-row">
                  <div class="vx-col sm:w-1/6 w-full">
                      <vs-input
                          v-validate="'required'"
                          name="Weight"
                          label="Weight:"
                          v-model="Weight"
                          class="w-full mt-5" />
                      <span v-show="errors.has('edit-new-product.weight')" class="text-danger">{{ errors.first('edit-new-product.weight') }}</span>
                  </div>
                  <div class="vx-col sm:w-1/6 w-full">
                      <vs-input
                          v-validate="'required'"
                          name="Stock"
                          label="Stock:"
                          v-model="Stock"
                          class="w-full mt-5" />
                      <span v-show="errors.has('edit-new-product.stock')" class="text-danger">{{ errors.first('edit-new-product.stock') }}</span>
                  </div>
                  <div class="vx-col sm:w-1/6 w-full">
                      <vs-input
                          v-validate="'required'"
                          name="Discount"
                          label="Discount:"
                          v-model="Discount"
                          class="w-full mt-5" />
                      <span v-show="errors.has('edit-new-product.discount')" class="text-danger">{{ errors.first('edit-new-product.discount') }}</span>
                  </div>
                  <div class="vx-col">
                      <p class="mt-4 text-sm">Quantity</p>
                      <vx-tooltip text="Quantity">
                          <vs-input-number style="width:100px" v-model="number1" />
                      </vx-tooltip>
                      <span v-show="errors.has('edit-new-product.quantity')" class="text-danger">{{ errors.first('edit-new-product.quantity') }}</span>
                  </div>
                  <div class="mt-10"> 
                    <vs-button radius color="primary" type="border" icon-pack="feather" icon="icon-archive"></vs-button>
                  </div>
              </div>
              <div class="vx-row">
                <div class="vx-col">
                  <p class="mt-4 ml-2 text-sm">Pictures:</p>
                  <vs-upload
                      v-validate="'required'"
                      name="Picture"
                      v-model="Picture"
                      class="w-full" />
                  <span v-show="errors.has('edit-new-product.picture')" class="text-danger">{{ errors.first('edit-new-product.picture') }}</span>                  
                </div>
              </div>
              <div class="vx-row">
      <div class="vx-col w-full">
        <div class="mt-8 flex flex-wrap items-center justify-end">
          <vs-button class="ml-auto mt-2" type="border" color="danger" @click="reset_data">Cancel</vs-button>
          <vs-button class="ml-4 mt-2" @click="save_changes">Save</vs-button>
        </div>
      </div>
    </div>
          </form>
          <template slot="codeContainer">
            &lt;template&gt;
              &lt;vx-tooltip text=&quot;Tooltip Default&quot;&gt;
                &lt;vs-input-number v-model=&quot;number1&quot; /&gt;
              &lt;/vx-tooltip&gt;
            &lt;/template&gt;

            <vs-button radius color="primary" type="border" icon-pack="feather" icon="icon-search"></vs-button>

            &lt;script&gt;
            export default {
              data(){
                return {
                  switch1:true,
                  number1:10,
                }
              }
            }
            &lt;/script&gt;
          </template>
      </vx-card>
    </div>
</template>

<script>
import { FormWizard, TabContent } from 'vue-form-wizard'
import 'vue-form-wizard/dist/vue-form-wizard.min.css'

export default {
  data () {
    return {
      switch1: true,
      number1: 10
    }
  },
  components: {
    FormWizard,
    TabContent
  }
}
</script>
