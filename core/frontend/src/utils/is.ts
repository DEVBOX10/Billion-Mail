/**
 * @description: 判断值是否未某个类型
 */
export function is<T>(val: unknown, type: string): val is T {
	return toString.call(val) === `[object ${type}]`
}

/**
 * @description: 是否为对象
 * @param { unknown } val 参数
 * @returns { boolean }
 */
export function isObject<T>(val: unknown): val is T {
	return val !== null && is(val, 'object')
}
