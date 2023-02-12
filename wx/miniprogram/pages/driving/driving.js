"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const trip_1 = require("../../service/trip");
const format_1 = require("../../utils/format");
const routing_1 = require("../../utils/routing");
const centPerSec = 1; // 每秒一分钱
function durationStr(sec) {
    const dur = (0, format_1.formatDuration)(sec);
    return `${dur.hh}:${dur.mm}:${dur.ss}`;
}
Page({
    timer: undefined,
    data: {
        location: {
            latitude: 23.130324,
            longitude: 113.348478,
        },
        scale: 14,
        elapsed: '00:00:00',
        fee: '0.00',
    },
    // 页面打开更新位置开始计时
    // 以及从开锁页面获取行程ID
    onLoad(opt) {
        const o = opt;
        console.log('current trip_id', o.trip_id);
        trip_1.TripService.GetTrip(o.trip_id).then(console.log);
        this.setupLocationUpdator();
        this.setupTimer();
    },
    // 页面关闭停止位置更新和清空计时
    onUnload() {
        wx.stopLocationUpdate();
        if (this.timer) {
            clearInterval(this.timer);
        }
    },
    // 实时位置更新
    setupLocationUpdator() {
        // 开启小程序进入前后台时均接收位置消息，需引导用户开启授权。
        // 授权以后，小程序在运行中或进入后台均可接受位置消息变化。
        // wx.authorize提前向用户发起授权请求。调用后会立刻弹窗询问用户是否同意授权小程序使用某项功能或获取用户的某些数据，
        // 但不会实际调用对应接口。如果用户之前已经同意授权，则不会出现弹窗，直接返回成功。
        wx.authorize({
            scope: 'scope.userLocationBackground',
            success: () => {
                wx.startLocationUpdateBackground({
                    fail: console.error,
                });
                wx.onLocationChange(res => {
                    console.log('current location:', res);
                    this.setData({
                        location: {
                            latitude: res.latitude,
                            longitude: res.longitude,
                        },
                    });
                });
            },
            // 获取后台更新位置权限失败提示授权
            fail: () => {
                wx.showToast({
                    icon: 'none',
                    title: '点击右上角授权位置信息',
                });
            },
        });
    },
    // 计时
    setupTimer() {
        let elapsedSec = 0; // 租车时间
        let cents = 0; // 租车费用
        this.timer = setInterval(() => {
            elapsedSec++; // 时间增加
            cents += centPerSec;
            // 把秒转化为小时分钟秒后显示
            this.setData({
                elapsed: durationStr(elapsedSec),
                fee: (0, format_1.formatFee)(cents),
            });
        }, 1000);
    },
    // 结束行程
    onEndTripTap() {
        // TODO:支付模块接入
        wx.redirectTo({
            //url: '/pages/index/index',
            url: routing_1.routing.mytrips(),
        });
    },
});
