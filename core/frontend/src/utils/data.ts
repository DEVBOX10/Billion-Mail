import { toNumber } from 'lodash-es'

/**
 * @description 转化数字，将非数字转化为0
 * @param val 值
 * @returns
 */
export const getNumber = (val: unknown) => {
	return !Number.isNaN(toNumber(val)) ? toNumber(val) : 0
}
