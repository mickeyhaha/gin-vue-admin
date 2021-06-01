import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

//获取原型对象上的push函数
const originalPush = Router.prototype.push
//修改原型对象中的push方法
Router.prototype.push = function push(location) {
    return originalPush.call(this, location).catch(err => err)
}

const baseRouters = [{
    path: '/',
    redirect: '/login'
},
    {
        path: "/smtboard",
        name: 'SmtBoard',
        component: () =>
            import('@/view/datav/index_line.vue')
    },
    {
        path: "/deptboard",
        name: 'DeptBoard',
        component: () =>
            import('@/view/datav/index_dept.vue')
    },
    {
        path: "/deptdashboard",
        name: 'DeptDashboard',
        component: () =>
            import('@/view/datavdash/index_dept_dash.vue')
    },
    // {
    //     path: "/spcdashboard",
    //     name: 'SpcDashboard',
    //     component: () =>
    //         import('@/view/datavspc/index.vue')
    // },
{
    path: "/init",
    name: 'init',
    component: () =>
        import('@/view/init/init.vue')
},
{
    path: '/login',
    name: 'login',
    component: () =>
        import('@/view/login/login.vue')
}
]

// 需要通过后台数据来生成的组件

const createRouter = () => new Router({
    routes: baseRouters
})

const router = createRouter()

export default router
