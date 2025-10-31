<script>
	import { isLoggedIn } from '@/utils/auth'
	export default {
		onLaunch: function() {
			console.log('App Launch')
		},
		onShow: function() {
			console.log('App Show')
			// if (!isLoggedIn()) {
			// 	uni.redirectTo({
			// 		url: '/pages/login/login' // 替换为你的登录页路径
			// 	})
			// }
		},
		onHide: function() {
			console.log('App Hide')
		}
	}
</script>

<style lang="scss">
	@import "@/uni.scss";
	/*每个页面公共css */
</style>
