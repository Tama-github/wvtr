import './assets/main.css'

import { createApp, ref } from 'vue'
import App from './App.vue'
import VueCookies from 'vue-cookies';
import { NavigationHandler } from './model/utils';



const app = createApp(App).use(VueCookies);

app.provide('navigationHandler', new NavigationHandler())

app.mount('#app')
