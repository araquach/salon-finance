import Vue from 'vue'
import App from './App.vue'
import BaseComponent from './components/BaseComponent'
import JakComponent from './components/JakComponent'
import PkComponent from './components/PkComponent'

import Buefy from 'buefy'

Vue.use(Buefy)
Vue.use(require('vue-moment'))

Vue.component('BaseComponent', BaseComponent)
Vue.component('JakComponent', JakComponent)
Vue.component('PkComponent', PkComponent)

window.axios = require('axios')

new Vue({
    el: '#app',
    render: h => h(App)
})