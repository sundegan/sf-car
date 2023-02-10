import { IAppOption } from "./appoption"
import { sfcar } from "./service/request"

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
    sfcar.login()

  },
})