import { createRouter, createWebHistory,createWebHashHistory } from 'vue-router'

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
    {
        path: '/',
        component: () => import('/@/views/Index.vue'),
    },
    {
        path: '/test',
        component: () => import('/@/views/Files.vue'),
    },


    ]
})

router.beforeEach(() => {
    console.log('路由拦截')
})

export default router