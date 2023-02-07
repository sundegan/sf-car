import camelcaseKeys from "camelcase-keys"
import { IAppOption } from "./appoption"
import { auth } from "./service/proto_gen/auth/auth_pb"

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
        // 测试GRPC-Gateway服务
        wx.request({
          url: 'http://localhost:8080/v1/auth/login',
          method: 'POST',
          data: {
            code: res.code,
          } as auth.v1.ILoginRequest,
          success: res => {
            const loginResp: auth.v1.ILoginResponse = 
              auth.v1.LoginResponse.fromObject(
                camelcaseKeys(res.data as object, {deep: true}),
              )
              console.log(loginResp)
          },
          fail: console.error,
        })
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
      },
    })
  },
})