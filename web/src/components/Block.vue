<script lang="ts" setup>
import { useBlockchain } from '@/composables/use-blockchain';
import { useNotification } from 'naive-ui';
import { watch } from 'vue';

const { blockchain, isLoadingBlocks, blocksError } = useBlockchain()

const formatDate = (date: string) =>
  new Intl.DateTimeFormat('en-GB', {
    year: 'numeric',
    month: 'long',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
  }).format(new Date(date)).replace(' at', ',')

const notification = useNotification()

watch(blocksError, (error) => {
  if (error) {
    notification.error({
      content: error.message,
      duration: 2500,
      keepAliveOnHover: true,
    })
  }
},
  { immediate: true }
)

</script>

<template>
  <div class="py-4">
    <n-spin v-if="isLoadingBlocks" size="large" />

    <n-empty v-else-if="!blockchain" description="No blockchain data available" />

    <template v-else>
      <div class="flex gap-8 mb-4">
        <n-statistic label="Total Blocks">
          {{ blockchain.stats.totalBlocks }}
        </n-statistic>
        <n-statistic label="Average Mine Time"> {{ blockchain.stats.avgMineTime }}s </n-statistic>
      </div>

      <div class="grid grid-cols-[repeat(auto-fill,minmax(300px,1fr))] gap-4">
        <div v-for="(item, index) in blockchain.blocks" :key="item.hash">
          <n-card :title="item.index === 0 ? 'Block 0' : `Block ${index}`" class="min-h-96">
            <n-space vertical>
              <n-text ellipsis>Hash: {{ item.hash }}</n-text>
              <n-text>Prev hash: {{ !!item.prevHash ? item.prevHash : 0 }}</n-text>
              <n-text>Data: {{ item.data }}</n-text>
              <n-text>Timestamp: {{ formatDate(item.timestamp) }}</n-text>
              <n-text>Nonce: {{ item.nonce }}</n-text>
              <n-text>Difficulty: {{ item.difficulty }}</n-text>
            </n-space>
          </n-card>
        </div>
      </div>
    </template>
  </div>
</template>
