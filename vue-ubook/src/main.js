import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
// 导入插件
import './plugins'
// 导入element-ui组件
import ElementUI from 'element-ui'
// 导入样式
import 'element-ui/lib/theme-chalk/index.css'
import './assets/download/font_2587456_8idn4x5fjje/iconfont.css'
Vue.use(ElementUI)

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
