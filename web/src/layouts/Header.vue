<script lang="ts" setup>
import Icon from '@/components/Icon.vue'
import { useAppStore } from '@/stores/app'
import { useDark, useToggle } from '@vueuse/core'
import { NButton } from 'naive-ui'

const appStore = useAppStore()
const isDark = useDark({
  onChanged(dark: boolean) {
    appStore.setIsDark(dark)
  },
})
const toggleDark = useToggle(isDark)
</script>

<template>
  <n-layout-header bordered>
    <n-breadcrumb>
      <n-breadcrumb-item>Blockchain</n-breadcrumb-item>
    </n-breadcrumb>
    <n-space :size="20" class="items-center" style="line-height: 1">
      <n-button circle @click.prevent="toggleDark()">
        <template #icon>
          <Icon v-if="isDark" type="sunny" />
          <Icon v-else type="moon" />
        </template>
      </n-button>
    </n-space>
  </n-layout-header>
</template>

<style scoped lang="scss">
.n-layout-header {
  position: relative;
  display: flex;
  align-items: center;
  padding: 9px 18px;
  height: 60px;
}

.n-button {
  margin-right: 15px;
}

.n-space {
  margin-left: auto;
}
</style>
