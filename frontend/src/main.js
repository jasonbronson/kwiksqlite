import { createApp } from 'vue'

import App from './App.vue'
import router from './router'
import store from './store'
import axios from "axios";
import VueAxios from "vue-axios";
import './css/main.css'

//dev mode
let baseURL = "http://192.168.1.80:1234/api";
//build mode
//baseURL = "/api"

const client = axios.create({
  baseURL: baseURL,
});


/* Default title tag */
const defaultDocumentTitle = 'KwikSqlite Admin'

/* Collapse mobile aside menu on route change & set document title from route meta */
router.afterEach(to => {
  store.dispatch('asideMobileToggle', false)

  if (to.meta && to.meta.title) {
    document.title = `${to.meta.title} - ${defaultDocumentTitle}`
  } else {
    document.title = defaultDocumentTitle
  }
})

createApp(App).use(store).use(router).use(VueAxios, client).mount('#app')
