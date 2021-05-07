import { createApp } from 'vue'
import App from './App.vue'
import './index.css'
import router from './router'
import ElementPlus from 'element-plus'
import '/@modules/element-plus/lib/theme-chalk/index.css'

const app = createApp(App)

app.use(ElementPlus, { size: 'small', zIndex: 3000, })

import ElUploadOss from './components/UploadOSS.vue'
import ElManageOss from './components/ManageOSS.vue'
import ElMarkdownDisplay from './components/MarkdownDisplay.vue'

const components = [
ElUploadOss,
ElManageOss,
ElMarkdownDisplay,
]
components.forEach(component => {
  app.component(component.name, component)
})

app.use(router)
app.mount('#app')
