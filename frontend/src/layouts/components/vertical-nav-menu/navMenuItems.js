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
    name: 'Dashboard',
    icon: 'HomeIcon',
    i18n: 'Dashboard',
    submenu: [
      {
        url: '/',
        name: 'Login',
        slug: 'pages-login',
        i18n: 'Login',
        target: '_blank'
      },
      {
        url: '/register',
        name: 'Register',
        slug: 'pages-register',
        i18n: 'Register',
        target: '_blank'
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
            url: '/apps/store/browse',
            name: 'Browse',
            slug: 'store-browse',
            i18n: 'Browse'
          },
          {
            url: '/apps/store/item/',
            name: 'Item Details',
            slug: 'store-item-detail-view',
            i18n: 'ItemDetails'
          },
          {
            url: '/apps/store/cart',
            name: 'Cart',
            slug: 'store-cart',
            i18n: 'Cart'
          },
          {
            url: '/apps/store/checkout',
            name: 'Checkout',
            slug: 'store-checkout',
            i18n: 'Checkout'
          },
          {
            url: '/apps/store/history',
            name: 'History',
            slug: 'store-history',
            i18n: 'History'
          }
        ]
      }
    ]
  },
  {
    header: 'Advance',
    icon: 'PackageIcon',
    i18n: 'Advance',
    items: [
      {
        url: null,
        name: 'Store',
        icon: 'ShoppingCartIcon',
        i18n: 'Store',
        submenu: [
          {
            url: '/advance/store/product/add',
            name: 'Add Product',
            slug: 'store-add-product',
            i18n: 'AddProduct'
          },
          {
            url: '/advance/store/product/edit',
            name: 'Edit Product',
            slug: 'store-edit-product',
            i18n: 'EditProduct'
          },
          {
            url: '/advance/store/view',
            name: 'View',
            slug: 'store-view',
            i18n: 'ViewStore'
          }
        ]
      },
      {
        url: null,
        name: 'User',
        icon: 'UserIcon',
        i18n: 'User',
        submenu: [
          {
            url: '/advance/user/user-list/268',
            name: 'List User',
            slug: 'advance-user-list',
            i18n: 'ListUser'
          },
          {
            url: '/advance/store/store-list',
            name: 'List Store',
            slug: 'advance-store-list',
            i18n: 'ListStore'
          }
        ]
      }
    ]
  }
]

