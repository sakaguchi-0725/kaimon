export type ButtonSize = 'full' | 'sm' | 'md' | 'lg'
export type ButtonVariant = 'primary' | 'outline'

export const SIZE_CLASSES: Record<ButtonSize, string> = {
  sm: 'py-1 px-3',
  md: 'py-2 px-4',
  lg: 'py-3 px-6',
  full: 'w-full py-2 px-4',
}

export const VARIANT_CLASSES: Record<ButtonVariant, string> = {
  primary: 'bg-primary text-white font-bold',
  outline: 'bg-transparent text-primary text-outline border border-primary',
}

export const DISABLED_CLASS = 'opacity-50 cursor-not-allowed'
export const ENABLED_CLASS = 'cursor-pointer'
