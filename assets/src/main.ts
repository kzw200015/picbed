import { ElButton, ElCard, ElCol, ElContainer, ElHeader, ElMain, ElMenu, ElMenuItem, ElNotification, ElPopconfirm, ElPopover, ElProgress, ElRow, ElSwitch, ElTooltip, ElUpload } from 'element-plus'
import 'element-plus/packages/theme-chalk/src/base.scss'
import { createApp } from 'vue'
import App from './App.vue'
import { router } from './routes'

const components = [ElSwitch, ElPopconfirm, ElPopover, ElTooltip, ElButton, ElCard, ElContainer, ElHeader, ElMain, ElMenu, ElMenuItem, ElNotification, ElUpload, ElRow, ElCol, ElProgress]

const app = createApp(App).use(router)

components.forEach(component => { app.use(component) })

app.mount('#app')