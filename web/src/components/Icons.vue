<template>
  <n-icon v-bind="props">
    <icon />
  </n-icon>
</template>

<script lang="ts" setup>
import {
  AddOutline as add,
  MoonOutline as moon,
  SunnyOutline as sunny,
} from '@vicons/ionicons5'
import { NIcon } from 'naive-ui'

const icons = {
  add,
  sunny,
  moon,
}

type IconTypes = keyof typeof icons

const props = defineProps({
  ...NIcon.props,
  type: {
    type: String,
    default: 'fallback',
  },
})

const icon = icons[props.type as IconTypes]
</script>
