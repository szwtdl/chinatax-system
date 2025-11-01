export function getBrowserInfo() {
  const ua = navigator.userAgent.toLowerCase()
  return {
    isAndroid: /android/.test(ua),
    isIOS: /iphone|ipad|ipod/.test(ua),
    isWeChat: /micromessenger/.test(ua),
    isMobile: /android|webos|iphone|ipad|ipod|blackberry|iemobile|opera mini/.test(ua),
    isDesktop: !/android|webos|iphone|ipad|ipod|blackberry|iemobile|opera mini/.test(ua)
  }
}
