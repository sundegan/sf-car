"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
// app.ts
App({
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
        // 测试GRPC-Gateway服务
        wx.request({
            url: 'http://localhost:8080/helloworld/greeter/sayhello',
            method: 'GET',
            success: console.log,
            fail: console.error
        });
        // 展示本地存储能力
        const logs = wx.getStorageSync('logs') || [];
        logs.unshift(Date.now());
        wx.setStorageSync('logs', logs);
        // 登录
        wx.login({
            success: res => {
                console.log(res.code);
                // 发送 res.code 到后台换取 openId, sessionKey, unionId
            },
        });
    },
});
