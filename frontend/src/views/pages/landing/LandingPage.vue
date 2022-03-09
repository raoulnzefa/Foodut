<template>
  <div class="h-screen flex w-full bg-img vx-row no-gutter items-center justify-center" id="landing-page">
    <div class="vx-col sm:m-0 m-4">
      <vx-card>
        <div slot="no-body" class="full-page-bg-color">

          <div class="vx-row no-gutter justify-center items-center">

            <div class="vx-col hidden lg:block lg:w-1/2">
              <img src="@/assets/images/logo/foodut_shadow.png" width="100%" alt="foodut logo" class="mx-auto">
            </div>

            <div class="vx-col sm:w-full md:w-full lg:w-1/2 d-theme-dark-bg">
              <div class="px-8 pt-8 login-tabs-container">

                <div class="vx-card__title mb-4">
                  <h4 class="mb-4">Foodut</h4>
                  <p>Welcome to the most complete vast variant home made snack website ever.</p>
                </div>

                <vs-tabs>
                  <vs-tab label="LOGIN">
                    <login></login>
                  </vs-tab>

                  <vs-tab label="REGISTER">
                    <register></register>
                  </vs-tab>

                  <vs-tab label="GUEST">
                    <div class="vx-card__title mb-4">
                      <h4 class="mb-4">Just passing by?</h4>
                      <p>We really welcome you. As a guest, you can browse products and shops. You just can't make any purchases. 
                        Make sure to register if you want to buy the product.</p>
                    </div>
                    <!-- Redirect to home page products -->
                    <vs-button type="gradient" to="/pages/login">Start Wandering</vs-button>
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
import Login from './Login.vue'
import Register from './Register.vue'

export default {
  components: {
    Login,
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
