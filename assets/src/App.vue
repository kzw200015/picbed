<template>
  <el-container>
    <el-header><Nav></Nav></el-header>
    <el-main><router-view></router-view></el-main>
  </el-container>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import Nav from './components/Nav.vue'

export default defineComponent({
  components: { Nav }
})
</script>

<style>
body {
  margin: 0;
}
</style>
