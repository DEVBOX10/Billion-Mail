import { format, startOfDay, endOfDay } from 'date-fns'
import { toNumber } from 'lodash-es'
import { isString, isDate, isNumber, getNumber } from '@/utils'

export type Time = Date | number | string

/**
 * @description Format time
 * @param time Timestamp / Date / Time string
 * @param fmt Time format, optional
 * @returns { string } Formatted time
 * @throws { Error } When time is invalid
 */
export const formatTime = (time?: Time, fmt = 'yyyy-MM-dd HH:mm:ss'): string => {
	if (!time) return '--'

	let timestamp: number

	try {
		if (isString(time)) {
			timestamp = isNaN(toNumber(time)) ? new Date(time).getTime() : getNumber(time)
		} else if (isDate(time)) {
			timestamp = time.getTime()
		} else if (isNumber(time)) {
			const str = time.toString()
			timestamp = str.length === 10 ? time * 1000 : time
		} else {
			return '--'
		}

		return format(timestamp, fmt)
	} catch {
		// console.error('Time formatting error:', error)
		return '--'
	}
}

/**
 * @description Get time range for a given date
 * @param { Date } date Date object, defaults to current date
 * @returns { [number, number] } Tuple containing start and end timestamps
 */
export const getDayTimeRange = (date: Date = new Date()): [number, number] => {
	const start = startOfDay(date)
	const end = endOfDay(date)
	return [start.getTime(), end.getTime()]
}
