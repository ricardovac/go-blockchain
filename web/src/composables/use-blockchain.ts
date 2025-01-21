import { getBlocks, mineBlock, verifyChain } from '@/api/blockchain'
import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query'
import { createGlobalState } from '@vueuse/core'
import { useNotification, type NotificationType } from 'naive-ui'
import { ref } from 'vue'

export const useBlockchain = createGlobalState(() => {
  const notification = useNotification()
  const queryClient = useQueryClient()

  const miningValue = ref('')
  const difficultySlider = ref(1)

  function notify(type: NotificationType, options: { content: string }) {
    notification[type]({
      content: options.content,
      duration: 2500,
      keepAliveOnHover: true,
    })
  }

  const {
    data: blockchain,
    isLoading: isLoadingBlocks,
    error: blocksError,
  } = useQuery({
    queryKey: ['blocks'],
    queryFn: getBlocks,
  })

  const { mutate: verifyChainMutation, isPending: isVerifying } = useMutation({
    mutationFn: verifyChain,
    onSuccess: (result) => {
      notify(result.valid ? 'success' : 'error', {
        content: result.valid ? 'Chain is valid' : 'Chain is invalid',
      })
    },
    onError: (error) => notify('error', { content: error.message }),
  })

  const { mutate: mineMutation, isPending: isMining } = useMutation({
    mutationFn: (payload: { data: string; difficulty: number }) => mineBlock(payload),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['blocks'] })
      miningValue.value = ''
      notify('success', {
        content: 'Block mined and added to the chain',
      })
    },
    onError: (error) => notify('error', { content: error.message }),
  })

  function handleMineBlock() {
    if (!miningValue.value) return

    mineMutation({
      data: miningValue.value,
      difficulty: difficultySlider.value,
    })
  }

  return {
    miningValue,
    difficultySlider,
    blockchain,
    isLoadingBlocks,
    isMining,
    isVerifying,
    blocksError,
    verifyChain: verifyChainMutation,
    handleMineBlock,
  }
})
