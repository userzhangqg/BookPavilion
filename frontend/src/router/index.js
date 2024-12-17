import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import BookDetail from '../views/BookDetail.vue'
import ReadingView from '../views/ReadingView.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/books/:id',
    name: 'BookDetail',
    component: BookDetail
  },
  {
    path: '/reading/:id',
    name: 'ReadingView',
    component: ReadingView
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
