<script lang="ts" setup>
import { getBlocks } from '@/api/blockchain';
import { useQuery } from '@tanstack/vue-query';

import BlockCard from '@/components/Block.vue';
import ChainVerification from '@/components/ChainVerification.vue';
import Mine from '@/components/Mine.vue';
const { data, isLoading, error } = useQuery({ queryKey: ['blocks'], queryFn: getBlocks })

</script>

<template>
  <n-alert v-if="error" type="error" :title="error.message" />
  <n-space vertical size="large">
    <mine />
    <chain-verification />
    <block-card :data="data" :is-loading="isLoading" />
  </n-space>
</template>
