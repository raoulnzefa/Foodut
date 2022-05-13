<!-- =========================================================================================
    File Name: ForgotPassword.vue
    Description: FOrgot Password Page
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->

<template>
    <div class="h-screen flex w-full bg-img">
        <div class="vx-col sm:w-3/5 md:w-3/5 lg:w-3/4 xl:w-3/5 mx-auto self-center">
            <vx-card>
                <div slot="no-body" class="full-page-bg-color">
                    <div class="vx-row">
                        <div class="vx-col hidden sm:hidden md:hidden lg:block lg:w-1/2 mx-auto self-center">
                            <img src="@/assets/images/pages/reset-password.png" alt="login" class="mx-auto">
                        </div>
                        <div class="vx-col sm:w-full md:w-full lg:w-1/2 mx-auto self-center  d-theme-dark-bg">
                            <div class="p-8">
                                <div class="vx-card__title mb-8">
                                    <h4 class="mb-4">Reset Password</h4>
                                    <p>Please enter your new password.</p>
                                </div>
                                <vs-input type="email" icon-no-border icon="icon icon-mail" icon-pack="feather" label-placeholder="Email" v-model="email" class="w-full mb-6" />
                                <vs-input type="password" icon-no-border icon="icon icon-lock" icon-pack="feather" label-placeholder="New Password" v-model="new_password" class="w-full mb-6" />
                                <vs-input type="password" icon-no-border icon="icon icon-lock" icon-pack="feather" label-placeholder="Confirm Password" v-model="confirm_new_password" class="w-full mb-8" />

                                <div class="flex flex-wrap justify-between flex-col-reverse sm:flex-row">
                                    <vs-button type="border" to="/" class="w-full sm:w-auto mb-8 sm:mb-auto mt-3 sm:mt-auto">Go Back To Login</vs-button>
                                    <vs-button class="w-full sm:w-auto" @click="UpdatePasswordUser">Reset</vs-button>
                                </div>

                            </div>
                        </div>
                    </div>
                </div>
            </vx-card>
        </div>
    </div>
</template>

<script>
import apiUser from '../../api/user';

export default {
  data () {
    return {
      email: '',
      new_password: '',
      confirm_new_password: ''
    }
  },
  methods: {
    UpdatePasswordUser() {
      this.userId = localStorage.getItem('userId')
      if(this.email != '' && this.new_password == this.confirm_new_password){
        //searching email milik user brp
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
  }
}
</script>
