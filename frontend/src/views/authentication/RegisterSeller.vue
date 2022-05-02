<!-- =========================================================================================
File Name: RegisterJWT.vue
Description: Register Page for JWT
----------------------------------------------------------------------------------------
Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
  Author: Pixinvent
Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
  <div class="clearfix">
    <vs-input
      v-validate="'required|alpha_dash|min:3'"
      data-vv-validate-on="blur"
      label-placeholder="Username"
      name="username"
      placeholder="Username"
      v-model="username"
      class="w-full" />
    <span class="text-danger text-sm">{{ errors.first('username') }}</span>

    <vs-input
      v-validate="'required|min:3'"
      data-vv-validate-on="blur"
      label-placeholder="Name"
      name="name"
      placeholder="Name"
      v-model="name"
      class="w-full" />
    <span class="text-danger text-sm">{{ errors.first('name') }}</span>

    <vs-input
      v-validate="'required|email'"
      data-vv-validate-on="blur"
      name="email"
      type="email"
      label-placeholder="Email"
      placeholder="Email"
      v-model="email"
      class="w-full mt-6" />
    <span class="text-danger text-sm">{{ errors.first('email') }}</span>

    <vs-input
      ref="password"
      type="password"
      data-vv-validate-on="blur"
      v-validate="'required|min:5|max:20'"
      name="password"
      label-placeholder="Password"
      placeholder="Password"
      v-model="password"
      class="w-full mt-6" />
    <span class="text-danger text-sm">{{ errors.first('password') }}</span>

    <vs-input
      v-validate="'required|min:3'"
      data-vv-validate-on="blur"
      label-placeholder="Address"
      name="address"
      placeholder="Address"
      v-model="address"
      class="w-full" />
    <span class="text-danger text-sm">{{ errors.first('address') }}</span>

    <vs-input
      v-validate="'required|min:3'"
      data-vv-validate-on="blur"
      label-placeholder="Store Name"
      name="storeName"
      placeholder="Store Name"
      v-model="storeName"
      class="w-full" />
    <span class="text-danger text-sm">{{ errors.first('storeName') }}</span>

    <vs-input
      v-validate="'required|min:3'"
      data-vv-validate-on="blur"
      label-placeholder="City"
      name="city"
      placeholder="City"
      v-model="city"
      class="w-full" />
    <span class="text-danger text-sm">{{ errors.first('city') }}</span>

    <vs-checkbox v-model="isTermsConditionAccepted" class="mt-6">I accept the terms & conditions.</vs-checkbox>
    <vs-button  type="border" to="/" class="mt-6">Login</vs-button>
    <vs-button class="float-right mt-6" @click="registerSeller" :disabled="!validateForm">Register</vs-button>
  </div>
</template>

<script>
import apiUser from '../../api/user'

export default {
  data () {
    return {
      username: '',
      name: '',
      email: '',
      password: '',
      address: '',
      storeName: '',
      city: '',
      isTermsConditionAccepted: false
    }
  },
  computed: {
    validateForm () {
      return !this.errors.any() && this.username !== '' && this.name !== '' && this.email !== '' && this.password !== '' && this.address !== '' && this.storeName !== '' && this.city !== '' && this.isTermsConditionAccepted === true
    }
  },
  methods: {
    checkLogin () {
      // If user is already logged in notify
      if (this.$store.state.auth.isUserLoggedIn()) {
        this.$vs.notify({
          title: 'Login Attempt',
          text: 'You are already logged in!',
          iconPack: 'feather',
          icon: 'icon-alert-circle',
          color: 'warning'
        })
        return false
      }
      return true
    },
    registerSeller () {
      // If form is not validated or user is already login return
      if (!this.validateForm || !this.checkLogin()) return
      apiUser
        .RegisterSeller(this.username, this.email, this.name, this.password, this. address, this.storeName, this.city)
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to login',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
            return
          }else{
              this.$vs.notify({
              title: 'Success',
              text: 'Succes to create seller account',
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
