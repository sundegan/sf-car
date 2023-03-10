"use strict";
// pages/lock/lock.ts
Object.defineProperty(exports, "__esModule", { value: true });
const trip_1 = require("../../service/trip");
const routing_1 = require("../../utils/routing");
// 全局变量
const shareLocationKey = "shareLocation";
Page({
    carID: '',
    data: {
        avatarUrl: '',
        shareLocation: false, // 是否共享我的位置
    },
    // 每次登录从本地缓存读取shareLocationKey
    // 从本地缓存读取avatarUrl
    // 从index页面获取扫码的carID
    async onLoad(opt) {
        const o = opt;
        // 从扫码页面获取传递过来的car_id参数
        this.carID = o.car_id;
        const avatarUrl = wx.getStorageSync('avatarUrl');
        this.setData({
            avatarUrl: avatarUrl,
            shareLocation: wx.getStorageSync(shareLocationKey) || false,
        });
    },
    // 获取用户头像用于在地图上向别人实时分享我的位置
    onChooseAvatar(e) {
        this.setData({
            avatarUrl: e.detail.avatarUrl,
        }),
            // 将获取的avatarUrl保存到本地缓存中
            wx.setStorageSync('avatarUrl', e.detail.avatarUrl);
    },
    // 打开/关闭位置共享
    onShareLocation(e) {
        this.data.shareLocation = e.detail.value;
        // 将打开/关闭位置共享的设置保存到本地缓存，避免每次都需要重新设置
        wx.setStorageSync(shareLocationKey, this.data.shareLocation);
    },
    // 点击开锁按钮处理函数
    onUnlockTap() {
        // 获得当前位置信息
        wx.getLocation({
            type: 'gcj02',
            // 模拟向服务器发送位置位置信息、头像链接、汽车ID
            success: async (res) => {
                console.log('starting a trip', {
                    location: {
                        latitude: res.latitude,
                        longitude: res.longitude,
                    },
                    // TODO: 需要双向绑定
                    avatarUrl: this.data.shareLocation
                        ? this.data.avatarUrl : '',
                });
                // 向后台发送请求创建行程
                if (!this.carID) {
                    console.error('no car_id specified, please bring car_id parameter n the lock page');
                    return;
                }
                const trip = await trip_1.TripService.CreateTrip({
                    carId: this.carID,
                    start: {
                        latitude: res.latitude,
                        longitude: res.longitude,
                    }
                });
                if (!trip.id) {
                    console.error('no tripID in response', trip);
                    return;
                }
                // 显示开锁中提示框
                wx.showLoading({
                    title: '开锁中',
                    mask: true,
                });
                // 模拟两秒钟后开锁完成，跳转到行程页面
                setTimeout(() => {
                    wx.redirectTo({
                        // url: `/pages/driving/driving?trip_id=${tripID}`
                        url: routing_1.routing.drving({
                            trip_id: trip.id,
                        }),
                        // 不管成功或者失败都要取消开锁中提示
                        complete: () => {
                            wx.hideLoading();
                        }
                    });
                }, 2000);
            },
            fail: () => {
                wx.showToast({
                    icon: 'none',
                    title: '点击右上角授权位置信息',
                });
            },
        });
    },
});
