<script lang="ts" setup>
import Icon from '@/components/Icons.vue'
import { useAppStore } from '@/composables/use-app'
import { useToggle } from '@vueuse/core'
import { NButton } from 'naive-ui'
import { useRouter } from 'vue-router'

const { isDark } = useAppStore()
const router = useRouter()
const toggleDark = useToggle(isDark)
</script>

<template>
  <n-layout-header bordered class="relative h-16">
    <n-grid :cols="3" class="h-full">
      <n-gi class="flex items-center">
        <div class="flex items-center gap-2 px-4">
          <Icon type="cube" class="text-2xl" />
          <span class="text-lg font-semibold">Blockchain Simulator</span>
        </div>
      </n-gi>

      <n-gi class="flex items-center justify-center">
        <n-breadcrumb>
          <n-breadcrumb-item @click="router.push('/what-is-blockchain')">
            What is Blockchain?
          </n-breadcrumb-item>
          <n-breadcrumb-item @click="router.push('/')">
            Blockchain
          </n-breadcrumb-item>
        </n-breadcrumb>
      </n-gi>

      <n-gi class="flex items-center justify-end gap-4 px-4">
        <ChainVerification />
        <n-button circle @click.prevent="toggleDark()">
          <template #icon>
            <Icon v-if="isDark" type="sunny" />
            <Icon v-else type="moon" />
          </template>
        </n-button>
      </n-gi>
    </n-grid>
  </n-layout-header>
</template>
