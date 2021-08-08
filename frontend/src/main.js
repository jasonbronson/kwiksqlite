import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store/index'
import VueAxios from "vue-axios";
import ElementPlus from 'element-plus';
import 'element-plus/lib/theme-chalk/index.css';
import './css/main.css'
import { instance as axios } from './plugins/axios'

/* Default title tag */
const defaultDocumentTitle = 'KwikSqlite Admin'

/* Collapse mobile aside menu on route change & set document title from route meta */
router.afterEach(to => {
  
  if (to.meta && to.meta.title) {
    document.title = `${to.meta.title} - ${defaultDocumentTitle}`
  } else {
    document.title = defaultDocumentTitle
  }
})

const app = createApp(App).use(store).use(router).use(VueAxios, axios).use(ElementPlus).mount('#app')
// app.config.globalProperties.$axios 
// app.config.globalProperties.$localstorage

