/*=========================================================================================
  File Name: router.js
  Description: Routes for vue-router. Lazy loading is enabled.
  Object Strucutre:
    path => router path
    name => router name
    component(lazy loading) => component to load
    meta : {
      rule => which user can have access (ACL)
      breadcrumb => Add breadcrumb to specific page
      pageTitle => Display title besides breadcrumb
    }
==========================================================================================*/


import Vue from 'vue'
import Router from 'vue-router'
//import auth from '@/auth/authService'

import 'firebase/auth'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  scrollBehavior () {
    return { x: 0, y: 0 }
  },
  routes: [
    {
      path: '',
      component: () => import('@/layouts/full-page/FullPage.vue'),
      children: [
        // =============================================================================
        // Dashboard Routes
        // =============================================================================
        {
          path: '/',
          name: 'login-page',
          component: () => import('@/views/dashboard/Login.vue'),
          meta: {
            rule: 'editor'
          }
        },
        {
          path: '/register',
          name: 'register-page',
          component: () => import('@/views/dashboard/Register.vue'),
          meta: {
            rule: 'editor'
          }
        },
        {
          path: '/forgot-password',
          name: 'page-forgot-password',
          component: () => import('@/views/dashboard/ForgotPassword.vue'),
          meta: {
            rule: 'editor'
          }
        }
      ]
    },
    {
      path: '',
      component: () => import('./layouts/main/Main.vue'),
      children: [
        // =============================================================================
        // Application Routes
        // =============================================================================
        {
          path: '/apps/store/browse',
          name: 'store-browse',
          component: () => import('./views/apps/store/Browse.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home'},
              { title: 'Store'},
              { title: 'Browse', active: true }
            ],
            pageTitle: 'Browse',
            rule: 'editor'
          }
        },
        {
          path: '/apps/store/item/',
          redirect: '/apps/eCommerce/item/5546604'
        },
        {
          path: '/apps/eCommerce/item/:item_id',
          name: 'ecommerce-item-detail-view',
          component: () => import('./views/apps/store/ItemDetailView.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home'},
              { title: 'Store'},
              { title: 'Shop', url: {name: 'store-shop'} },
              { title: 'Item Details', active: true }
            ],
            parent: 'store-item-detail-view',
            pageTitle: 'Item Details',
            rule: 'editor'
          }
        },
        {
          path: '/apps/store/cart',
          name: 'store-cart',
          component: () => import('./views/apps/store/Cart.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home'},
              { title: 'Store'},
              { title: 'Cart', active: true }
            ],
            pageTitle: 'Cart',
            rule: 'editor'
          }
        },
        {
          path: '/apps/store/checkout',
          name: 'store-checkout',
          component: () => import('./views/apps/store/Checkout.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home'},
              { title: 'Store'},
              { title: 'Checkout', active: true }
            ],
            pageTitle: 'Checkout',
            rule: 'editor'
          }
        },
        {
          path: '/apps/store/history',
          name: 'store-history',
          component: () => import('./views/apps/store/History.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home'},
              { title: 'Store'},
              { title: 'History', active: true }
            ],
            pageTitle: 'History',
            rule: 'editor'
          }
        },
        // =============================================================================
        // Advance Routes
        // =============================================================================
        {
          path: '/advance/store/product/add',
          name: 'store-add-product',
          component: () => import('./views/advance/store/AddProduct.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home'},
              { title: 'Store'},
              { title: 'Product'},
              { title: 'Add', active: true }
            ],
            pageTitle: 'Add',
            rule: 'editor'
          }
        },
        {
          path: '/advance/store/product/edit',
          name: 'store-edit-product',
          component: () => import('./views/advance/store/EditProduct.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home'},
              { title: 'Store'},
              { title: 'Product'},
              { title: 'Edit', active: true }
            ],
            pageTitle: 'Edit',
            rule: 'editor'
          }
        },
        {
          path: '/advance/store/view',
          name: 'store-view',
          component: () => import('./views/advance/store/ViewStore.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home'},
              { title: 'Store'},
              { title: 'View', active: true }
            ],
            pageTitle: 'ViewStore',
            rule: 'editor'
          }
        },
        {
          path: '/advance/store/product/view',
          name: 'store-product-view',
          component: () => import('./views/advance/store/ViewSoldProduct.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home'},
              { title: 'Store'},
              { title: 'Product'},
              { title: 'View', active: true }
            ],
            pageTitle: 'ViewProduct',
            rule: 'editor'
          }
        },
        {
          path: '/advance/user/user-list/:userId',
          name: 'advance-user-list',
          component: () => import('@/views/advance/user/ListUser.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home'},
              { title: 'User' },
              { title: 'List', active: true }
            ],
            pageTitle: 'List User',
            rule: 'editor'
          }
        },
        {
          path: '/advance/store/store-list',
          name: 'advance-store-list',
          component: () => import('@/views/advance/user/ListStore.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home'},
              { title: 'Store' },
              { title: 'List', active: true }
            ],
            pageTitle: 'List Store',
            rule: 'editor'
          }
        },
        // =============================================================================
        // Profile Routes
        // =============================================================================
        {
          path: '/profile',
          name: 'profile',
          component: () => import('@/views/profile/EditProfile.vue'),
          meta: {
            rule: 'editor'
          }
        }
      ]
    },

    // Redirect to 404 page, if no match found
    {
      path: '*',
      redirect: '/pages/error-404'
    }
  ]
})

router.afterEach(() => {
  // Remove initial loading
  const appLoading = document.getElementById('loading-bg')
  if (appLoading) {
    appLoading.style.display = 'none'
  }
})

// router.beforeEach((to, from, next) => {
//   firebase.auth().onAuthStateChanged(() => {

//     // get firebase current user
//     const firebaseCurrentUser = firebase.auth().currentUser

//     // if (
//     //     to.path === "/pages/login" ||
//     //     to.path === "/pages/forgot-password" ||
//     //     to.path === "/pages/error-404" ||
//     //     to.path === "/pages/error-500" ||
//     //     to.path === "/pages/register" ||
//     //     to.path === "/callback" ||
//     //     to.path === "/pages/comingsoon" ||
//     //     (auth.isAuthenticated() || firebaseCurrentUser)
//     // ) {
//     //     return next();
//     // }

//     // If auth required, check login. If login fails redirect to login page
//     if (to.meta.authRequired) {
//       if (!(auth.isAuthenticated() || firebaseCurrentUser)) {
//         router.push({ path: '/pages/login', query: { to: to.path } })
//       }
//     }

//     return next()
//     // Specify the current path as the customState parameter, meaning it
//     // will be returned to the application after auth
//     // auth.login({ target: to.path });

//   })

// })

export default router
