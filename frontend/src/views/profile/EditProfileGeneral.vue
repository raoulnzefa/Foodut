<template>
  <vx-card no-shadow>
    <!-- Img Row -->
    <div class="flex flex-wrap items-center mb-base">
      <vs-avatar :src="activeUserInfo.photoURL" size="70px" class="mr-4 mb-4" />
      <div>
        <vs-button class="mr-4 sm:mb-0 mb-2">Upload photo</vs-button>
        <vs-button type="border" color="danger">Remove</vs-button>
        <p class="text-sm mt-2">Allowed JPG, GIF or PNG. Max size of 800kB</p>
      </div>
    </div>

    <!-- Username -->
    <label class="text-sm">Username</label>
    <vs-input class="w-full mb-base" v-model="username"></vs-input>

    <!-- Name -->
    <label class="text-sm">Name</label>
    <vs-input class="w-full mb-base" v-model="name"></vs-input>

    <!-- Email -->
    <label class="text-sm">Email</label>
    <vs-input class="w-full" v-model="email"></vs-input>

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
      username: 'johny_01',
      name: this.$store.state.AppActiveUser.displayName,
      email: 'john@admin.com',
      company: 'SnowMash Technologies Pvt Ltd'
    }
  },
  computed: {
    activeUserInfo () {
      return this.$store.state.AppActiveUser
    }
  }
}
</script>
