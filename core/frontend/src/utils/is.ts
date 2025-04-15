/**
 * @description: Check if value is of a specific type
 */
export function is<T>(val: unknown, type: string): val is T {
	return Object.prototype.toString.call(val) === `[object ${type}]`
}

/**
 * @description: Check if value is an object
 * @param { unknown } val Parameter
 * @returns { boolean }
 */
export function isObject<T>(val: unknown): val is T {
	return val !== null && is(val, 'Object')
}

/**
 * @description Check if value is a number
 * @param { unknown } val Parameter
 * @returns { boolean }
 */
export const isNumber = (val: unknown): val is number => {
	return is(val, 'Number')
}

/**
 * @description Check if value is a string
 * @param { unknown } val Parameter
 * @returns { boolean }
 */
export const isString = (val: unknown): val is string => {
	return is(val, 'String')
}

/**
 * @description Check if value is a function
 * @param { unknown } val Parameter
 * @returns { boolean }
 */
export const isFunction = <T = Function>(val: unknown): val is T => {
	return is(val, 'Function')
}

/**
 * @description Check if value is undefined
 * @param { unknown } val Parameter
 * @returns { boolean }
 */
export const isUndefined = (val: unknown): val is undefined => {
	return typeof val === 'undefined'
}

/**
 * @description Check if value is an array
 * @param { unknown } val Parameter
 * @returns { boolean }
 */
export const isArray = <T = unknown>(val: unknown): val is T[] => {
	return Array.isArray(val)
}

/**
 * Check if value is a date
 * @param { unknown } val Parameter
 * @returns { boolean }
 */
export const isDate = (val: unknown): val is Date => {
	return is(val, 'Date')
}
