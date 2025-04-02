<template>
	<n-config-provider abstract inline-theme-disabled :theme-overrides="themeOverrides">
		<slot></slot>
	</n-config-provider>
</template>

<script lang="ts" setup>
const themeOverrides = {
	common: {
		lineHeight: 'normal',
		fontSize: '12px',
		fontSizeSmall: '12px',
		fontSizeMedium: '12px',
		fontSizeLarge: '14px',
		borderRadius: '4px',
		primaryColor: '#20A53A',
		primaryColorHover: '#1D9534',
		successColor: '#20A53A',
		successColorHover: '#1D9534',
	},
	Layout: {
		color: '#f4f7fa',
		textColor: '#333',
		headerColor: 'rgba(244, 247, 250, 0.7)',
		headerTextColor: '#5b6b79',
	},
	Menu: {
		fontSize: '14px',
	},
	Form: {
		feedbackHeightMedium: '20px',
		feedbackHeightLarge: '22px',
		feedbackFontSizeMedium: '12px',
		feedbackFontSizeLarge: '12px',
		feedbackPadding: '2px 0 0',
		labelFontSizeLeftMedium: '12px',
		labelPaddingHorizontal: '0 20px 0 0',
		labelFontSizeTopMedium: '12px',
	},
}
</script>
