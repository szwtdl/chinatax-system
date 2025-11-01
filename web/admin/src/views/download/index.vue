<template>
  <div class="download">
    <div v-if="isWechat" class="wechat">
      <img src="https://www.sumsemi.top/static/images/2.png" alt="下载页面">
    </div>
    <div v-else class="not-wechat">
      <div class="download-header">APP下载</div>
      <div class="download-body">
        <h1>立即下载手机客户端</h1>
        <ul>
          <li><a :href="androidUrl" target="_blank">安卓端下载</a></li>
          <li v-if="iosUrl"><a :href="iosUrl" target="_blank">苹果端下载</a></li>
        </ul>
      </div>
      <div class="download-footer">
        <p>版权所有: 广州技行信息技术科技有限公司</p>
        <p><a href="https://www.baidu.com" target="_blank">备案号：粤ICP备2022078965号-1</a></p>
      </div>
    </div>
  </div>
</template>
<script>
import { getBrowserInfo } from '@/utils/browser'
import { lastUpdate } from '@/api/public'
export default {
  name: 'Download',
  data() {
    return {
      browserInfo: {},
      isWechat: false,
      is_force_update: false,
      latest_version: '',
      min_supported_version: '',
      release_notes: '',
      androidUrl: 'https://www.baidu.com',
      iosUrl: 'https://www.baidu.com'
    }
  },
  mounted() {
    this.getAppUpdate()
  },
  created() {
    document.title = 'APP下载'
    this.browserInfo = getBrowserInfo()
    if (this.browserInfo.isWeChat === true && this.browserInfo.isAndroid === true) {
      this.isWechat = true
      document.body.style.backgroundColor = '#CDCDCD'
    }
  },
  methods: {
    getAppUpdate() {
      // 获取app版本信息
      lastUpdate().then(res => {
        this.is_force_update = res.data.is_force_update
        this.latest_version = res.data.latest_version
        this.min_supported_version = res.data.min_supported_version
        this.release_notes = res.data.release_notes
        this.androidUrl = res.data.update_android_url
        this.iosUrl = res.data.update_ios_url
      })
    }
  }
}
</script>
<style scoped lang="scss">
*{
  margin: 0;
  padding: 0;
  user-select: none;
}
.download-header {
  width: 100%;
  height: 60px;
  line-height: 60px;
  text-align: center;
  background-color: #2893ff;
  color: #fff;
  font-size: 18px;
  font-weight: bold;
}
.download-body{
  margin-top: 100px;
  h1 {
    text-align: center;
    font-size: 28px;
  }
  ul {
    list-style: none;
    display: flex;
    justify-content: center;
    margin-top: 30px;
    li {
      margin: 10px;
      a {
        display: block;
        width: 130px;
        height: 45px;
        line-height: 45px;
        text-align: center;
        background-color: #2893ff;
        border-radius: 5px;
        font-size: 16px;
        font-weight: bold;
        letter-spacing: 1px;
        color: #fff;
        text-decoration: none;
      }
      &:first-child a {
        background-color: #48d484;
      }
    }
  }
  .qr-code{
    width: 100%;
    text-align: center;
    margin-top: 30px;
    img{
      width: 200px;
      height: 200px;
    }
  }
}
.download-footer {
  width: 100%;
  text-align: center;
  color: black;
  font-size: 16px;
  position: fixed;
  bottom: 10px;
  p{
    margin: 5px;
  }
}
.wechat{
  width: 100%;
  height: 100%;
  text-align: center;
  img{
    width: 70%;
    padding-top: 20px;
  }
}
</style>
