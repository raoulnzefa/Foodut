<template>
  <vx-card no-shadow>
    <label class="text-sm">Old Password</label>
    <vs-input class="w-full mb-base" v-model="old_password" />

    <label class="text-sm">New Password</label>
    <vs-input class="w-full mb-base" v-model="new_password" />

    <label class="text-sm">Confirm New Password</label>
    <vs-input class="w-full mb-base" v-model="confirm_new_password" />

    <!-- Save & Reset Button -->
    <div class="flex flex-wrap items-center justify-end">
      <vs-button class="ml-auto mt-2">Save Changes</vs-button>
      <vs-button class="ml-4 mt-2" type="border" color="warning">Reset</vs-button>
    </div>
  </vx-card>
</template>

<script>
export default {
  data () {
    return {
      old_password: '',
      new_password: '',
      confirm_new_password: ''
    }
  },
  computed: {
    activeUserInfo () {
      return this.$store.state.AppActiveUser
    }
  }
}
</script>
