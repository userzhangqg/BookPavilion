<template>
  <div class="book-detail">
    <el-container v-loading="loading">
      <el-header>
        <div class="header-content">
          <el-page-header @back="goBack">
            <template #content>
              <span class="page-title">图书详情</span>
            </template>
          </el-page-header>
        </div>
      </el-header>

      <el-main>
        <template v-if="book">
          <el-card class="detail-card">
            <template #header>
              <div class="card-header">
                <h2>{{ book.title }}</h2>
                <div class="header-actions">
                  <el-button
                    type="primary"
                    @click="startReading"
                  >
                    开始阅读
                  </el-button>
                </div>
              </div>
            </template>

            <el-descriptions :column="2" border>
              <el-descriptions-item label="作者">
                {{ book.author || '未知' }}
              </el-descriptions-item>
              <el-descriptions-item label="格式">
                {{ book.format?.toUpperCase() }}
              </el-descriptions-item>
              <el-descriptions-item label="文件大小">
                {{ formatFileSize(book.fileSize) }}
              </el-descriptions-item>
              <el-descriptions-item label="上传时间">
                {{ formatDate(book.createdAt) }}
              </el-descriptions-item>
            </el-descriptions>

            <div class="actions">
              <el-button
                type="danger"
                @click="handleDelete"
              >
                删除图书
              </el-button>
            </div>
          </el-card>
        </template>

        <el-empty
          v-else-if="!loading"
          description="未找到图书信息"
        />
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { computed, onMounted } from 'vue'
import { useStore } from 'vuex'
import { useRouter, useRoute } from 'vue-router'
import { ElMessageBox, ElMessage } from 'element-plus'

export default {
  name: 'BookDetail',
  setup() {
    const store = useStore()
    const router = useRouter()
    const route = useRoute()

    const book = computed(() => store.state.currentBook)
    const loading = computed(() => store.state.loading)

    onMounted(async () => {
      const bookId = parseInt(route.params.id)
      if (bookId) {
        await store.dispatch('fetchBookById', bookId)
      }
    })

    const goBack = () => {
      router.push('/')
    }

    const startReading = () => {
      router.push(`/reading/${book.value.id}`)
    }

    const handleDelete = () => {
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
            await store.dispatch('deleteBook', book.value.id)
            ElMessage({
              type: 'success',
              message: '删除成功',
            })
            router.push('/')
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
      if (!size) return '0 B'
      if (size < 1024) return size + ' B'
      if (size < 1024 * 1024) return (size / 1024).toFixed(2) + ' KB'
      if (size < 1024 * 1024 * 1024) return (size / (1024 * 1024)).toFixed(2) + ' MB'
      return (size / (1024 * 1024 * 1024)).toFixed(2) + ' GB'
    }

    const formatDate = (date) => {
      if (!date) return '未知'
      return new Date(date).toLocaleString()
    }

    return {
      book,
      loading,
      goBack,
      startReading,
      handleDelete,
      formatFileSize,
      formatDate
    }
  }
}
</script>

<style scoped>
.book-detail {
  min-height: 100vh;
}

.el-header {
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12);
  position: relative;
  z-index: 1;
}

.header-content {
  height: 60px;
  display: flex;
  align-items: center;
}

.page-title {
  font-size: 18px;
  font-weight: bold;
}

.el-main {
  background-color: #f5f7fa;
  padding: 20px;
}

.detail-card {
  max-width: 800px;
  margin: 0 auto;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-header h2 {
  margin: 0;
  font-size: 20px;
}

.actions {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>
