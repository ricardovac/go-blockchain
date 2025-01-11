import { useDark, useStorage } from '@vueuse/core'
import { acceptHMRUpdate, defineStore } from 'pinia'

export const useAppStore = defineStore('app', () => {
  const isDark = useDark({
    selector: 'html',
    attribute: 'class',
    valueDark: 'dark',
    valueLight: 'light',
  })

  const sidebarCollapsed = useStorage('sidebarCollapsed', false)

  const toggleSidebarCollapsed = async () => {
    sidebarCollapsed.value = !sidebarCollapsed.value
  }

  const setIsDark = (value: boolean) => {
    isDark.value = value
  }

  return {
    toggleSidebarCollapsed,
    sidebarCollapsed,
    setIsDark,
    isDark,
  } as const
})

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useAppStore, import.meta.hot))
}
