<script lang="ts" setup>
import { mineBlock } from '@/api/blockchain'
import { useMutation, useQueryClient } from '@tanstack/vue-query'
import { useNotification } from 'naive-ui'
import { ref } from 'vue'
import Icon from './Icons.vue'

interface MineBlockPayload {
  data: string
  difficulty: number
}

const value = ref('')
const difficultySlider = ref(1)
const queryClient = useQueryClient()
const notification = useNotification()

const { mutate, error, isPending } = useMutation({
  mutationFn: (payload: MineBlockPayload) => mineBlock(payload),
  onSuccess: () => {
    queryClient.invalidateQueries({ queryKey: ['blocks'] })
    value.value = ''
    notification.success({
      content: `Block mined and added to the chain`,
      duration: 2500,
      keepAliveOnHover: true,
    })
  },
})

function handleMineBlock() {
  mutate({
    data: value.value,
    difficulty: difficultySlider.value,
  })
}
</script>

<template>
  <n-card>
    <div class="space-y-4">
      <n-alert v-if="error" type="error" :title="error?.response?.data?.error || 'Error'" closable>
        {{ error?.response?.data?.details || 'Failed to mine block' }}
      </n-alert>

      <div class="flex gap-2">
        <n-input v-model:value="value" type="text" placeholder="Enter block data" />
        <n-button :loading="isPending" @click="handleMineBlock" :disabled="isPending || !value">
          <Icon type="add" v-if="!isPending" />
          Mine Block
        </n-button>
      </div>

      <div class="flex items-center gap-4">
        <span>Difficulty:</span>
        <div class="w-96">
          <n-slider v-model:value="difficultySlider" :min="1" :max="4" :step="1" />
        </div>
        <span> {{ difficultySlider }} </span>
      </div>
    </div>
  </n-card>
</template>
