<template>
  <div class="h-screen flex w-full bg-img vx-row no-gutter items-center justify-center" id="landing-page">
    <div class="vx-col sm:m-0 m-4">
      <vx-card>
        <div slot="no-body" class="full-page-bg-color">
          <div class="vx-row no-gutter">
            <div class="vx-col hidden sm:hidden md:hidden lg:block lg:w-1/2 mx-auto self-center">
                <img src="@/assets/images/pages/register.jpg" alt="register" class="mx-auto">
            </div>
            <div class="vx-col sm:w-full md:w-full lg:w-1/2 mx-auto self-center  d-theme-dark-bg">
              <div class="px-8 pt-8 register-tabs-container">
                <div class="vx-card__title mb-4">
                    <h4 class="mb-4">Create Account</h4>
                    <p>Fill the below form to create a new account.</p>
                </div>

                <vs-tabs>
                  <vs-tab label="REGISTER">
                    <register></register>
                  </vs-tab>

                  <vs-tab label="GUEST">
                    
                    <h4 class="mb-4">Just passing by?</h4>
                    <p >We really welcome you. As a guest, you can </p> 
                    <p>browse products and shops. Make sure to  </p> 
                    <p>register if you want to buy the product.</p><br>
                    
                    <!-- Redirect to home page products -->
                    <vs-button type="gradient" to="../apps/store/browse" class="mb-4">Start Wandering</vs-button>
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
import Register from './RegisterJWT.vue'

export default {
  components: {
    Register
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
