<template>
  <div>
    <vs-input
      v-validate="'required|email|min:3'"
      data-vv-validate-on="blur"
      name="email"
      icon-no-border
      icon="icon icon-user"
      icon-pack="feather"
      label-placeholder="Email"
      v-model="email"
      class="w-full"
    />
    <span class="text-danger text-sm">{{ errors.first("email") }}</span>

    <vs-input
      data-vv-validate-on="blur"
      v-validate="'required|min:6|max:10'"
      type="password"
      name="password"
      icon-no-border
      icon="icon icon-lock"
      icon-pack="feather"
      label-placeholder="Password"
      v-model="password"
      class="w-full mt-6"
    />
    <span class="text-danger text-sm">{{ errors.first("password") }}</span>
    <div class="flex flex-wrap justify-between my-5">
      <vs-checkbox v-model="checkbox_remember_me" class="mb-3"
        >Remember Me</vs-checkbox
      >
      <router-link to="/forgot-password">Forgot Password?</router-link>
    </div>
    <div class="flex flex-wrap justify-between mb-3">
      <vs-button type="border" @click="registerUser">Register</vs-button>
      <vs-button :disabled="!validateForm" @click="loginJWT">Login</vs-button>
    </div>
  </div>
</template>

<script>
import apiUser from '../../../../api/user'

export default {
  data () {
    return {
      email: '',
      password: '',
      checkbox_remember_me: false
    }
  },
  computed: {
    validateForm () {
      return !this.errors.any() && this.email !== '' && this.password !== ''
    }
  },
  methods: {
    redirectUser(){
        //Get User ID from local storage that has been set from login api
        const userID = localStorage.getItem('userId');
        // Call API User Function Get User By ID to find what level user is { 1,2,3 }
        apiUser.GetUserById(userID).then((response)=>{
           //Get First data from data api
           const firstData = response[0];
           //Check if data from api is not match with currect localstorage state 
           if(firstData["id"] != userID ){
              this.$vs.notify({
              title: 'Error',
              text: 'Failed to Match API',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
             return;
            }
           // Get curret  level-type user
           const lvl = firstData["level"];
           switch (lvl) {
             // Level Customer
             case 1:
               this.$router.push({ name : 'customer-browse'}).catch(() => {})
               break;
             case 2:
               this.$router.push({ name : 'seller-browse'}).catch(() => {})
               break;
             case 3:
               this.$router.push({ name : 'admin-browse'}).catch(() => {})
               break; 
             default:
                this.$vs.notify({
                title: 'Failed To Redirect',
                text: 'User Level-Type Not Found',
                iconPack: 'feather',
                icon: 'icon-alert-circle',
                color: 'warning'
              })
           }
        }).catch((error) =>{
          this.$vs.notify({
            title: 'Error',
            text: 'Failed To Redirect',
            iconPack: 'feather',
            icon: 'icon-alert-circle',
            color: 'danger'
          });
          alert(error);
        });
    },
    checkLogin () {
      const userID = localStorage.getItem('userId');
      if(userID.length > 0){
        this.$vs.notify({
          title: 'Login Attempt',
          text: 'You are already logged in!, You will be redirect in 5 seconds',
          iconPack: 'feather',
          icon: 'icon-alert-circle',
          color: 'warning'
        })
        setTimeout(this.redirectUser, 5000);
      }
    },
    loginJWT () {
      apiUser
        .Login(this.email,this.password)
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
          }
          this.redirectUser()
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
        this.$vs.loading.close()
    },
    registerUser () {
      this.$router.push('/register').catch(() => {})
    }
  },
  created(){
    this.checkLogin();
  }
}
</script>

