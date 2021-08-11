import { createRouter, createWebHashHistory } from "vue-router";
import DBConnect from "../views/DBConnect";

const routes = [
  {
    // Document title tag
    // We combine it with defaultDocumentTitle set in `src/main.js` on router.afterEach hook
    meta: {
      title: "Database",
    },
    path: "/",
    name: "DBConnect",
    component: DBConnect,
  },
  {
    meta: {
      title: "Table",
    },
    path: "/table/:tablename",
    name: "table",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "table" */ "../views/Table"),
  },
  {
    meta: {
      title: "Tables",
    },
    path: "/tables",
    name: "tables",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "table" */ "../views/Table"),
  },
  {
    meta: {
      title: "Query",
    },
    path: "/query",
    name: "Query",
    component: () =>
      import(/* webpackChunkName: "dbconnect" */ "../views/DBConnect"),
  },
  {
    meta: {
      title: "Create",
    },
    path: "/create",
    name: "Create",
    component: () =>
      import(/* webpackChunkName: "dbconnect" */ "../views/DBConnect"),
  },
  {
    meta: {
      title: "Login",
    },
    path: "/login",
    name: "login",
    component: () => import(/* webpackChunkName: "login" */ "../views/Login"),
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    return savedPosition || { x: 0, y: 0 };
  },
});

export default router;
