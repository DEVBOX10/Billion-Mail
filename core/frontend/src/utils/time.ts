import { format, startOfDay, endOfDay } from 'date-fns'
import { isString, isDate, isNumber, getNumber } from '@/utils'
import { toNumber } from 'lodash-es'

export type Time = Date | number | string

/**
 * @description Format time
 * @param time Timestamp / Date / Time string
 * @param fmt Time format, optional
 * @returns { string } Formatted time
 */
export const formatTime = (time?: Time, fmt = 'yyyy-MM-dd HH:mm:ss'): string => {
	if (isString(time)) {
		time = isNaN(toNumber(time)) ? new Date(time).getTime() : getNumber(time)
	}

	let val = 0
	if (isDate(time)) {
		val = time.getTime()
	} else if (isNumber(time)) {
		time = Math.round(time)
		const str = time.toString()
		if (str.length === 13) val = time
		if (str.length === 10) val = time * 1000
	}

	if (val === 0) return '--'

	return format(val, fmt)
}

/**
 * @description Get time range
 * @param { Date } date Date
 * @returns { number[] }
 */
export const getDayTimeRange = (date = new Date()): [number, number] => {
	const start = startOfDay(date)
	const end = endOfDay(date)
	return [start.getTime(), end.getTime()]
}
