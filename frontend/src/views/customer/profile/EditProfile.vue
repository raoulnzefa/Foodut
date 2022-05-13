<template>
  <vs-tabs :position="isSmallerScreen ? 'top' : 'left'" class="tabs-shadow-none" id="profile-tabs" :key="isSmallerScreen">

    <!-- GENERAL -->
    <vs-tab icon-pack="feather" icon="icon-user" :label="!isSmallerScreen ? 'General' : ''">
      <div class="tab-general md:ml-4 md:mt-0 mt-4 ml-0">
        <edit-profile-general />
      </div>
    </vs-tab>
    <vs-tab icon-pack="feather" icon="icon-info" :label="!isSmallerScreen ? 'Info' : ''">
      <div class="tab-info md:ml-4 md:mt-0 mt-4 ml-0">
        <edit-profile-info />
      </div>
    </vs-tab>
    <vs-tab icon-pack="feather" icon="icon-lock" :label="!isSmallerScreen ? 'Change Password' : ''">
      <div class="tab-change-pwd md:ml-4 md:mt-0 mt-4 ml-0">
        <edit-profile-change-password />
      </div>
    </vs-tab>
  </vs-tabs>
</template>

<script>
import EditProfileGeneral from './components/General.vue'
import EditProfileInfo from './components/Info.vue'
import EditProfileChangePassword from './components/ChangePassword.vue'
import apiUser from '../../../api/user'

export default {
  components: {
    EditProfileGeneral,
    EditProfileInfo,
    EditProfileChangePassword
  },
  data () {
    return {
      

    }
  },
  method:{
    GetUserID(userId){
        apiUser 
        .GetUserById(userId)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to get user',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
          }else{
            this.$vs.notify({
              title: 'Success',
              text: 'Succes to get user',
              color: 'success',
              iconPack: 'feather',
              icon: 'icon-check'
            })
            console.log(response)
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
    isSmallerScreen () {
      return this.$store.state.windowWidth < 768
    },
    infoUser() {
      const userId = localStorage.getItem('userId')
      this.GetUserID(userId)
      return "EditProfileInfo"
    }
  }
}
</script>

<style lang="scss">
#profile-tabs {
  .vs-tabs--content {
    padding: 0;
  }
}
</style>
