<script lang="ts" setup>
import { useBlockchain } from '@/composables/use-blockchain';

const { verifyChain, isVerifying } = useBlockchain()
</script>

<template>
  <div class="flex justify-end">
    <n-button @click="verifyChain()" :loading="isVerifying">
      Validate Chain
    </n-button>
  </div>
</template>