<template>
  <div class="reading-view">
    <el-container>
      <el-header>
        <div class="header-content">
          <el-page-header @back="goBack">
            <template #content>
              <span class="book-title">{{ book?.title }}</span>
            </template>
          </el-page-header>
          <div class="header-actions">
            <el-button-group>
              <el-button :icon="ZoomOut" @click="zoomOut" />
              <el-button :icon="ZoomIn" @click="zoomIn" />
            </el-button-group>
          </div>
        </div>
      </el-header>

      <el-main>
        <div
          v-loading="loading"
          class="reader-container"
          :style="{ zoom: zoom }"
        >
          <!-- PDF阅读器 -->
          <template v-if="book?.format === 'pdf'">
            <iframe
              v-if="book?.filePath"
              :src="`/api/books/${book.id}/content`"
              class="pdf-viewer"
            />
          </template>

          <!-- EPUB阅读器 -->
          <template v-else-if="book?.format === 'epub'">
            <div class="epub-viewer">
              <div ref="epubViewer" />
            </div>
          </template>

          <!-- TXT阅读器 -->
          <template v-else-if="book?.format === 'txt'">
            <div
              v-if="textContent"
              class="txt-viewer"
            >
              {{ textContent }}
            </div>
          </template>

          <!-- 不支持的格式 -->
          <template v-else>
            <el-empty description="暂不支持该格式" />
          </template>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { useStore } from 'vuex'
import { useRouter, useRoute } from 'vue-router'
import { ZoomIn, ZoomOut } from '@element-plus/icons-vue'
import axios from 'axios'

export default {
  name: 'ReadingView',
  components: {
    ZoomIn,
    ZoomOut
  },
  setup() {
    const store = useStore()
    const router = useRouter()
    const route = useRoute()
    const epubViewer = ref(null)
    const loading = ref(true)
    const zoom = ref(1)
    const textContent = ref('')

    const book = computed(() => store.state.currentBook)

    onMounted(async () => {
      const bookId = parseInt(route.params.id)
      if (bookId) {
        await store.dispatch('fetchBookById', bookId)
        await loadContent()
      }
    })

    const loadContent = async () => {
      if (!book.value) return

      try {
        loading.value = true
        console.log(`Loading content for book ID: ${book.value.id}`); // Log the book ID being loaded

        if (book.value.format === 'txt') {
          const response = await axios.get(`/api/books/${book.value.id}/content`)
          textContent.value = response.data
        }
        // 其他格式的处理可以在这里添加

      } catch (error) {
        console.error('加载内容失败:', error)
      } finally {
        loading.value = false
      }
    }

    const goBack = () => {
      router.push(`/books/${book.value.id}`)
    }

    const zoomIn = () => {
      zoom.value = Math.min(zoom.value + 0.1, 2)
    }

    const zoomOut = () => {
      zoom.value = Math.max(zoom.value - 0.1, 0.5)
    }

    return {
      book,
      loading,
      epubViewer,
      zoom,
      textContent,
      goBack,
      zoomIn,
      zoomOut,
      ZoomIn,
      ZoomOut
    }
  }
}
</script>

<style scoped>
.reading-view {
  height: 100vh;
  display: flex;
  flex-direction: column;
}

.el-container {
  height: 100%;
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
  justify-content: space-between;
}

.book-title {
  font-size: 18px;
  font-weight: bold;
}

.el-main {
  padding: 0;
  background-color: #f5f7fa;
  overflow: hidden;
}

.reader-container {
  height: 100%;
  overflow: auto;
  background-color: #fff;
}

.pdf-viewer {
  width: 100%;
  height: 100%;
  border: none;
}

.epub-viewer {
  height: 100%;
  overflow: hidden;
}

.txt-viewer {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  line-height: 1.6;
  white-space: pre-wrap;
  font-family: 'Courier New', Courier, monospace;
}
</style>
