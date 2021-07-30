import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'


export const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'upload',
        component: () => import('../components/Upload.vue'),
        meta: { title: '上传图片' }
    }
]

export const router = createRouter({
    history: createWebHashHistory(),
    routes
})

router.afterEach((to) => {
    document.title = to.meta.title as string
})