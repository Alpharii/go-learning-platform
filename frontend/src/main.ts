import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

import App from './App.vue'
import router from './router'

const app = createApp(App)

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)
app.use(pinia)

app.use(router)

app.config.errorHandler = (err, instance, info) => {
    console.error('Global error handler triggered:', err)
    console.error('Component instance:', instance)
    console.error('Error info:', info)
}


app.mount('#app')
