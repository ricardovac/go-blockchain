import { useDark } from '@vueuse/core'
import { acceptHMRUpdate, defineStore } from 'pinia'

export const useAppStore = defineStore('app', () => {
  const isDark = useDark({
    selector: 'html',
    attribute: 'class',
    valueDark: 'dark',
    valueLight: 'light',
  })

  const setIsDark = (value: boolean) => {
    isDark.value = value
  }

  return {
    setIsDark,
    isDark,
  } as const
})

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useAppStore, import.meta.hot))
}
