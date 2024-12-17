<template>
  <div class="home">
    <el-container>
      <el-header>
        <div class="header-content">
          <h1>图书阅读馆</h1>
          <el-button type="primary" @click="showUpload = true">
            上传图书
          </el-button>
        </div>
      </el-header>

      <el-main>
        <el-dialog
          v-model="showUpload"
          title="上传图书"
          width="50%"
          destroy-on-close
        >
          <book-upload @upload-success="handleUploadSuccess" />
        </el-dialog>

        <book-list />
      </el-main>

      <el-footer>
        <p>© 2024 图书阅读馆. All rights reserved.</p>
      </el-footer>
    </el-container>
  </div>
</template>

<script>
import { ref } from 'vue'
import BookList from '../components/BookList.vue'
import BookUpload from '../components/BookUpload.vue'

export default {
  name: 'Home',
  components: {
    BookList,
    BookUpload
  },
  setup() {
    const showUpload = ref(false)

    const handleUploadSuccess = () => {
      showUpload.value = false
    }

    return {
      showUpload,
      handleUploadSuccess
    }
  }
}
</script>

<style scoped>
.home {
  min-height: 100vh;
}

.el-container {
  min-height: 100vh;
}

.el-header {
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12);
  position: relative;
  z-index: 1;
}

.header-content {
  max-width: 1200px;
  margin: 0 auto;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.header-content h1 {
  margin: 0;
  font-size: 24px;
  color: #409EFF;
}

.el-main {
  background-color: #f5f7fa;
  padding: 20px;
}

.el-footer {
  text-align: center;
  color: #909399;
  padding: 20px;
  background-color: #fff;
}
</style>
