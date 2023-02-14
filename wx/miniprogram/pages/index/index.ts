// index.ts
import { ProfileService } from "../../service/ profile"
import { rental } from "../../service/proto_gen/rental/rental_pb"
import { TripService } from "../../service/trip"
import { routing } from "../../utils/routing"

Page({
  // 判断页面是否在前台显示,当页面不在前台显示时,不更新汽车位置
  isPageShowing: false, 
  
  // 初始经纬度数据
  location: {
    latitude: 23.130324,
    longitude: 113.348478, 
  },

  data: {
    avatarUrl: '',  // 用户头像链接

    // map组件设置数据
    setting: {
      skew: 0,
      rotate: 0,
      showLocation: true,
      showScale: true,
      subKey: '',
      layerStyle: -1,
      enableZoom: true,
      enableRotate: false,
      showCompass: false,
      enable3D: false,
      enableOverlooking: false,
      enableSatellite: false,
      enableTraffic: false,
    },

    // 缩放等级
    scale: 15,

    // 汽车图标数据
    markers: [
      {
        iconPath: "/resources/car.png",
        id: 0,
        // 暨南大学坐标
        latitude: 23.130324,
        longitude: 113.348478,
        width: 50,
        height: 50,
      },
    ]
  },

  // 点击定位图标的处理函数,返回当前位置的经纬度
  myLocationTap() {
    wx.getLocation({
      type: 'gcj02',
      success: res => {
        this.setData({
          location: {
            latitude: res.latitude,
            longitude: res.longitude,
          },
        })
      },
      fail: () => {
        wx.showToast({
          icon: 'none',
          title: '点击右上角设置授权位置信息',
        })
      }
    }) 
  },

  // 每次加载首页时，获取当前位置信息以及从本地缓存获取头像链接
  async onLoad() {
    wx.getLocation({
      type: 'gcj02',
      success: res => {
        this.setData({
          location: {
            latitude: res.latitude,
            longitude: res.longitude,
          },
        })
      },
      fail: () => {
        wx.showToast({
          icon: 'none',
          title: '点击右上角设置授权位置信息',
        })
      }
    }) 
    
    const avatarUrl = wx.getStorageSync('avatarUrl')
    this.setData({
      avatarUrl: avatarUrl,
    })
  },

  onShow() {
    this.isPageShowing = true
  },

  onHide() {
    this.isPageShowing = false 
  },

  // 更新汽车位置
  moveCars() {
    // 获取map对象
    const map = wx.createMapContext('map')

    const moveCar = () => {
      this.location.latitude += 0.1
      this.location.longitude += 0.1
      map.translateMarker({
        destination: {
          latitude: this.location.latitude,
          longitude: this.location.longitude,
        },
        markerId: 0,
        autoRotate: false,
        rotate: 0,
        duration: 5000, // 移动五秒钟
        // 五秒结束后继续移动
        animationEnd: () => {
          if (this.isPageShowing) {
            moveCar()
          }
        }
      })
    }
    moveCar()
  },

  // 扫码租车处理函数
  async onScanQRcodes() {
    const trips = await TripService.getTrips(rental.v1.TripStatus.IN_PROGRESS)
    if ((trips.trips?.length || 0) > 0) {
      wx.showModal({
        title: '行程中',
        content: '当前有正在进行的行程，将跳转到行程页面',
        success: () => {
          wx.navigateTo({
            url: routing.drving({
              trip_id: trips.trips![0].id!,
            }),
          })
        },
        fail: () => {},
      })

      return
    }
    wx.scanCode({
      success: async () => {
        // TODO: 从扫描结果获取汽车ID
        const carID = 'car123'
        const lockURL = routing.lock({
          car_id: carID,
        })
        // 如果身份未认证,跳转到身份认证页
        // 如果已认证,跳转到开锁页
        const profile = await ProfileService.getProfile()
        if (profile.identityStatus === rental.v1.IdentityStatus.VERIFIED) {
          wx.navigateTo({
            url: lockURL,
          })
        } else {
          wx.showModal({
            title: '未进行身份认证',
            content: '需要先经过驾驶资格认证',
            success: () => {
              wx.navigateTo({
                url: routing.register({
                  redirectURL: lockURL,
                })
                // 等价于:url: `/pages/register/register?redirect=${encodeURIComponent(redirectURL)}`,
              })
            },
            fail: () => {},
          })
        }
      },
      fail: () => {
        wx.showToast({
          icon: 'none',
          title: '扫码失败!',
        })
      }
    })
  },

  // 点击首页用户头像跳转到我的行程列表
  onMyTripsTap() {
    wx.navigateTo({
      //url: '/pages/mytrips/mytrips'
      url: routing.mytrips(),
    })
  },

})