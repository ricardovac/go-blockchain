<script lang="ts" setup>
import { getBlocks } from '@/api/blockchain'
import { useQuery } from '@tanstack/vue-query'

import BlockCard from '@/components/Block.vue'
import Mine from '@/components/Mine.vue'

const { data, isLoading, error } = useQuery({ queryKey: ['blocks'], queryFn: getBlocks })
</script>

<template>
  <div class="container">
    <n-alert v-if="error" type="error" :title="error.message" />

    <mine />

    <block-card :data="data" :is-loading="isLoading" />
  </div>
</template>
