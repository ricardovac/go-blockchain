<script setup lang="ts">
import Header from './Header.vue';
</script>

<template>
  <div class="min-h-screen">
    <Header />
    <main class="container mx-auto max-w-7xl py-8">
      <slot />
    </main>
  </div>
</template>
