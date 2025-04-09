import { useClipboard } from '@vueuse/core'
import { Message } from './'
import { isNumber } from './is'

// 服务器前缀
export const apiUrlPrefix: string = import.meta.env.API_URL_PREFIX

/**
 * @description 字节转换，到指定单位结束
 * @param { number } bytes 字节数
 * @param { boolean } isUnit 是否显示单位
 * @param { number } fixed 小数点位置
 * @param { string } endUnit 结束单位
 * @returns { string }
 */
export const getByteUnit = (
	bytes: number = 0,
	isUnit: boolean = true,
	fixed: number = 2,
	endUnit: string = ''
): string => {
	let newBytes = isNumber(bytes) ? bytes : Number(bytes)
	const c = 1024
	const units = [' B', ' KB', ' MB', ' GB', ' TB', ' PB', ' EB', ' ZB', ' YB', ' BB', ' NB', ' DB']
	for (let i = 0; i < units.length; i++) {
		const unit = units[i]
		const showValue = fixed === 0 ? Math.round(newBytes) : newBytes.toFixed(fixed)
		const result = i === 0 ? newBytes.toFixed(fixed) : showValue
		if (endUnit) {
			if (unit.trim() === endUnit.trim()) {
				return isUnit ? result + unit : `${result}`
			}
		} else if (newBytes < c) {
			return isUnit ? result + unit : `${result}`
		}
		newBytes /= c
	}
	return ''
}

/**
 * @description 复制文本
 * @param {string} value 复制的文本
 */
export const copyText = async (value: string) => {
	const { copy } = useClipboard({
		source: value,
	})

	try {
		copy(value)
		Message.success('复制成功')
	} catch {
		Message.error('复制失败')
	}
}
