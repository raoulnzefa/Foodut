<template>
  <div class="the-navbar__user-meta flex items-center" v-if="activeUserInfo.name">
<!-- id: '6',
        email: 'admin@gmail.com',
        name: 'admin',
        username: 'admin',
        level: '3',
        profilePhoto: '' -->
    <div class="text-right leading-tight hidden sm:block">
      <p class="font-semibold">{{ activeUserInfo.name }}</p>
      <small>{{ activeUserInfo.level }}</small>
    </div>

    <vs-dropdown vs-custom-content vs-trigger-click class="cursor-pointer">

      <div class="con-img ml-3">
        <img v-if="activeUserInfo.profilePhoto" key="onlineImg" :src="activeUserInfo.profilePhoto" alt="user-img" width="40" height="40" class="rounded-full shadow-md cursor-pointer block" />
      </div>

      <vs-dropdown-menu class="vx-navbar-dropdown">
        <ul style="min-width: 9.5rem">
          <li
            class="flex py-2 px-4 cursor-pointer hover:bg-primary hover:text-white"
            @click="$router.push('/profile').catch(() => {})">
            <feather-icon icon="SettingsIcon" svgClasses="w-4 h-4" />
            <span class="ml-2">Edit Profile</span>
          </li>

          <li
            class="flex py-2 px-4 cursor-pointer hover:bg-primary hover:text-white"
            @click="logout">
            <feather-icon icon="LogOutIcon" svgClasses="w-4 h-4" />
            <span class="ml-2">Logout</span>
          </li>
        </ul>
      </vs-dropdown-menu>
    </vs-dropdown>
  </div>
</template>

<script>
import apiUser from '../../../../api/user'

export default {
  data () {
    return {
      activeUserInfo: {
        //Dummy
        id: '0',
        email: 'test@gmail.com',
        name: 'test',
        username: 'test',
        level: '0',
        profilePhoto: require('@/assets/images/portrait/small/avatar-s-11.jpg')
      }
    }
  },
  created() {
    apiUser
      .GetUserById(localStorage.getItem('userId'))
      .then((response) => { 
        this.activeUserInfo.id = response[0].id
        this.activeUserInfo.email = response[0].email
        this.activeUserInfo.name = response[0].name
        this.activeUserInfo.username = response[0].username
        this.activeUserInfo.level = response[0].level
        // userData.profilePhoto = response[0].profilePhoto
        console.log('Response Api Get User By ID')
        console.log(response)
      })
      .catch((error) => { console.log(error) })
  },
  computed: {
    // activeUserInfo () {
      // const userId = localStorage.getItem('userId')
      // await apiUser
      //   .GetUserById(userId)
      //   .then((response) => { 
      //     userData.id = response[0].id
      //     userData.email = response[0].email
      //     userData.name = response[0].name
      //     userData.username = response[0].username
      //     userData.level = response[0].level
      //     // userData.profilePhoto = response[0].profilePhoto
      //     console.log('Response Get User By ID')
      //     console.log(response)
      //   })
      //   .catch((error) => { console.log(error) })
      // return userData
      // console.log('App Active User')
      // console.log(this.$store.state.AppActiveUser)
    //   return this.$store.state.AppActiveUser
    // }
  },
  methods: {
    logout () {
      apiUser
        .Logout()
        .then((response) => {
          if(!response){
            this.$vs.notify({
              title: 'Error',
              text: 'Failed to logout',
              iconPack: 'feather',
              icon: 'icon-alert-circle',
              color: 'danger'
            })
            return
          }
          this.$vs.loading.close()
          this.$router.push({ name : 'login-page'}).catch(() => {})
        })
        .catch((error) => {
          this.$vs.loading.close()
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
