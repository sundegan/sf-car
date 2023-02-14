import { ProfileService } from "../../service/ profile"
import { rental } from "../../service/proto_gen/rental/rental_pb"
import { sfcar } from "../../service/request"
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
    this.renderIdentity(p.identity!)
    this.setData({
      state: rental.v1.IdentityStatus[p.identityStatus || 0],
    }) 
  },

  renderIdentity(i?: rental.v1.IIdentity) {
    this.setData({
      licNo: i?.licNumber || '',
      name: i?.name || '',
      genderIndex: i?.gender || 0,
      birthDate: formatDate(i?.birthDateMillis || 0),
    }) 
  },

  onLoad(opt: Record<'redirect', string>) {
    const o: routing.RegisterOpts = opt
    if (o.redirect) {
      this.redirectURL = decodeURIComponent(o.redirect)
    }
    // 获取身份认证信息和驾照图片信息
    ProfileService.getProfile().then(p => this.renderProfile(p))
    ProfileService.getProfilePhoto().then(p => {
      console.log(p.url)
      this.setData({
        licImgURL: p.url || '',
      })
    })
  },

  // 上传驾驶证照片
  onUploadLic() {
    wx.chooseMedia({
      success: async res => {
        if (res.tempFiles.length === 0) {
          return 
        }
        this.setData({
          licImgURL: res.tempFiles[0].tempFilePath,
        })
        // 获取图片上传地址（预签名URL）
        const photoRes = await ProfileService.createProfilePhoto()
        console.log(photoRes)
        // 上传图片(如果上传地址不为空)
        if (!photoRes.uploadUrl) {
          return
        }
        await sfcar.uploadfile({
          localPath: res.tempFiles[0].tempFilePath,
          url: photoRes.uploadUrl,
        })
        // 上传完成通知服务器上传成功,服务器返回身份信息
        const identity = await ProfileService.completeProfilePhoto()
        // 保存身份信息到本地
        this.renderIdentity(identity)
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

  // 重新审核按钮处理函数
  onResubmit() {
    // 重新提交审核时清理之前上传的图片
    ProfileService.clearProfile().then(p => this.renderProfile(p))
    ProfileService.clearProfilePhoto().then(() => {
      this.setData({
          licImgURL: '',
      })
    })
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