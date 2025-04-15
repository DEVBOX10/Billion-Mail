import { toNumber } from 'lodash-es'

/**
 * @description Convert value to number, return 0 if not a number
 * @param val Value
 * @returns
 */
export const getNumber = (val: unknown) => {
	return !Number.isNaN(toNumber(val)) ? toNumber(val) : 0
}
