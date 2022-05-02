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

    <!-- Save & Reset Button -->
    <div class="flex flex-wrap items-center justify-end">
      <vs-button class="ml-auto mt-2" @click="UpdateGeneralUser">Save Changes</vs-button>
      <vs-button class="ml-4 mt-2" type="border" color="warning">Reset</vs-button>
    </div>
  </vx-card>
</template>

<script>
import apiUser from '../../api/user';

export default {
  data () {
    return {
      profilePhoto: '',
      username: '',
      name: ''
    }
  },
  methods: {
    UpdateGeneralUser() {
      this.userId = localStorage.getItem('userId')
      apiUser
        .UpdateGeneralUser(this.userId, this.profilePhoto, this.username, this.name)
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
  },
  computed: {
    activeUserInfo () {
      return this.$store.state.AppActiveUser
    }
  }
}
</script>
