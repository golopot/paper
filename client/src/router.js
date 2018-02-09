import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import About from './views/About.vue'
import Paper from './views/Paper.vue'
import Login from './views/Login.vue'

Vue.use(Router)
export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/about',
      name: 'about',
      component: About
    },
    {
      path: '/paper/:id',
      name: 'paper',
      component: Paper,
    },
    {
      path: '/login/',
      name: 'login',
      component: Login,
    }
  ]
})
