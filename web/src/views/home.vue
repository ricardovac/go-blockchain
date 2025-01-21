<script lang="ts" setup>

import BlockCard from '@/components/Block.vue';
import ChainVerification from '@/components/ChainVerification.vue';
import Mine from '@/components/Mine.vue';

</script>

<template>
  <n-space vertical size="large">
    <mine />
    <chain-verification />
    <block-card />
  </n-space>
  <n-back-top />
</template>
