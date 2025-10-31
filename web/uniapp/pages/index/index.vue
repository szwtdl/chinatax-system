<template>
	<view class="container">
		<view class="no-login">
			<view class="no-login-body">
				<text class="no-login-body-title">
					伽玛账-用手机自己记账报税
				</text>
				<text class="no-login-body-sub-title">
					每月最快一分钟完成
				</text>
			</view>
			<view class="no-login-btn" @click="() => go('/pages/login/login')">
				<text class="no-login-btn-txt">注册/登录</text>
				<uni-icons class="login-check-more" type="right" color="#3C84F9" size="18"></uni-icons>
			</view>
		</view>
		<view class="not-kefu">
			<view class="not-kefu-avatr">
				<image src="/static/logo.png" mode="aspectFit"></image>
			</view>
			<view class="not-kefu-body">
				<view class="not-kefu-card">
					<view class="not-kefu-card-header">
						<text class="not-kefu-card-header-nickname">专属客服</text>
						<text class="not-kefu-card-header-live">在线</text>
					</view>
					<text class="not-kefu-card-desc">
						使用后即享1对1专属客服
					</text>
				</view>
				<view class="not-kefu-btn">联系TA</view>
			</view>
		</view>
		<view class="register-tip">
			<text class="register-tip-title">2025年度注册资金实缴登记</text>
			<text class="register-tip-btn">立即登记 ></text>
		</view>
	</view>
</template>

<script>
	import { navigateTo,redirectTo } from '@/utils/navigate.uts'
	export default {
		data() {
			return {
				
			}
		},
		onLoad() {

		},
		methods: {
			go(url) {
			  navigateTo(url)
			}
		}
	}
</script>


<style scoped lang="scss">
	.container {
		width: 100vw;
		height: 100vh;
		background-color: #F3F8FC;
		.no-login{
			display: block;
			padding: 30rpx;
			margin-bottom: 30rpx;
			.no-login-body{
				display: inline-block;
				.no-login-body-title{
					display: block;
					font-size: $uni-font-size-base;
				}
				.no-login-body-sub-title{
					display: block;
					font-size: $uni-font-size-sm;
					color: #9B9D9E;
				}
			}
			.no-login-btn{
				float: right;
				display: inline-block;
				background-color: #E3F1FE;
				padding: 20rpx 40rpx;
				border-radius: 50rpx;
				.no-login-btn-txt{
					color: #5C97FA;
					font-size: $uni-font-size-base;
				}
			}
		}
		.not-kefu{
			background-color: white;
			margin: 20rpx 30rpx;
			border-radius: 30rpx;
			padding: 30rpx;
			display: flex;
			align-items: center;
			.not-kefu-avatr{
				width: 90rpx;
				height: 90rpx;
				margin-right: 20rpx; 
				display: inline-block;
				image{
					width: 100%;
					height: 100%;
					border-radius: 100rpx;
					overflow: hidden;
				}
			}
			.not-kefu-body {
				flex: 1;
				display: flex;
				flex-direction: row; /* 横向排列 */
				align-items: center; /* 垂直居中 */
				justify-content: space-between; /* 左右分布 */
			
				.not-kefu-card {
					flex: 1;
					display: block;
			
					.not-kefu-card-header {
						font-size: $uni-font-size-title;
					}
					.not-kefu-card-desc {
						font-size: $uni-font-size-sm;
						color: #838383;
					}
				}
			
				.not-kefu-btn {
					text-align: center;
					padding: 15rpx 35rpx;
					background-color: #E3F1FE;
					color: #609BFA;
					border-radius: 999rpx;
					font-size: $uni-font-size-base;
				}
			}

		}
		.register-tip{
			border: 1px solid #F79F41;
			margin: 30rpx;
			font-size: $uni-font-size-base;
			padding: 10rpx 20rpx;
			border-radius: 10rpx;
			display: block;
			font-weight: bold;
			.register-tip-title{
				display: inline-block;
				color: #F79F41;
			}
			.register-tip-btn{
				float: right;
				color: #F79F41;
			}
		}
	}
</style>
