<template>
  <el-row justify="center">
    <el-col :md="10">
      <el-upload
        accept="image/*"
        drag
        :action="action"
        multiple
        style="text-align: center"
        :show-file-list="false"
        :http-request="uploadRequest"
      >
        <span class="el-icon-upload"></span>
        <div>将图片拖到此处，或点击上传</div>
      </el-upload>
    </el-col>
  </el-row>
  <el-row justify="center">
    <el-col :sm="18" :md="12">
      <el-card
        shadow="hover"
        class="file-item"
        v-for="(file, index) in fileList"
        :key="index"
      >
        <el-row :gutter="10"
          ><el-col class="file-item-col" :xs="6" :sm="8" :md="8" :lg="10">
            <el-tooltip :content="file.name">
              <div class="file-item-name">{{ file.name }}</div></el-tooltip
            > </el-col
          ><el-col class="file-item-col" :xs="8" :sm="8" :md="8" :lg="8"
            ><el-progress
              style="width: 100%"
              :percentage="file.percentage"
              :status="file.status"
            ></el-progress></el-col
          ><el-col class="file-item-col" :xs="8" :sm="6" :md="6" :lg="4"
            ><el-popover placement="top" trigger="hover">
              <div class="copy-button-container">
                <el-button type="primary" @click="handleCopyMarkdown(file)"
                  >Markdown</el-button
                ><el-button type="primary" @click="handleCopyHTML(file)"
                  >HTML</el-button
                ><el-button type="primary" @click="handleCopyURL(file)"
                  >URL</el-button
                >
              </div>
              <template #reference>
                <el-button>复制链接</el-button>
              </template>
            </el-popover>
          </el-col>
          <el-col class="file-item-col" :xs="2" :sm="2" :md="2" :lg="2">
            <el-popconfirm
              v-if="file.status === 'success'"
              title="确定要删除吗？"
              confirmButtonText="是"
              cancelButtonText="否"
              hideIcon
              @confirm="handleDeleteConfirm(index)"
            >
              <template #reference>
                <span class="el-icon-delete" style="color: #f56c6c"></span>
              </template>
            </el-popconfirm>
          </el-col>
        </el-row>
      </el-card>
    </el-col>
  </el-row>
</template>
<script lang="ts">
import { AxiosResponse } from 'axios'
import { defineComponent, ref } from 'vue'
import { backend, backendAPIURL, backendURL } from '../api'
import { ElNotification } from 'element-plus'
import { Image } from '../types'

interface ImageItem extends Image {
  percentage: number
  status?: string
}

const formatURL = (url: string) => {
  return backendURL + url
}

const handleCopy = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text)
    ElNotification({
      title: '成功',
      message: '复制成功',
      type: 'success'
    })
  } catch (error) {
    ElNotification({
      title: '错误',
      message: '复制失败',
      type: 'error'
    })
  }
}

export default defineComponent({
  setup() {
    const fileList = ref<Array<ImageItem>>([])
    const action = ref(backendAPIURL + '/upload')

    async function uploadRequest(options: unknown) {
      const ops = options as { action: string, file: File }
      const index = fileList.value.push({ id: '', name: ops.file.name, url: '', percentage: 0 }) - 1
      const formData = new FormData()
      formData.append('file', ops.file)
      let resp: AxiosResponse<Image> | undefined
      try {
        resp = await backend.post<Image>('/upload', formData, {
          onUploadProgress: e => {
            fileList.value[index].percentage = e.loaded / e.total * 100
          }
        })
      } catch (error) {
        console.log(error)
        fileList.value[index].status = 'exception'
        ElNotification({
          title: '错误',
          message: ops.file.name + ' ' + '上传失败',
          type: 'error'
        })
      }

      if (resp) {
        fileList.value[index].id = resp.data.id
        fileList.value[index].url = resp.data.url
        fileList.value[index].status = 'success'
      }
    }


    const handleCopyMarkdown = async (img: ImageItem) => {
      await handleCopy(`![${img.name}](${formatURL(img.url)})`)
    }

    const handleCopyHTML = async (img: ImageItem) => {
      await handleCopy(`<img src="${formatURL(img.url)}" alt="${img.name}">`)
    }

    const handleCopyURL = async (img: ImageItem) => {
      await handleCopy(formatURL(img.url))
    }


    const handleDeleteConfirm = async (index: number) => {
      const file = fileList.value[index]
      try {
        await backend.delete('/remove/' + file.id)
        fileList.value.splice(index, 1)
        ElNotification({
          title: '成功',
          message: '删除成功',
          type: 'success'
        })
      } catch (error) {
        console.log(error)
        ElNotification({
          title: '错误',
          message: '删除失败',
          type: 'error'
        })
      }
    }

    return {
      action, uploadRequest, fileList,
      handleCopyMarkdown, handleCopyHTML, handleCopyURL,
      handleDeleteConfirm
    }
  }
})
</script>
<style scoped>
.file-item {
  margin-top: 7px;
}
.file-item-col {
  display: flex;
  justify-content: center;
  align-items: center;
}
.file-item-name {
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}
.copy-button-container {
  display: flex;
  justify-content: space-between;
}
</style>