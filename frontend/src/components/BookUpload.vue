<template>
  <div class="book-upload">
    <el-form
      ref="uploadForm"
      :model="formData"
      :rules="rules"
      label-width="80px"
    >
      <el-form-item
        label="书名"
        prop="title"
      >
        <el-input v-model="formData.title" placeholder="请输入书名" />
      </el-form-item>

      <el-form-item
        label="作者"
        prop="author"
      >
        <el-input v-model="formData.author" placeholder="请输入作者" />
      </el-form-item>

      <el-form-item
        label="文件"
        prop="file"
      >
        <el-upload
          ref="upload"
          class="upload-demo"
          drag
          action="#"
          :auto-upload="false"
          :on-change="handleFileChange"
          :limit="1"
        >
          <el-icon class="el-icon--upload">
            <upload-filled />
          </el-icon>
          <div class="el-upload__text">
            拖拽文件到此处或 <em>点击上传</em>
          </div>
          <template #tip>
            <div class="el-upload__tip">
              支持的格式：PDF、EPUB、TXT、MOBI
            </div>
          </template>
        </el-upload>
      </el-form-item>

      <el-form-item>
        <el-button
          type="primary"
          :loading="uploading"
          @click="submitUpload"
        >
          上传
        </el-button>
        <el-button @click="resetForm">重置</el-button>
      </el-form-item>
    </el-form>

    <el-dialog
      v-model="uploadProgress.visible"
      title="上传进度"
      width="30%"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      :show-close="false"
    >
      <el-progress
        :percentage="uploadProgress.percent"
        :status="uploadProgress.status"
      />
      <div class="progress-text">{{ uploadProgress.text }}</div>
    </el-dialog>
  </div>
</template>

<script>
import { ref, reactive } from 'vue'
import { useStore } from 'vuex'
import { ElMessage } from 'element-plus'
import { UploadFilled } from '@element-plus/icons-vue'

export default {
  name: 'BookUpload',
  components: {
    UploadFilled
  },
  emits: ['upload-success'],
  setup(props, { emit }) {
    const store = useStore()
    const uploadForm = ref(null)
    const upload = ref(null)
    const uploading = ref(false)

    const formData = reactive({
      title: '',
      author: '',
      file: null
    })

    const rules = {
      title: [
        { required: true, message: '请输入书名', trigger: 'blur' }
      ],
      file: [
        { required: true, message: '请选择文件', trigger: 'change' }
      ]
    }

    const uploadProgress = reactive({
      visible: false,
      percent: 0,
      status: '',
      text: '正在上传...'
    })

    const handleFileChange = (file) => {
      const allowedTypes = ['application/pdf', 'application/epub+zip', 'text/plain']
      const maxSize = 100 * 1024 * 1024 // 100MB

      if (!allowedTypes.includes(file.raw.type)) {
        ElMessage.error('不支持的文件格式')
        upload.value.clearFiles()
        return false
      }

      if (file.size > maxSize) {
        ElMessage.error('文件大小不能超过100MB')
        upload.value.clearFiles()
        return false
      }

      formData.file = file.raw
    }

    const submitUpload = async () => {
      if (!formData.file) {
        ElMessage.error('请选择文件')
        return
      }

      await uploadForm.value.validate(async (valid) => {
        if (valid) {
          try {
            uploading.value = true
            uploadProgress.visible = true

            const form = new FormData()
            form.append('title', formData.title)
            form.append('author', formData.author)
            form.append('file', formData.file)

            await store.dispatch('createBook', form)

            uploadProgress.percent = 100
            uploadProgress.status = 'success'
            uploadProgress.text = '上传成功'

            ElMessage.success('上传成功')
            emit('upload-success')
            resetForm()

            setTimeout(() => {
              uploadProgress.visible = false
              uploadProgress.percent = 0
              uploadProgress.status = ''
              uploadProgress.text = '正在上传...'
            }, 1500)
          } catch (error) {
            uploadProgress.status = 'exception'
            uploadProgress.text = '上传失败'
            ElMessage.error('上传失败：' + error.message)
          } finally {
            uploading.value = false
          }
        }
      })
    }

    const resetForm = () => {
      uploadForm.value.resetFields()
      upload.value.clearFiles()
      formData.file = null
    }

    return {
      uploadForm,
      upload,
      formData,
      rules,
      uploading,
      uploadProgress,
      handleFileChange,
      submitUpload,
      resetForm
    }
  }
}
</script>

<style scoped>
.book-upload {
  max-width: 600px;
  margin: 20px auto;
  padding: 20px;
}

.progress-text {
  text-align: center;
  margin-top: 10px;
}

.el-upload__tip {
  color: #909399;
}
</style>
