<script setup lang="ts">
import { darkTheme, NConfigProvider, NNotificationProvider } from 'naive-ui'
import { computed } from 'vue'
import { useAppStore } from './composables/use-app'
import Layout from './layouts/Layout.vue'

const { isDark } = useAppStore()
const theme = computed(() => (isDark.value ? darkTheme : null))

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