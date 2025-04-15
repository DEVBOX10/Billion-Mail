import { useClipboard } from '@vueuse/core'
import { Message } from './'
import { isNumber } from './is'

// Server prefix
export const apiUrlPrefix: string = import.meta.env.API_URL_PREFIX

/**
 * @description Byte conversion, ends at specified unit
 * @param { number } bytes Number of bytes
 * @param { boolean } isUnit Whether to display unit
 * @param { number } fixed Decimal point position
 * @param { string } endUnit End unit
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
 * @description Copy text
 * @param {string} value Text to be copied
 */
export const copyText = async (value: string) => {
	const { copy } = useClipboard({
		source: value,
	})

	try {
		copy(value)
		Message.success('Copy successful')
	} catch {
		Message.error('Copy failed')
	}
}
