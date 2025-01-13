import request from '@/utils/request'

interface BlockchainStats {
  totalBlocks: number
  avgMineTime: number
}

interface BlockProgress {
  currentHash: string
  hashesPerSec: number
  target: string
}

interface BlockStats {
  minedAt: string
  miningTime: number
  attempts: number
  progress: BlockProgress
}

interface Block {
  index: number
  timestamp: string
  bpm: number
  hash: string
  prevHash: string
  nonce: number
  difficulty: number
  stats: BlockStats
}

export interface BlockchainResponse {
  blocks: Block[]
  difficulty: number
  stats: BlockchainStats
}

export const getBlockchain = async (): Promise<BlockchainResponse> => {
  const response = await request.get('/blocks')
  return response.data
}
