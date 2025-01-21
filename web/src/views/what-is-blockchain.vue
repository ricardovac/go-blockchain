<script setup lang="ts">
import { NSpace, NCard, NText, NH1, NH2, NList, NListItem, NDivider } from 'naive-ui'
</script>

<template>
  <n-space vertical size="large" class="p-8">
    <n-card>
      <n-h1>What is Blockchain?</n-h1>
      <n-text>
        A blockchain is a distributed digital ledger that stores data in blocks that are linked together in
        chronological order.
        Think of it as a chain of blocks, where each block contains information and is connected to the previous block.
      </n-text>
    </n-card>

    <n-card>
      <n-h2>Key Components</n-h2>
      <n-list>
        <n-list-item>
          <n-text strong>Block</n-text>
          <n-text>A container that stores data like transactions, timestamps, and a link to the previous block</n-text>
        </n-list-item>
        <n-list-item>
          <n-text strong>Hash</n-text>
          <n-text>A unique digital fingerprint that identifies each block and its contents</n-text>
        </n-list-item>
        <n-list-item>
          <n-text strong>Chain</n-text>
          <n-text>The connection between blocks using their hashes, creating an immutable sequence</n-text>
        </n-list-item>
      </n-list>
    </n-card>

    <n-card>
      <n-h2>How Does It Work?</n-h2>
      <n-space vertical>
        <n-text>1. When new data needs to be stored, it's placed in a block</n-text>
        <n-text>2. The block is given a unique hash based on its contents</n-text>
        <n-text>3. The block stores the hash of the previous block, creating a chain</n-text>
        <n-text>4. If any data in a block changes, its hash changes, breaking the chain</n-text>
        <n-text>5. Multiple copies of the chain are stored across a network of computers</n-text>
      </n-space>
    </n-card>

    <n-card>
      <n-h2>Benefits of Blockchain</n-h2>
      <n-list>
        <n-list-item>
          <n-text strong>Decentralization</n-text>
          <n-text>No single entity controls the data</n-text>
        </n-list-item>
        <n-list-item>
          <n-text strong>Transparency</n-text>
          <n-text>All transactions are visible to network participants</n-text>
        </n-list-item>
        <n-list-item>
          <n-text strong>Immutability</n-text>
          <n-text>Once data is recorded, it cannot be changed without detection</n-text>
        </n-list-item>
        <n-list-item>
          <n-text strong>Security</n-text>
          <n-text>Cryptography and distributed nature make it highly secure</n-text>
        </n-list-item>
      </n-list>
    </n-card>

    <n-card>
      <n-h2>Common Applications</n-h2>
      <n-space vertical>
        <n-text>• Cryptocurrencies like Bitcoin</n-text>
        <n-text>• Smart contracts</n-text>
        <n-text>• Supply chain tracking</n-text>
        <n-text>• Digital identity verification</n-text>
        <n-text>• Voting systems</n-text>
      </n-space>
    </n-card>
  </n-space>
</template>
