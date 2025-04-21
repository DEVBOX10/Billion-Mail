import { isNumber } from './is'

// Server prefix
export const apiUrlPrefix: string = import.meta.env.API_URL_PREFIX

/**
 * Byte conversion with specified unit
 * @param bytes Number of bytes
 * @param isUnit Whether to display unit
 * @param fixed Decimal point position
 * @param endUnit End unit
 * @returns Formatted string with unit
 */
export const getByteUnit = (
	bytes: number = 0,
	isUnit: boolean = true,
	fixed: number = 2,
	endUnit: string = ''
): string => {
	if (!isNumber(bytes)) {
		bytes = Number(bytes)
		if (isNaN(bytes)) return ''
	}

	const units = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB', 'BB', 'NB', 'DB']
	const base = 1024
	let value = bytes

	for (let i = 0; i < units.length; i++) {
		const unit = units[i]
		const formattedValue = fixed === 0 ? Math.round(value).toString() : value.toFixed(fixed)

		if (endUnit && unit === endUnit.trim()) {
			return isUnit ? `${formattedValue} ${unit}` : formattedValue
		}

		if (value < base) {
			return isUnit ? `${formattedValue} ${unit}` : formattedValue
		}

		value /= base
	}

	return '--'
}
