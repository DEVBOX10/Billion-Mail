/**
 * @description: 判断值是否未某个类型
 */
export function is<T>(val: unknown, type: string): val is T {
	return Object.prototype.toString.call(val) === `[object ${type}]`
}

/**
 * @description: 是否为对象
 * @param { unknown } val 参数
 * @returns { boolean }
 */
export function isObject<T>(val: unknown): val is T {
	return val !== null && is(val, 'Object')
}

/**
 * @description 是否为字符串
 * @param { unknown } val 参数
 * @returns { boolean }
 */
export const isString = (val: unknown): val is string => {
	return is(val, 'String')
}

/**
 * @description 是否为函数
 * @param { unknown } val 参数
 * @returns { boolean }
 */
export const isFunction = <T = Function>(val: unknown): val is T => {
	return is(val, 'Function')
}

/**
 * @description 是否未定义
 * @param { unknown } val 参数
 * @returns { boolean }
 */
export const isUndefined = (val: unknown): val is undefined => {
	return typeof val === 'undefined'
}

/**
 * @description 是否为数组
 * @param { unknown } val 参数
 * @returns { boolean }
 */
export const isArray = <T = unknown>(val: unknown): val is T[] => {
	return Array.isArray(val)
}
