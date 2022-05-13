<template>
  <vx-card no-shadow>

    <!-- Username -->
    <label class="text-sm">Store Name</label>
    <vs-input class="w-full mb-base" v-model="storename"></vs-input>

    <!-- Name -->
    <label class="text-sm">City</label>
    <vs-input class="w-full mb-base" v-model="city"></vs-input>

    <!-- Save & Reset Button -->
    <div class="flex flex-wrap items-center justify-end">
      <vs-button class="ml-auto mt-2" @click="UpdateInfoUser">Save Changes</vs-button>
      <vs-button class="ml-4 mt-2" type="border" color="warning">Reset</vs-button>
    </div>
  </vx-card>
</template>

<script>
import flatPickr from 'vue-flatpickr-component'
import 'flatpickr/dist/flatpickr.css'
import vSelect from 'vue-select'
import apiUser from '../../../../api/user'

export default {
  components: {
    flatPickr,
    vSelect
  },
  data () {
    return {
      storename: "",
      city: ""
    }
  },
  computed: {
    activeUserInfo () {
      return this.$store.state.AppActiveUser
    }
  },
  methods: {
    UpdateInfoUser() {
      this.userId = localStorage.getItem('userId')
      apiUser
        .UpdateInfoSeller(this.userId, this.storename, this.city)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to update password',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to update password',
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
