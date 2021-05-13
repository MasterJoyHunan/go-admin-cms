import Vue from 'vue'

import 'normalize.css/normalize.css' // A modern alternative to CSS resets

import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

import '@/styles/index.scss' // global css

import App from './App'
import store from './store'
import router from './router'

import '@/icons' // icon
import '@/permission' // permission control
import permission from '@/directive/permission/index'

Vue.use(ElementUI)
Vue.use(permission)
Vue.config.productionTip = false
permission.install(Vue)

new Vue({
    el: '#app',
    router,
    store,
    render: h => h(App)
})
