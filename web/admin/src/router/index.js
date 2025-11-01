import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'
// eslint-disable-next-line no-unused-vars
import { title } from '@/settings'

/* Router Modules */
// import componentsRouter from './modules/components'
// import chartsRouter from './modules/charts'
// import tableRouter from './modules/table'
// import nestedRouter from './modules/nested'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
 roles: ['admin','editor']    control the page roles (you can set multiple roles)
 title: 'title'               the name show in sidebar and breadcrumb (recommend set)
 icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
 noCache: true                if set true, the page will no be cached(default is false)
 affix: true                  if set true, the tag will affix in the tags-view
 breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
 activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
 }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/redirect',
    component: Layout,
    hidden: true,
    children: [
      {
        path: '/redirect/:path(.*)',
        component: () => import('@/views/redirect/index')
      }
    ]
  },
  {
    path: '/download',
    component: () => import('@/views/download/index'),
    meta: { noAuth: true },
    hidden: true
  },
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  {
    path: '/auth-redirect',
    component: () => import('@/views/login/auth-redirect'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@/views/error-page/404'),
    hidden: true
  },
  {
    path: '/401',
    component: () => import('@/views/error-page/401'),
    hidden: true
  },
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        component: () => import('@/views/dashboard/index'),
        name: 'Dashboard',
        meta: { title: '控制台', icon: 'dashboard', affix: true }
      }
    ]
  },
  {
    path: '/profile',
    component: Layout,
    redirect: '/profile/index',
    hidden: true,
    children: [
      {
        path: 'index',
        component: () => import('@/views/profile/index'),
        name: 'Profile',
        meta: { title: '个人中心', icon: 'user', noCache: true }
      }
    ]
  }
]

/**
 * asyncRoutes
 * the routes that need to be dynamically loaded based on user roles
 */
export const asyncRoutes = [
  // {
  //   path: '/permission',
  //   component: Layout,
  //   redirect: '/permission/page',
  //   alwaysShow: true, // will always show the root menu
  //   name: 'Permission',
  //   meta: {
  //     title: 'Permission',
  //     icon: 'lock',
  //     roles: ['admin', 'editor'] // you can set roles in root nav
  //   },
  //   children: [
  //     {
  //       path: 'page',
  //       component: () => import('@/views/permission/page'),
  //       name: 'PagePermission',
  //       meta: {
  //         title: 'Page Permission',
  //         roles: ['admin'] // or you can only set roles in sub nav
  //       }
  //     },
  //     {
  //       path: 'directive',
  //       component: () => import('@/views/permission/directive'),
  //       name: 'DirectivePermission',
  //       meta: {
  //         title: 'Directive Permission'
  //         // if do not set roles, means: this page does not require permission
  //       }
  //     },
  //     {
  //       path: 'role',
  //       component: () => import('@/views/permission/role'),
  //       name: 'RolePermission',
  //       meta: {
  //         title: 'Role Permission',
  //         roles: ['admin']
  //       }
  //     }
  //   ]
  // },
  {
    path: '/users',
    component: Layout,
    name: '任务平台',
    meta: { title: '任务平台', icon: 'user', roles: ['admin', 'editor'] },
    children: [
      {
        path: 'users',
        component: () => import('@/views/users/index'),
        name: 'Users',
        meta: { title: '会员管理', icon: 'user', noCache: true, roles: ['admin', 'editor'] }
      }
    ]
  },
  {
    path: '/agent',
    component: Layout,
    name: '代理管理',
    meta: { title: '代理管理', icon: 'list', roles: ['admin'] },
    children: [
      {
        path: 'list',
        component: () => import('@/views/agent/index'),
        name: 'School',
        meta: { title: '代理列表', icon: 'list', noCache: true }
      },
      {
        path: 'order',
        component: () => import('@/views/order/index'),
        name: 'Order',
        meta: { title: '订单列表', icon: 'shopping', noCache: true, roles: ['admin'] }
      }
      // {
      //   path: 'refund',
      //   component: () => import('@/views/order/refund'),
      //   name: 'Refund',
      //   meta: { title: '退款管理', icon: 'money', noCache: true }
      // }
    ]
  },

  {
    path: '/admin',
    component: Layout,
    name: '管理员管理',
    icon: 'lock',
    meta: { title: '管理员管理', icon: 'lock', roles: ['admin'] },
    children: [
      {
        path: 'permission',
        component: () => import('@/views/admin/permission'),
        name: 'Permission',
        meta: { title: '权限管理', icon: 'tree', noCache: true }
      },
      {
        path: 'role',
        component: () => import('@/views/admin/role'),
        name: 'Role',
        meta: { title: '角色管理', icon: 'chart', noCache: true }
      },
      {
        path: 'account',
        component: () => import('@/views/admin/index'),
        name: 'Admin',
        meta: { title: '账号管理', icon: 'user', noCache: true }
      }
    ]
  },
  {
    path: '/system',
    component: Layout,
    name: '系统管理',
    meta: { title: '系统管理', icon: 'el-icon-setting', roles: ['admin'] },
    children: [
      {
        path: 'city',
        component: () => import('@/views/system/area'),
        name: 'City',
        meta: { title: '开放地区', icon: 'el-icon-setting', noCache: true }
      },
      {
        path: 'pay',
        component: () => import('@/views/system/pay'),
        name: 'Pay',
        meta: { title: '支付配置', icon: 'el-icon-setting', noCache: true }
      },
      {
        path: 'index',
        component: () => import('@/views/system/index'),
        name: 'System',
        meta: { title: '系统配置', icon: 'el-icon-setting', noCache: true }
      }
    ]
  },
  // {
  //   path: '/icon',
  //   component: Layout,
  //   children: [
  //     {
  //       path: 'index',
  //       component: () => import('@/views/icons/index'),
  //       name: 'Icons',
  //       meta: { title: 'Icons', icon: 'icon', noCache: true }
  //     }
  //   ]
  // },

  /** when your routing map is too long, you can split it into small modules **/
  // tableRouter,
  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
