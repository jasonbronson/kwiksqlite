import { createRouter, createWebHashHistory } from 'vue-router'
import DBConnect from '../views/DBConnect'

const routes = [
  {
    // Document title tag
    // We combine it with defaultDocumentTitle set in `src/main.js` on router.afterEach hook
    meta: {
      title: 'Database'
    },
    path: '/',
    name: 'DBConnect',
    component: DBConnect
  },
  {
    meta: {
      title: 'Tables'
    },
    path: '/tables',
    name: 'tables',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "tables" */ '../views/Tables')
  },
    // {
    //   meta: {
    //     title: 'Database Connect'
    //   },
    //   path: '/dbconnect',
    //   name: 'dbconnect',
    //   component: () => import(/* webpackChunkName: "dbconnect" */ '../views/DBConnect')
    // },
  // {
  //   meta: {
  //     title: 'Profile'
  //   },
  //   path: '/profile',
  //   name: 'profile',
  //   component: () => import(/* webpackChunkName: "profile" */ '../views/Profile')
  // },
  {
    meta: {
      title: 'Login'
    },
    path: '/login',
    name: 'login',
    component: () => import(/* webpackChunkName: "login" */ '../views/Login')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
  scrollBehavior (to, from, savedPosition) {
    return savedPosition || { x: 0, y: 0 }
  }
})

export default router
