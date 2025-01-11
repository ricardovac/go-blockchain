<script setup lang="ts">
import { darkTheme, NConfigProvider, NGlobalStyle } from 'naive-ui'
import { computed } from 'vue'
import Layout from './layouts/Layout.vue'
import { useAppStore } from './stores/app'

const appStore = useAppStore()
const theme = computed(() => (appStore.isDark ? darkTheme : null))
</script>

<template>
  <n-config-provider :theme="theme">
    <n-global-style />
    <Layout>
      <router-view />
    </Layout>
  </n-config-provider>
</template>
