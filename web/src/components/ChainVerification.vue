<script lang="ts" setup>
import { verifyChain } from '@/api/blockchain';
import { useMutation } from '@tanstack/vue-query';
import { useNotification, type NotificationType } from 'naive-ui';

const notification = useNotification()

const { mutate: verify, isPending: isLoading } = useMutation({
  mutationFn: verifyChain,
  onSuccess: (result) => {
    notify(result.valid ? 'success' : 'error')
  }
})

function notify(type: NotificationType) {
  notification[type]({
    content: type === 'success' ? 'Blockchain is valid' : 'Blockchain is invalid',
    duration: 2500,
    keepAliveOnHover: true,
  })
}
</script>

<template>
  <div class="flex justify-end">
    <n-button @click="verify()" :loading="isLoading">
      Validate Chain
    </n-button>
  </div>
</template>