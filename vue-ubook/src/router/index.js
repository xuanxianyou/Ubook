import Vue from 'vue'
import VueRouter from 'vue-router'
import login from '../views/login'
import register from '../views/register'
import home from '../views/home.vue'
import test from '../views/test'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Login',
    component: login
  },
  {
    path:'/login',
    name:'login',
    component:login
  },
  {
    path:'/register',
    name:'register',
    component:register
  },
  {
    path:'/home',
    name:'home',
    component:home
  },
  {
    path:'/test',
    name:'test',
    component:test
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
