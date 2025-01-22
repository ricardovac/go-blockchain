import request from '@/utils/request'
import type { BlockchainResponse } from '@/utils/types'

export interface ApiErrorResponse {
  error: string
  code: number
  details: string
}

export const getBlocks = async (): Promise<BlockchainResponse> => {
  const response = await request.get('/blocks')
  return response.data
}

export const verifyChain = async (): Promise<{ valid: boolean; error?: string }> => {
  const response = await request.get('/blocks/verify')
  return response.data
}

export const mineBlock = async (payload: {
  data: string
  difficulty: number
}): Promise<BlockchainResponse> => {
  const response = await request.post('/blocks', payload)
  return response.data
}