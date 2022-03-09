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
  ----------------------------------------------------------------------------------------
  Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
  Author: Pixinvent
  Author URL: http://www.themeforest.net/user/pixinvent
==========================================================================================*/


import Vue from 'vue'
import Router from 'vue-router'
import auth from '@/auth/authService'

import firebase from 'firebase/app'
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
          name: 'landing-page',
          component: () => import('@/views/pages/landing/LandingPage.vue'),
          meta: {
            rule: 'editor'
          }
        }
      ]
    },
    {
      path: '',
      component: () => import('@/layouts/full-page/FullPage.vue'),
      children: [
        // =============================================================================
        // Dashboard Routes
        // =============================================================================
        {
          path: '/callback',
          name: 'auth-callback',
          component: () => import('@/views/Callback.vue'),
          meta: {
            rule: 'editor'
          }
        },
        {
          path: '/pages/login',
          name: 'page-login',
          component: () => import('@/views/pages/login/Login.vue'),
          meta: {
            rule: 'editor'
          }
        },
        {
          path: '/pages/register',
          name: 'page-register',
          component: () => import('@/views/pages/register/Register.vue'),
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
          component: () => import('./views/apps/eCommerce/ECommerceShop.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Store'},
              { title: 'Browse', active: true }
            ],
            pageTitle: 'Browse',
            rule: 'editor'
          }
        },
        {
          path: '/apps/eCommerce/item/',
          redirect: '/apps/eCommerce/item/5546604'
        },
        {
          path: '/apps/eCommerce/item/:item_id',
          name: 'ecommerce-item-detail-view',
          component: () => import('./views/apps/eCommerce/ECommerceItemDetailView.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'eCommerce'},
              { title: 'Shop', url: {name: 'ecommerce-shop'} },
              { title: 'Item Details', active: true }
            ],
            parent: 'ecommerce-item-detail-view',
            pageTitle: 'Item Details',
            rule: 'editor'
          }
        },
        {
          path: '/apps/store/cart',
          name: 'store-cart',
          component: () => import('./views/apps/eCommerce/ECommerceWishList.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Store', url:'/apps/store/shop'},
              { title: 'Cart', active: true }
            ],
            pageTitle: 'Cart',
            rule: 'editor'
          }
        },
        {
          path: '/apps/store/checkout',
          name: 'store-checkout',
          component: () => import('./views/apps/eCommerce/ECommerceCheckout.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Store', url:'/apps/store/shop'},
              { title: 'Checkout', active: true }
            ],
            pageTitle: 'Checkout',
            rule: 'editor'
          }
        },
        {
          path: '/apps/store/history',
          name: 'store-history',
          component: () => import('./views/apps/eCommerce/ECommerceHistory'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Store', url:'/apps/eCommerce/shop'},
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
              { title: 'Home', url: '/' },
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
              { title: 'Home', url: '/' },
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
              { title: 'Home', url: '/' },
              { title: 'Store'},
              { title: 'View', active: true }
            ],
            pageTitle: 'View',
            rule: 'editor'
          }
        },
        {
          path: '/advance/user/user-list',
          name: 'advance-user-list',
          component: () => import('@/views/advance/user/ListUser.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'User' },
              { title: 'List', active: true }
            ],
            pageTitle: 'User List',
            rule: 'editor'
          }
        },
        {
          path: '/advance/store/store-list',
          name: 'advance-store-list',
          component: () => import('@/views/advance/user/ListStore.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Store' },
              { title: 'List', active: true }
            ],
            pageTitle: 'Store List',
            rule: 'editor'
          }
        },

        // =============================================================================
        // COMPONENT ROUTES
        // =============================================================================
        {
          path: '/components/alert',
          name: 'component-alert',
          component: () => import('@/views/components/vuesax/alert/Alert.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Alert', active: true }
            ],
            pageTitle: 'Alert',
            rule: 'editor'
          }
        },
        {
          path: '/components/avatar',
          name: 'component-avatar',
          component: () => import('@/views/components/vuesax/avatar/Avatar.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Avatar', active: true }
            ],
            pageTitle: 'Avatar',
            rule: 'editor'
          }
        },
        {
          path: '/components/breadcrumb',
          name: 'component-breadcrumb',
          component: () => import('@/views/components/vuesax/breadcrumb/Breadcrumb.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Breadcrumb', active: true }
            ],
            pageTitle: 'Breadcrumb',
            rule: 'editor'
          }
        },
        {
          path: '/components/button',
          name: 'component-button',
          component: () => import('@/views/components/vuesax/button/Button.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Button', active: true }
            ],
            pageTitle: 'Button',
            rule: 'editor'
          }
        },
        {
          path: '/components/button-group',
          name: 'component-button-group',
          component: () => import('@/views/components/vuesax/button-group/ButtonGroup.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Button Group', active: true }
            ],
            pageTitle: 'Button Group',
            rule: 'editor'
          }
        },
        {
          path: '/components/chip',
          name: 'component-chip',
          component: () => import('@/views/components/vuesax/chip/Chip.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Chip', active: true }
            ],
            pageTitle: 'Chip',
            rule: 'editor'
          }
        },
        {
          path: '/components/collapse',
          name: 'component-collapse',
          component: () => import('@/views/components/vuesax/collapse/Collapse.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Collapse', active: true }
            ],
            pageTitle: 'Collapse',
            rule: 'editor'
          }
        },
        {
          path: '/components/dialogs',
          name: 'component-dialog',
          component: () => import('@/views/components/vuesax/dialogs/Dialogs.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Dialogs', active: true }
            ],
            pageTitle: 'Dialogs',
            rule: 'editor'
          }
        },
        {
          path: '/components/divider',
          name: 'component-divider',
          component: () => import('@/views/components/vuesax/divider/Divider.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Divider', active: true }
            ],
            pageTitle: 'Divider',
            rule: 'editor'
          }
        },
        {
          path: '/components/dropdown',
          name: 'component-drop-down',
          component: () => import('@/views/components/vuesax/dropdown/Dropdown.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Dropdown', active: true }
            ],
            pageTitle: 'Dropdown',
            rule: 'editor'
          }
        },
        {
          path: '/components/list',
          name: 'component-list',
          component: () => import('@/views/components/vuesax/list/List.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'List', active: true }
            ],
            pageTitle: 'List',
            rule: 'editor'
          }
        },
        {
          path: '/components/loading',
          name: 'component-loading',
          component: () => import('@/views/components/vuesax/loading/Loading.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Loading', active: true }
            ],
            pageTitle: 'Loading',
            rule: 'editor'
          }
        },
        {
          path: '/components/navbar',
          name: 'component-navbar',
          component: () => import('@/views/components/vuesax/navbar/Navbar.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Navbar', active: true }
            ],
            pageTitle: 'Navbar',
            rule: 'editor'
          }
        },
        {
          path: '/components/notifications',
          name: 'component-notifications',
          component: () => import('@/views/components/vuesax/notifications/Notifications.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Notifications', active: true }
            ],
            pageTitle: 'Notifications',
            rule: 'editor'
          }
        },
        {
          path: '/components/pagination',
          name: 'component-pagination',
          component: () => import('@/views/components/vuesax/pagination/Pagination.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Pagination', active: true }
            ],
            pageTitle: 'Pagination',
            rule: 'editor'
          }
        },
        {
          path: '/components/popup',
          name: 'component-popup',
          component: () => import('@/views/components/vuesax/popup/Popup.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Popup', active: true }
            ],
            pageTitle: 'Popup',
            rule: 'editor'
          }
        },
        {
          path: '/components/progress',
          name: 'component-progress',
          component: () => import('@/views/components/vuesax/progress/Progress.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Progress', active: true }
            ],
            pageTitle: 'Progress',
            rule: 'editor'
          }
        },
        {
          path: '/components/sidebar',
          name: 'component-sidebar',
          component: () => import('@/views/components/vuesax/sidebar/Sidebar.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Sidebar', active: true }
            ],
            pageTitle: 'Sidebar',
            rule: 'editor'
          }
        },
        {
          path: '/components/slider',
          name: 'component-slider',
          component: () => import('@/views/components/vuesax/slider/Slider.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Slider', active: true }
            ],
            pageTitle: 'Slider',
            rule: 'editor'
          }
        },
        {
          path: '/components/tabs',
          name: 'component-tabs',
          component: () => import('@/views/components/vuesax/tabs/Tabs.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Tabs', active: true }
            ],
            pageTitle: 'Tabs',
            rule: 'editor'
          }
        },
        {
          path: '/components/tooltip',
          name: 'component-tooltip',
          component: () => import('@/views/components/vuesax/tooltip/Tooltip.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Tooltip', active: true }
            ],
            pageTitle: 'Tooltip',
            rule: 'editor'
          }
        },
        {
          path: '/components/upload',
          name: 'component-upload',
          component: () => import('@/views/components/vuesax/upload/Upload.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Components' },
              { title: 'Upload', active: true }
            ],
            pageTitle: 'Upload',
            rule: 'editor'
          }
        },


        // =============================================================================
        // EXTENSIONS
        // =============================================================================
        {
          path: '/extensions/select',
          name: 'extra-component-select',
          component: () => import('@/views/components/extra-components/select/Select.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Extra Components' },
              { title: 'Select', active: true }
            ],
            pageTitle: 'Select',
            rule: 'editor'
          }
        },
        {
          path: '/extensions/quill-editor',
          name: 'extra-component-quill-editor',
          component: () => import('@/views/components/extra-components/quill-editor/QuillEditor.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Extra Components' },
              { title: 'Quill Editor', active: true }
            ],
            pageTitle: 'Quill Editor',
            rule: 'editor'
          }
        },
        {
          path: '/extensions/drag-and-drop',
          name: 'extra-component-drag-and-drop',
          component: () => import('@/views/components/extra-components/drag-and-drop/DragAndDrop.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Extra Components' },
              { title: 'Drag & Drop', active: true }
            ],
            pageTitle: 'Drag & Drop',
            rule: 'editor'
          }
        },
        {
          path: '/extensions/datepicker',
          name: 'extra-component-datepicker',
          component: () => import('@/views/components/extra-components/datepicker/Datepicker.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Extra Components' },
              { title: 'Datepicker', active: true }
            ],
            pageTitle: 'Datepicker',
            rule: 'editor'
          }
        },
        {
          path: '/extensions/datetime-picker',
          name: 'extra-component-datetime-picker',
          component: () => import('@/views/components/extra-components/datetime-picker/DatetimePicker.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Extra Components' },
              { title: 'Datetime Picker', active: true }
            ],
            pageTitle: 'Datetime Picker',
            rule: 'editor'
          }
        },
        {
          path: '/extensions/access-control',
          name: 'extra-component-access-control',
          component: () => import('@/views/components/extra-components/access-control/AccessControl.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Extensions' },
              { title: 'Access Control', active: true }
            ],
            pageTitle: 'Access Control',
            rule: 'editor'
          }
        },
        {
          path: '/extensions/i18n',
          name: 'extra-component-i18n',
          component: () => import('@/views/components/extra-components/I18n.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Extensions' },
              { title: 'I18n', active: true }
            ],
            pageTitle: 'I18n',
            rule: 'editor'
          }
        },
        {
          path: '/extensions/carousel',
          name: 'extra-component-carousel',
          component: () => import('@/views/components/extra-components/carousel/Carousel.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Extensions' },
              { title: 'Carousel', active: true }
            ],
            pageTitle: 'Carousel',
            rule: 'editor'
          }
        },
        {
          path: '/extensions/clipboard',
          name: 'extra-component-clipboard',
          component: () => import('@/views/components/extra-components/clipboard/Clipboard.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Extensions' },
              { title: 'Clipboard', active: true }
            ],
            pageTitle: 'Clipboard',
            rule: 'editor'
          }
        },
        {
          path: '/extensions/context-menu',
          name: 'extra-component-context-menu',
          component: () => import('@/views/components/extra-components/context-menu/ContextMenu.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Extensions' },
              { title: 'Context Menu', active: true }
            ],
            pageTitle: 'Context Menu',
            rule: 'editor'
          }
        },
        {
          path: '/extensions/star-ratings',
          name: 'extra-component-star-ratings',
          component: () => import('@/views/components/extra-components/star-ratings/StarRatings.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Extensions' },
              { title: 'Star Ratings', active: true }
            ],
            pageTitle: 'Star Ratings',
            rule: 'editor'
          }
        },
        {
          path: '/extensions/autocomplete',
          name: 'extra-component-autocomplete',
          component: () => import('@/views/components/extra-components/autocomplete/Autocomplete.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Extensions' },
              { title: 'Autocomplete', active: true }
            ],
            pageTitle: 'Autocomplete',
            rule: 'editor'
          }
        },
        {
          path: '/extensions/tree',
          name: 'extra-component-tree',
          component: () => import('@/views/components/extra-components/tree/Tree.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Extensions' },
              { title: 'Tree', active: true }
            ],
            pageTitle: 'Tree',
            rule: 'editor'
          }
        },
        {
          path: '/import-export/import',
          name: 'import-excel',
          component: () => import('@/views/components/extra-components/import-export/Import.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Extensions' },
              { title: 'Import/Export' },
              { title: 'Import', active: true }
            ],
            pageTitle: 'Import Excel',
            rule: 'editor'
          }
        },
        {
          path: '/import-export/export',
          name: 'export-excel',
          component: () => import('@/views/components/extra-components/import-export/Export.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Extensions' },
              { title: 'Import/Export' },
              { title: 'Export', active: true }
            ],
            pageTitle: 'Export Excel',
            rule: 'editor'
          }
        },
        {
          path: '/import-export/export-selected',
          name: 'export-excel-selected',
          component: () => import('@/views/components/extra-components/import-export/ExportSelected.vue'),
          meta: {
            breadcrumb: [
              { title: 'Home', url: '/' },
              { title: 'Extensions' },
              { title: 'Import/Export' },
              { title: 'Export Selected', active: true }
            ],
            pageTitle: 'Export Excel',
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

router.beforeEach((to, from, next) => {
  firebase.auth().onAuthStateChanged(() => {

    // get firebase current user
    const firebaseCurrentUser = firebase.auth().currentUser

    // if (
    //     to.path === "/pages/login" ||
    //     to.path === "/pages/forgot-password" ||
    //     to.path === "/pages/error-404" ||
    //     to.path === "/pages/error-500" ||
    //     to.path === "/pages/register" ||
    //     to.path === "/callback" ||
    //     to.path === "/pages/comingsoon" ||
    //     (auth.isAuthenticated() || firebaseCurrentUser)
    // ) {
    //     return next();
    // }

    // If auth required, check login. If login fails redirect to login page
    if (to.meta.authRequired) {
      if (!(auth.isAuthenticated() || firebaseCurrentUser)) {
        router.push({ path: '/pages/login', query: { to: to.path } })
      }
    }

    return next()
    // Specify the current path as the customState parameter, meaning it
    // will be returned to the application after auth
    // auth.login({ target: to.path });

  })

})

export default router
