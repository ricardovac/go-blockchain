<script lang="ts" setup>
import type { BlockchainResponse } from '@/api/blockchain'

const props = defineProps<{
  data?: BlockchainResponse
  isLoading: boolean
}>()
</script>

<template>
  <div class="p-4">
    <n-spin v-if="isLoading" size="large" />

    <n-empty v-else-if="!props.data" description="No blockchain data available" />

    <template v-else>
      <div class="flex gap-8 mb-4">
        <n-statistic label="Total Blocks">
          {{ props.data.stats.totalBlocks }}
        </n-statistic>
        <n-statistic label="Average Mine Time"> {{ props.data.stats.avgMineTime }}s </n-statistic>
      </div>

      <div class="grid grid-cols-[repeat(auto-fill,minmax(300px,1fr))] gap-4">
        <div v-for="(item, index) in props.data.blocks" :key="item.hash">
          <n-card :title="item.index === 0 ? 'Genesis Block' : `Block ${index}`">
            <n-space vertical>
              <n-text ellipsis>Hash: {{ item.hash }}</n-text>
              <n-text>Prev hash: {{ !!item.prevHash ? item.prevHash : 0 }}</n-text>
              <n-text>Time: {{ new Date(item.timestamp) }}</n-text>
              <n-text>Difficulty: {{ item.difficulty }}</n-text>
            </n-space>
          </n-card>
        </div>
      </div>
    </template>
  </div>
</template>
