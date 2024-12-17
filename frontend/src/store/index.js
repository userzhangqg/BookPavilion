import { createStore } from 'vuex'
import axios from 'axios'

export default createStore({
  state: {
    books: [],
    currentBook: null,
    totalBooks: 0,
    currentPage: 1,
    pageSize: 10,
    loading: false,
    error: null
  },
  mutations: {
    SET_BOOKS(state, books) {
      state.books = books
    },
    SET_CURRENT_BOOK(state, book) {
      state.currentBook = book
    },
    SET_TOTAL_BOOKS(state, total) {
      state.totalBooks = total
    },
    SET_CURRENT_PAGE(state, page) {
      state.currentPage = page
    },
    SET_LOADING(state, loading) {
      state.loading = loading
    },
    SET_ERROR(state, error) {
      state.error = error
    }
  },
  actions: {
    async fetchBooks({ commit }, { page = 1, pageSize = 10 } = {}) {
      commit('SET_LOADING', true)
      try {
        const response = await axios.get(`/api/books?page=${page}&page_size=${pageSize}`)
        commit('SET_BOOKS', response.data.books)
        commit('SET_TOTAL_BOOKS', response.data.total)
        commit('SET_CURRENT_PAGE', page)
      } catch (error) {
        commit('SET_ERROR', error.message)
      } finally {
        commit('SET_LOADING', false)
      }
    },
    async fetchBookById({ commit }, id) {
      console.log(`Fetching book with ID: ${id}`); // Log the book ID being fetched
      commit('SET_LOADING', true)
      try {
        const response = await axios.get(`/api/books/${id}`)
        commit('SET_CURRENT_BOOK', response.data)
      } catch (error) {
        commit('SET_ERROR', error.message)
      } finally {
        commit('SET_LOADING', false)
      }
    },
    async createBook({ dispatch }, formData) {
      try {
        await axios.post('/api/books', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })
        await dispatch('fetchBooks')
      } catch (error) {
        throw error
      }
    },
    async deleteBook({ dispatch }, id) {
      try {
        await axios.delete(`/api/books/${id}`)
        await dispatch('fetchBooks')
      } catch (error) {
        throw error
      }
    }
  },
  getters: {
    getBookById: (state) => (id) => {
      return state.books.find(book => book.id === id)
    }
  }
})
