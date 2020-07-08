import Vue from 'vue'
import App from './App.vue'
import router from './router'
import ElementUI from 'element-ui'
// import marked from 'marked'

import 'element-ui/lib/theme-chalk/index.css'
import '@/assets/style/index.less'

Vue.config.productionTip = false

Vue.use(ElementUI)
// Vue.use(marked())

new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>',
  render: h => h(App)
}).$mount('#app')
