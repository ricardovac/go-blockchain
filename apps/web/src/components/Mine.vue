<script lang="ts" setup>
import { useBlockchain } from '@/composables/use-blockchain';
import { computed } from 'vue';
import Icon from './Icons.vue';

const { handleMineBlock, isMining, miningValue, difficultySlider } = useBlockchain()

const difficulty = computed(() => {
  return difficultySlider.value
})
</script>

<template>
  <n-card>
    <div class="space-y-4">
      <div class="flex gap-2">
        <n-input v-model:value="miningValue" type="text" placeholder="Enter block data" />
        <n-button :loading="isMining" @click="handleMineBlock" :disabled="isMining || !miningValue">
          <Icon type="add" v-if="!isMining" />
          Mine Block
        </n-button>
      </div>

      <div class="flex items-center gap-4">
        <span>Difficulty:</span>
        <div class="w-96">
          <n-slider v-model:value="difficultySlider" :min="1" :max="4" :step="1" />
        </div>
        <span> {{ difficulty }} </span>
      </div>
    </div>
  </n-card>
</template>
