<script setup lang="ts">
import { darkTheme, NConfigProvider, NNotificationProvider } from 'naive-ui'
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
    <n-notification-provider>
      <Layout>
        <router-view />
      </Layout>
      <n-back-top :right="50" />
    </n-notification-provider>
  </n-config-provider>
</template>

<style>
:root {
  --bg-color: 255, 255, 255;
}

html.dark {
  --bg-color: 16, 16, 20;
}

body {
  background-color: rgb(var(--bg-color));
}
</style>