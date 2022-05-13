<!-- =========================================================================================
    File Name: Login.vue
    Description: Login Page
========================================================================================== -->

<template>
  <div class="h-screen flex w-full bg-img vx-row no-gutter items-center justify-center" id="landing-page">
    <div class="vx-col sm:m-0 m-4">
      <vx-card>
        <div slot="no-body" class="full-page-bg-color">

           <div class="vx-row no-gutter justify-center items-center">

            <div class="vx-col hidden lg:block lg:w-1/2">
              <img src="@/assets/images/pages/login.png" alt="login" class="mx-auto">
            </div>

            <div class="vx-col sm:w-full md:w-full lg:w-1/2 d-theme-dark-bg">
              <div class="px-8 pt-8 login-tabs-container">

                <div class="vx-card__title mb-4">
                  <h4 class="mb-4">Login</h4>
                  <p>Welcome back, please login to your account.</p>
                </div>

                <vs-tabs>
                  <vs-tab label="LOGIN">
                    <login></login>
                  </vs-tab>

                  <vs-tab label="GUEST">
                    
                    <h4 class="mb-4">Just passing by?</h4>
                    <p >We really welcome you. As a guest, you can </p> 
                    <p>browse products and shops. Make sure to  </p> 
                    <p>register if you want to buy the product.</p><br>
                    
                    <!-- Redirect to home page products -->
                    <vs-button type="gradient" to="../guest/browse" class="mb-4">Start Wandering</vs-button>
                  </vs-tab>
                </vs-tabs>

              </div>
            </div>
          </div>
        </div>
      </vx-card>
    </div>
  </div>
</template>


<script>
import Login from './component/JWT.vue'


export default {
  components: {
    Login
  }
}
</script>

<style lang="scss">
.login-tabs-container {
  min-height: 505px;

  .con-tab {
    padding-bottom: 14px;
  }

  .con-slot-tabs {
    margin-top: 1rem;
  }
}
</style>
