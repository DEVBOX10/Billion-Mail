import { useClipboard } from '@vueuse/core'
import { Message } from '@/utils/message'

/**
 * @description 复制文本
 * @param {string} value 复制的文本
 */
export const copyText = async (value: string) => {
	const { copy, isSupported } = useClipboard({ legacy: true })
	if (isSupported.value) {
		copy(value)
		Message.success('复制成功')
	} else {
		Message.error('复制失败')
	}
}
