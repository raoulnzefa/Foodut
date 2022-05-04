<template>
  <vx-card no-shadow>
    <label class="text-sm">Old Password</label>
    <vs-input data-vv-validate-on="blur"
      v-validate="'required|min:6|max:10'"
      type="password"
      name="old_password"
      icon-no-border
      icon="icon icon-lock"
      icon-pack="feather"
      class="w-full mb-base" 
      v-model="old_password" />

    <label class="text-sm">New Password</label>
    <vs-input 
      v-validate="'required|min:6|max:10'"
      type="password"
      name="new_password"
      icon-no-border
      icon="icon icon-lock"
      icon-pack="feather"
      class="w-full mb-base" 
      v-model="new_password" />

    <label class="text-sm">Confirm New Password</label>
    <vs-input v-validate="'required|min:6|max:10'"
      type="password"
      name="confirm_new_password"
      icon-no-border
      icon="icon icon-lock"
      icon-pack="feather" 
      class="w-full mb-base" 
      v-model="confirm_new_password" />

    <!-- Save & Reset Button -->
    <div class="flex flex-wrap items-center justify-end">
      <vs-button class="ml-auto mt-2" @click="UpdatePasswordUser">Save Changes</vs-button>
      <vs-button class="ml-4 mt-2" type="border" color="warning">Reset</vs-button>
    </div>
  </vx-card>
</template>

<script>
import apiUser from '../../../../api/user'

export default {
  data () {
    return {
      old_password: '',
      new_password: '',
      confirm_new_password: ''
    }
  },
  methods: {
    UpdatePasswordUser() {
      this.userId = localStorage.getItem('userId')
      if(this.old_password != this.new_password && this.new_password == this.confirm_new_password){
        apiUser
          .UpdatePasswordUser(this.userId, this.new_password)
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
      }else{
        this.$vs.notify({
          title: 'Error',
          text: 'Failed to update password',
          iconPack: 'feather',
          icon: 'icon-alert-circle',
          color: 'danger'
        })
      }
    }
  },
  computed: {
    activeUserInfo () {
      return this.$store.state.AppActiveUser
    }
  }
}
</script>
