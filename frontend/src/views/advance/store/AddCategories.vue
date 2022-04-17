<template>
  <div id="ecommerce-checkout-demo">
    <list-categories /><br>
    <vx-card
      title="Add New Categories"
      class="mb-base">
      <form data-vv-scope="add-new-categories">
        <div class="vx-row">
          <div class="vx-col sm:w-3/4 w-full">
            <vs-input
              style="padding-top:15px"
              v-validate="'required|alpha_spaces'"
              data-vv-as="field"
              name="Category"
              label="Category:"
              v-model="category"
              class="w-full mt-0"
            />
            <span
              v-show="errors.has('add-new-categories.productCategory')"
              class="text-danger"
              >{{ errors.first("add-new-categories.productCategory") }}</span
            >
          </div>
          <div class="vx-col sm:w-1/4 w-full">
            <div class="mt-8 flex flex-wrap items-center justify-end">
              <!-- :disabled="!validateForm" -->
              <vs-button class="ml-auto mt-2" @click="CreateCategories">Create</vs-button>
              <vs-button class="ml-4 mt-2" type="border" color="warning" @click="reset_data">Cancel</vs-button>
            </div>
          </div>
        </div>
      </form>
      <!-- Save & Reset Button -->
      <div class="vx-row">
        <div class="vx-col w-full">
          
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
import apiCategories from '../../../api/categories'

export default {
  components: {
    FormWizard,
    TabContent
  },
  data () {
    return {
      switch1: true,
      category: ""
    }
  },
  methods: {
    CreateCategories(){
      apiCategories
        .AddCategories(this.category)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to create categories',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to create categories',
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
  }
}
</script>
