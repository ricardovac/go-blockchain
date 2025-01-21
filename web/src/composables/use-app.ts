import { createGlobalState, useDark } from '@vueuse/core'

export const useAppStore = createGlobalState(() => {
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
