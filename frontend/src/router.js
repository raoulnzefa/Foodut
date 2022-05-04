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
    // =============================================================================
    // Global Authentication Routes
    // =============================================================================
    {
      path: '',
      component: () => import('@/layouts/full-page/FullPage.vue'),
      children: [
        {
          path: '/',
          name: 'login-page',
          component: () => import('@/views/authentication/login/Login.vue'),
          meta: {
            rule: 'editor'
          }
        },
        {
          path: '/register',
          name: 'register-page',
          component: () => import('@/views/authentication/register/Register.vue'),
          meta: {
            rule: 'editor'
          }
        },
        {
          path: '/forgot-password',
          name: 'page-forgot-password',
          component: () => import('@/views/authentication/ForgotPassword.vue'),
          meta: { 
            rule: 'editor'
          }
        }
      ]
    },
    // =============================================================================
    // Customer
    // =============================================================================
    {
      path: '',
      component: () => import('./views/customer/Main.vue'),
      children: [
        {
          path: '/customer/browse',
          name: 'customer-browse',
          component: () => import('./views/customer/browse/Browse.vue'),
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
          path: '/customer/product/:product_id',
          name: 'customer-product-detail',
          component: () => import('./views/customer/ProductDetail.vue'),
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
          path: '/customer/cart',
          name: 'customer-cart',
          component: () => import('./views/customer/cart/Cart.vue'),
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
          path: '/customer/cart/checkout',
          name: 'customer-checkout',
          component: () => import('./views/customer/checkout/Checkout.vue'),
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
          path: '/customer/transaction/history',
          name: 'customer-transaction-history',
          component: () => import('./views/customer/history/History.vue'),
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
        {
          path: '/customer/store',
          name: 'customer-store',
          component: () => import('./views/customer/store/Store.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home'},
              { title: 'Store'},
              { title: 'View', active: true }
            ],
            pageTitle: 'View Store',
            rule: 'editor'
          }
        },
        // =============================================================================
        // Profile
        // =============================================================================
        {
          path: '/customer/profile',
          name: 'customer-profile',
          component: () => import('@/views/customer/profile/EditProfile.vue'),
          meta: {
            rule: 'editor'
          }
        }
      ]
    },
    // =============================================================================
    // Seller Application Routes
    // =============================================================================
    {
      path: '',
      component: () => import('./views/seller/Main.vue'),
      children: [
        {
          path: '/seller/browse',
          name: 'seller-browse',
          component: () => import('./views/seller/browse/Browse.vue'),
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
          path: '/seller/product/add',
          name: 'seller-add-product',
          component: () => import('./views/seller/product/AddProduct.vue'),
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
          path: '/seller/product/edit',
          name: 'seller-edit-product',
          component: () => import('./views/seller/product/EditProduct.vue'),
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
          path: '/seller/product',
          name: 'seller-product',
          component: () => import('./views/seller/store/Store.vue'),
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
          path: '/seller/transaction/order',
          name: 'seller-transaction-order',
          component: () => import('./views/seller/order/ViewTransactionOrder.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home'},
              { title: 'Store'},
              { title: 'Orders'},
              { title: 'View', active: true }
            ],
            pageTitle: 'ViewOrders',
            rule: 'editor'
          }
        },
        // =============================================================================
        // Profile Routes
        // =============================================================================
        {
          path: '/seller/profile',
          name: 'seller-profile',
          component: () => import('@/views/seller/profile/EditProfile.vue'),
          meta: {
            rule: 'editor'
          }
        }
      ]
    },
    // =============================================================================
    // Admin  Application Routes
    // =============================================================================
    {
      path: '',
      component: () => import('./views/admin/Main.vue'),
      children: [
        {
          path: '/admin/browse',
          name: 'admin-browse',
          component: () => import('./views/admin/browse/Browse.vue'),
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
          path: '/admin/product/',
          redirect: '/admin/product/5546604'
        },
        {
          path: '/admin/product/:product_id',
          name: 'admin-product-detail',
          component: () => import('./views/admin/ProductDetail.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home'},
              { title: 'Store'},
              { title: 'Shop', url: {name: 'store-product'} },
              { title: 'Product Details', active: true }
            ],
            parent: 'admin-product-detail',
            pageTitle: 'Product Details',
            rule: 'editor'
          }
        },
        {
          path: '/admin/user/:userId',
          name: 'admin-user',
          component: () => import('@/views/admin/user/ListUser.vue'),
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
          path: '/admin/store',
          name: 'admin-store',
          component: () => import('@/views/admin/store/ListStore.vue'),
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
        // Test Api
        // =============================================================================
        {
          path: '/admin/test',
          name: 'admin-test-api',
          component: () => import('@/views/TestApi.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home'},
              { title: 'Test Api', active: true }
            ],
            pageTitle: 'TestApi',
            rule: 'editor'
          }
        },
        // =============================================================================
        // Profile Routes
        // =============================================================================
        {
          path: '/admin/profile',
          name: 'admin-profile',
          component: () => import('@/views/admin/profile/EditProfile.vue'),
          meta: {
            rule: 'editor'
          }
        }
      ]
    },
    // =============================================================================
    // Guest Application Routes
    // =============================================================================
    {
      path: '',
      component: () => import('./views/guest/Main.vue'),
      children: [
        {
          path: '/guest/browse',
          name: 'guest-browse',
          component: () => import('./views/guest/browse/Browse.vue'),
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
          path: '/guest/product/:product_id',
          name: 'guest-product-detail',
          component: () => import('./views/guest/ProductDetail.vue'),
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
          path: '/guest/store',
          name: 'guest-store',
          component: () => import('./views/guest/store/Store.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home'},
              { title: 'Store'},
              { title: 'View', active: true }
            ],
            pageTitle: 'ViewStore',
            rule: 'editor'
          }
        }
      ]
    },
    {
      path: '/pages/error-404',
      component: () => import('@/views/pages/Error404.vue'),
      meta: {
        rule: 'editor'
      }
    },
    // Redirect to 404 page, if no match found
    {
      path: '*',
      redirect: '/pages/error-404',
    }
  ]}
)

router.afterEach(() => {
  // Remove initial loading
  const appLoading = document.getElementById('loading-bg')
  if (appLoading) {
    appLoading.style.display = 'none'
  }
})

export default router
