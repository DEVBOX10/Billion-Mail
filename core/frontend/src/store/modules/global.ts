import { defineStore } from 'pinia'

export default defineStore('GlobalStore', () => {
	const isCollapse = ref(false)

	const setCollapse = () => {
		isCollapse.value = !isCollapse.value
	}

	return {
		isCollapse,
		setCollapse,
	}
})
