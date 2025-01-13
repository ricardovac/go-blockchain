<script lang="ts" setup>
import { getBlockchain } from '@/api/blockchain'
import { useQuery } from '@tanstack/vue-query'

import BlockCard from '@/components/BlockCard.vue'

const { data, isLoading, error } = useQuery({ queryKey: ['blockchain'], queryFn: getBlockchain })
</script>

<template>
  <div class="container">
    <n-alert v-if="error" type="error" :title="error.message" />

    <block-card :data="data" :is-loading="isLoading" />
  </div>
</template>
