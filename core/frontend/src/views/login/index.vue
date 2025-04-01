<template>
	<div class="login-container">
		<div class="login-card">
			<div class="logo-container">
				<div class="logo">
					<svg viewBox="0 0 24 24" class="logo-icon">
						<path
							d="M16,12V4H17V2H7V4H8V12L6,14V16H11.2V22H12.8V16H18V14L16,12Z"
							fill="currentColor" />
					</svg>
				</div>
			</div>

			<h2 class="login-title">Billion Mail</h2>

			<n-form size="large">
				<n-form-item :show-label="false">
					<n-input v-model:value="form.username" placeholder="Username"></n-input>
				</n-form-item>
				<n-form-item :show-label="false">
					<n-input
						v-model:value="form.password"
						type="password"
						placeholder="Password"
						show-password-on="click">
					</n-input>
				</n-form-item>

				<n-button type="primary" size="large" block @click="handleLogin">Login</n-button>
			</n-form>
		</div>
	</div>
</template>

<script setup>
const form = reactive({
	username: '',
	password: '',
})

const handleLogin = () => {}
</script>

<style lang="scss" scoped>
.login-container {
	--primary-color: #00b9f5;
	--primary-hover: #0095c8;
	--text-dark: #333;
	--text-light: #666;
	--accent-green: #25cdb1;
	--accent-purple: #7e6ed5;
	--border-light: #ddd;
}

.login-container {
	display: flex;
	justify-content: center;
	align-items: center;
	min-height: 100vh;
	background: linear-gradient(135deg, #f0f3ff 0%, #e5f8f7 100%);
	position: relative;
	overflow: hidden;
	&::before {
		content: '';
		position: absolute;
		width: 40vw;
		height: 40vw;
		border-radius: 50%;
		background-color: var(--accent-green);
		top: -15vh;
		right: -15vw;
		z-index: 0;
	}
	&::after {
		content: '';
		position: absolute;
		width: 40vw;
		height: 40vw;
		border-radius: 50%;
		background-color: var(--accent-purple);
		bottom: -15vh;
		left: -15vw;
		z-index: 0;
	}
}

.login-card {
	background-color: white;
	padding: 52px 32px 62px;
	border-radius: 8px;
	box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
	width: 100%;
	max-width: 400px;
	z-index: 1;
}

.logo {
	display: flex;
	align-items: center;
	justify-content: center;
	width: 40px;
	height: 40px;
	background-color: var(--primary-color);
	border-radius: 5px;

	&-container {
		display: flex;
		align-items: center;
		justify-content: center;
		margin-bottom: 24px;
	}

	&-icon {
		width: 24px;
		height: 24px;
		color: white;
	}

	&-text {
		font-size: 20px;
		font-weight: 600;
		color: var(--text-dark);
	}
}

.login-title {
	margin-bottom: 32px;
	text-align: center;
	font-weight: 500;
	color: var(--text-dark);
}
</style>
