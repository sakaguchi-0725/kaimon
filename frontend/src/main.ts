import { createApp } from 'vue'
import './app/assets/css/index.css'
import App from './app/app.vue'
import { createAppRouter } from './app/routes'

const app = createApp(App)
app.use(createAppRouter('web'))
app.mount('#app')
