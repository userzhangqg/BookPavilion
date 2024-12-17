<template>
  <div class="book-list">
    <el-table
      v-loading="loading"
      :data="books"
      style="width: 100%"
    >
      <el-table-column
        prop="title"
        label="书名"
        width="250"
      />
      <el-table-column
        prop="author"
        label="作者"
        width="150"
      />
      <el-table-column
        prop="format"
        label="格式"
        width="100"
      />
      <el-table-column
        label="文件大小"
        width="120"
      >
        <template #default="{ row }">
          {{ formatFileSize(row.fileSize) }}
        </template>
      </el-table-column>
      <el-table-column
        prop="createdAt"
        label="上传时间"
        width="180"
      >
        <template #default="{ row }">
          {{ formatDate(row.createdAt) }}
        </template>
      </el-table-column>
      <el-table-column
        label="操作"
        width="200"
      >
        <template #default="{ row }">
          <el-button
            size="small"
            @click="viewBook(row.id)"
          >
            查看
          </el-button>
          <el-button
            size="small"
            type="primary"
            @click="readBook(row.id)"
          >
            阅读
          </el-button>
          <el-button
            size="small"
            type="danger"
            @click="handleDelete(row.id)"
          >
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <div class="pagination">
      <el-pagination
        v-model:current-page="currentPage"
        :page-size="pageSize"
        :total="totalBooks"
        layout="total, prev, pager, next"
        @current-change="handlePageChange"
      />
    </div>
  </div>
</template>

<script>
import { computed, onMounted } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'

export default {
  name: 'BookList',
  setup() {
    const store = useStore()
    const router = useRouter()

    const books = computed(() => store.state.books)
    const loading = computed(() => store.state.loading)
    const totalBooks = computed(() => store.state.totalBooks)
    const currentPage = computed(() => store.state.currentPage)
    const pageSize = computed(() => store.state.pageSize)

    onMounted(() => {
      loadBooks()
    })

    const loadBooks = async () => {
      await store.dispatch('fetchBooks', {
        page: currentPage.value,
        pageSize: pageSize.value
      })
    }

    const handlePageChange = (page) => {
      store.dispatch('fetchBooks', { page, pageSize: pageSize.value })
    }

    const viewBook = (id) => {
      router.push(`/books/${id}`)
    }

    const readBook = (id) => {
      router.push(`/reading/${id}`)
    }

    const handleDelete = (id) => {
      ElMessageBox.confirm(
        '确定要删除这本书吗？',
        '警告',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          try {
            await store.dispatch('deleteBook', id)
            ElMessage({
              type: 'success',
              message: '删除成功',
            })
          } catch (error) {
            ElMessage({
              type: 'error',
              message: '删除失败',
            })
          }
        })
        .catch(() => {
          ElMessage({
            type: 'info',
            message: '已取消删除',
          })
        })
    }

    const formatFileSize = (size) => {
      if (size < 1024) return size + ' B'
      if (size < 1024 * 1024) return (size / 1024).toFixed(2) + ' KB'
      if (size < 1024 * 1024 * 1024) return (size / (1024 * 1024)).toFixed(2) + ' MB'
      return (size / (1024 * 1024 * 1024)).toFixed(2) + ' GB'
    }

    const formatDate = (date) => {
      return new Date(date).toLocaleString()
    }

    return {
      books,
      loading,
      totalBooks,
      currentPage,
      pageSize,
      handlePageChange,
      viewBook,
      readBook,
      handleDelete,
      formatFileSize,
      formatDate
    }
  }
}
</script>

<style scoped>
.book-list {
  padding: 20px;
}

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: center;
}
</style>
