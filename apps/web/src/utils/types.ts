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
  data: string
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
