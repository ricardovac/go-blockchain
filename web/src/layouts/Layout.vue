<script setup lang="ts">
import { useAppStore } from '@/stores/app'
import { computed } from 'vue'
import Header from './Header.vue'
const appStore = useAppStore()
const isSidebarCollapsed = computed(() => appStore.sidebarCollapsed)
const toggleCollapsedSidebar = async () => await appStore.toggleSidebarCollapsed()
</script>

<template>
  <n-layout>
    <Header />
    <n-layout has-sider>
      <n-layout-sider
        bordered
        show-trigger
        collapse-mode="width"
        :collapsed="isSidebarCollapsed"
        :collapsed-width="64"
        :width="240"
        :native-scrollbar="false"
        @update:collapsed="toggleCollapsedSidebar"
      >
        <n-menu :collapsed-width="64" :collapsed-icon-size="22" />
      </n-layout-sider>
      <n-layout-content content-style="padding: 24px;">
        <slot />
      </n-layout-content>
    </n-layout>
  </n-layout>
</template>

<style scoped>
.n-layout {
  height: 100vh;
}
</style>
