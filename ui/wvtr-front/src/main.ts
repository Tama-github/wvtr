import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import VueCookies from 'vue-cookies';


createApp(App).use(VueCookies).mount('#app')
