import { ProfileService } from "../../service/ profile"
import { rental } from "../../service/proto_gen/rental/rental_pb"
import { padString } from "../../utils/format"
import { routing } from "../../utils/routing"

function formatDate(millis: number) {
  const dt = new Date(millis)
  const y = dt.getFullYear()
  const m = dt.getMonth() + 1
  const d = dt.getDate()
  return `${padString(y)}-${padString(m)}-${padString(d)}`
}

// pages/register/register.ts
Page({
  redirectURL: '',
  profileRefersher: 0,

  data: {
    licNo: '',
    name: '',
    genderIndex: 0,
    genders: ['未知', '男', '女'],
    birthDate: '1990-01-01',
    licImgURL: '',
    state: rental.v1.IdentityStatus[rental.v1.IdentityStatus.UNSUBMITTED],
},
  
  renderProfile(p: rental.v1.IProfile) {
    this.setData({
      licNo: p.identity?.licNumber || '',
      name: p.identity?.name || '',
      genderIndex: p.identity?.gender || 0,
      birthDate: formatDate(p.identity?.birthDateMillis || 0),
      state: rental.v1.IdentityStatus[p.identityStatus || 0],
    })
  },

  onLoad(opt: Record<'redirect', string>) {
    const o: routing.RegisterOpts = opt
    if (o.redirect) {
      this.redirectURL = decodeURIComponent(o.redirect)
    }
    ProfileService.getProfile().then(p => this.renderProfile(p))
  },

  // 上传驾驶证照片
  onUploadLic() {
    wx.chooseMedia({
      success: res => {
        this.setData({
            licImgURL: res.tempFiles[0].tempFilePath,
        })
        // TODO: 上传图片到服务器存储
        // 服务器返回name,gender,licID等信息,自动填写这些信息
        setTimeout(() => {
          this.setData({
            name: '张三',
            genderIndex: 1,
            licNo: '123456',
            phone: '123456',
            birthDate: '1990-01-01',
          })
        }, 1000);
      }
    })
  },

  onGenderChange(e: any) {
    this.setData({
        genderIndex: parseInt(e.detail.value),
    })
  },

  // 日期选择处理函数
  onBirthDateChange(e: any) {
    this.setData({
      birthDate: e.detail.value,
    })
  },

  // 提交审核处理函数
  onSubmit() {
    ProfileService.submitProfile({
      licNumber: this.data.licNo,
      name: this.data.name,
      gender: this.data.genderIndex,
      birthDateMillis: Date.parse(this.data.birthDate),
    }).then(p => {
      this.renderProfile(p)           // 重填身份信息
      this.scheduleProfileRefresher() // 刷新验证页面
    })
  },

  // 每秒刷新页面,审核完成停止刷新
  scheduleProfileRefresher() {
    this.profileRefersher = setInterval(() => {
      ProfileService.getProfile().then(p => {
        this.renderProfile(p)
        if (p.identityStatus !== rental.v1.IdentityStatus.PENDING) {
          this.clearProfileRefresher()
        }
        if (p.identityStatus === rental.v1.IdentityStatus.VERIFIED) {
          this.onLicVerified()
        }
      })
      
    }, 1000)
  },

  clearProfileRefresher() {
    if (this.profileRefersher) {
      clearInterval(this.profileRefersher)
      this.profileRefersher = 0
    }
  },

  onUnload(){
    this.clearProfileRefresher() 
  }, 

  // 重新审核处理函数
  onResubmit() {
    ProfileService.clearProfile().then(p => this.renderProfile(p))
  },

  // 服务器审核结束处理函数
  onLicVerified() {
    // 当从扫码页面跳转过来注册，当注册完后跳转到开锁页面
    // 当从我的行程页面跳转过来进行注册，则不会跳转到开锁页面
    if (this.redirectURL) {
      wx.redirectTo({
          url: this.redirectURL,
      })
    }
  },
})