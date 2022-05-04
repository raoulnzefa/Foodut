/*=========================================================================================
  File Name: sidebarItems.js
  Description: Sidebar Items list. Add / Remove menu items from here.
  Strucutre:
          url     => router path
          name    => name to display in sidebar
          slug    => router path name
          icon    => Feather Icon component/icon name
          tag     => text to display on badge
          tagColor  => class to apply on badge element
          i18n    => Internationalization
          submenu   => submenu of current item (current item will become dropdown )
                NOTE: Submenu don't have any icon(you can add icon if u want to display)
          isDisabled  => disable sidebar item/group
  ----------------------------------------------------------------------------------------
  Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
  Author: Pixinvent
  Author URL: http://www.themeforest.net/user/pixinvent
==========================================================================================*/


export default [
    {
      url: null,
      name: 'Authentication',
      icon: 'HomeIcon',
      i18n: 'Authentication',
      submenu: [
        {
          url: '/',
          name: 'Login',
          slug: 'login-page',
          i18n: 'Login'
        },
        {
          url: '/register',
          name: 'Register',
          slug: 'register-page',
          i18n: 'Register'
        }
      ]
    },
    {
      header: 'Apps',
      icon: 'PackageIcon',
      i18n: 'AppsPages',
      items: [
        {
          url: null,
          name: 'Store',
          icon: 'ShoppingCartIcon',
          i18n: 'Store',
          submenu: [
            {
              url: '/customer/browse',
              name: 'Browse',
              slug: 'customer-browse',
              i18n: 'Browse'
            },
            {
              url: '/customer/product/',
              name: 'Product Details',
              slug: 'customer-product-view',
              i18n: 'ProductDetails'
            },
            {
              url: '/customer/cart',
              name: 'Cart',
              slug: 'customer-cart',
              i18n: 'Cart'
            },
            {
              url: '/customer/cart/checkout',
              name: 'Checkout',
              slug: 'customer-checkout',
              i18n: 'Checkout'
            },
            {
              url: '/customer/transaction/history',
              name: 'History',
              slug: 'store-history',
              i18n: 'History'
            },
            {
              url: '/customer/store',
              name: 'Store',
              slug: 'customer-store',
              i18n: 'Store'
            }
          ]
        }
      ]
    }
  ]
  
  