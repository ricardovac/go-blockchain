<script setup lang="ts">
import { darkTheme, NConfigProvider } from 'naive-ui'
import { computed } from 'vue'
import Layout from './layouts/Layout.vue'
import { useAppStore } from './pinia/modules/app'

const appStore = useAppStore()
const theme = computed(() => (appStore.isDark ? darkTheme : null))

defineOptions({
  name: 'App',
})
</script>

<template>
  <n-config-provider :theme="theme">
    <Layout>
      <router-view />
    </Layout>
  </n-config-provider>
</template>
