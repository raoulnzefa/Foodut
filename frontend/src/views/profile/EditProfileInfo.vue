<template>
  <vx-card no-shadow>

    <!-- Bio -->
    <label class="text-sm">Bio</label>
    <vs-textarea v-model="bio" placeholder="Your bio..." />

    <!-- DOB -->
    <div class="mt-8">
      <label class="text-sm">Birth Date</label>
      <flat-pickr v-model="dob" :config="{ dateFormat: 'd F Y' }" class="w-full" />
    </div>

    <!-- Address -->
    <div class="mt-8">
      <label class="text-sm">Address</label>
      <vs-textarea style="height:150px" v-model="bio" placeholder="Your bio..." /> 
    </div>

    <!-- Mobile Number -->
    <label class="text-sm">Mobile  Number</label>
    <vs-input class="w-full mt-0" type="number" v-model="mobile" />

    <!-- Gender -->
    <div class="mt-8 mb-base">
      <label class="text-sm">Gender</label>
      <div class="mt-2">
        <vs-radio v-model="gender" vs-value="male" class="mr-4">Male</vs-radio>
        <vs-radio v-model="gender" vs-value="female" class="mr-4">Female</vs-radio>
      </div>
    </div>

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
