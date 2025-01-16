<script lang="ts" setup>
import { mineBlock } from '@/api/blockchain'
import { useMutation, useQueryClient } from '@tanstack/vue-query'
import { Add } from '@vicons/ionicons5'
import { NIcon } from 'naive-ui'
import { h, ref } from 'vue'

interface MineBlockPayload {
  data: string
  difficulty: number
}

const value = ref('')
const difficultySlider = ref(1)
const queryClient = useQueryClient()

function renderIcon() {
  return h(NIcon, null, {
    default: () => h(Add),
  })
}

const { mutate, error, isPending } = useMutation({
  mutationFn: (payload: MineBlockPayload) => mineBlock(payload),
  onSuccess: () => {
    queryClient.invalidateQueries({ queryKey: ['blocks'] })
    value.value = ''
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
        <n-button
          :render-icon="renderIcon"
          @click="handleMineBlock"
          :disabled="isPending || !value"
          :loading="isPending"
          >Mine Block</n-button
        >
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
