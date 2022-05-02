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
      <vs-button class="ml-auto mt-2">Save Changes</vs-button>
      <vs-button class="ml-4 mt-2" type="border" color="warning">Reset</vs-button>
    </div>
  </vx-card>
</template>

<script>
import flatPickr from 'vue-flatpickr-component'
import 'flatpickr/dist/flatpickr.css'
import vSelect from 'vue-select'

export default {
  components: {
    flatPickr,
    vSelect
  },
  data () {
    return {
      bio: this.$store.state.AppActiveUser.about,
      dob: null,
      address: '',
      gender: 'male',
      mobile: ''
    }
  },
  computed: {
    activeUserInfo () {
      return this.$store.state.AppActiveUser
    }
  }
}
</script>
