// app.ts
App<IAppOption>({
  globalData: {
    // 定义用户信息全局变量
    userInfo: {
      avatarUrl: '',
      city: '',
      country: '',
      gender: 0,
      language: 'zh_CN',
      nickName: '',
      province: '',
    },
  },

  onLaunch() {
    // 展示本地存储能力
    const logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)

    // 登录
    wx.login({
      success: res => {
        console.log(res.code)
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
      },
    })
  },
})