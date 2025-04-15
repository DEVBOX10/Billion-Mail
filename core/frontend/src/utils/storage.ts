import { isObject } from './is'

/**
 * Get data from localStorage
 * @param key Key name
 * @param defaultVal Default value
 * @returns Stored value or default value
 */
export const getLocalStorage = <T = string>(key: string, defaultVal: T = '' as T): T => {
	try {
		const val = localStorage.getItem(key)
		if (val === null) return defaultVal

		if (isObject<T>(defaultVal)) {
			return JSON.parse(val) as T
		}
		return val as T
	} catch {
		// console.error(`Failed to get localStorage item: ${key}`, error)
		return defaultVal
	}
}

/**
 * Store data in localStorage
 * @param key Key name
 * @param value Value to store
 */
export const setLocalStorage = <T>(key: string, value: T): void => {
	const val = isObject(value) ? JSON.stringify(value) : String(value)
	localStorage.setItem(key, val)
}

/**
 * Delete one or multiple data from localStorage
 * @param keys Key name or array of key names
 */
export const delLocalStorage = (keys: string | string[]): void => {
	const keyArray = Array.isArray(keys) ? keys : [keys]
	keyArray.forEach(key => {
		localStorage.removeItem(key)
	})
}

/**
 * Set cookie
 * @param key Key name
 * @param value Value
 * @param days Expiration days, default 30 days
 * @param path Cookie path, default root path
 */
export const setCookie = (key: string, value: string, days = 30, path = '/'): void => {
	const date = new Date()
	date.setTime(date.getTime() + days * 24 * 60 * 60 * 1000)
	const expires = `expires=${date.toUTCString()}`
	const cookiePath = `path=${path}`
	document.cookie = `${key}=${encodeURIComponent(value)};${expires};${cookiePath}`
}

/**
 * Get cookie
 * @param key Key name
 * @returns Cookie value or empty string
 */
export const getCookie = (key: string): string => {
	try {
		const name = `${key}=`
		const decodedCookie = decodeURIComponent(document.cookie)
		const cookieArray = decodedCookie.split(';')

		for (const cookie of cookieArray) {
			const trimmedCookie = cookie.trim()
			if (trimmedCookie.startsWith(name)) {
				return trimmedCookie.substring(name.length)
			}
		}
		return ''
	} catch {
		return ''
	}
}

/**
 * Delete cookie
 * @param key Key name
 * @param path Cookie path, default root path
 */
export const delCookie = (key: string, path = '/'): void => {
	setCookie(key, '', -1, path)
}
