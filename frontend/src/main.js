import { createApp } from 'vue'
import App from './App.vue'
import './index.css'
import ElementPlus from 'element-plus'
import '/@modules/element-plus/lib/theme-chalk/index.css'


const app = createApp(App)

app.use(ElementPlus, { size: 'small', zIndex: 3000, })

import ElUploadOss from './components/UploadOSS.vue'

const components = [
ElUploadOss
]
components.forEach(component => {
  app.component(component.name, component)
})

app.mount('#app')
