<template>
  <el-menu :default-active="route.path" mode="horizontal" router>
    <el-menu-item
      v-for="navItem in navItems"
      :key="navItem.name"
      :index="navItem.path"
      >{{ navItem.name }}</el-menu-item
    >
  </el-menu>
</template>
<script lang="ts">
import { defineComponent, ref } from 'vue'
import { useRoute } from 'vue-router'
import { routes } from '../routes'

interface NavItem {
  name: string
  path: string
}

const navItems = ref<Array<NavItem>>([])

routes.forEach(r => { navItems.value.push({ name: (r.meta?.title ?? '') as string, path: r.path }) })

export default defineComponent({
  setup() {
    const route = useRoute()

    return { route, navItems }
  }
})
</script>
