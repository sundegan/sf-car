import { routing } from "../../utils/routing"

// pages/register/register.ts
Page({
  redirectURL: '',
  data: {
    name: '',
    gender: '',
    licID: '',
    phone: '',
    date: '2023-01-25', 
    licImgURL: '',
    auditState: 'UNSUBMITTED' as 'UNSUBMITTED' | 'PENDING' | 'VERIFIED',
  },
  
  // 从扫码页面获取汽车ID
  onLoad(opt: Record<'redirect', string>) {
    const o: routing.RegisterOpts = opt
    if (o.redirect) {
      this.redirectURL = decodeURIComponent(o.redirect)
    }
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
            gender: '男',
            licID: '123456',
            phone: '123456',
            date: '1998-9-8',
          })
        }, 1000);
      }
    })
  },

  // 日期选择处理函数
  onDateChange(e: any) {
    this.setData({
      date: e.detail.value,
    })
  },

  // 提交审核处理函数
  onSubmit() {
    // TODO: 上传表单到服务器
    this.setData({
      auditState: 'PENDING',
    })
    // 模拟服务器审核,3秒审核结束
    setTimeout(() => {
      this.onLicVerified()
    }, 3000);
  },

  // 重新审核处理函数
  onResubmit() {
    this.setData({
      auditState: 'UNSUBMITTED',  // 清空状态
      licImgURL: '',
    })
  },

  // 服务器审核结束处理函数
  onLicVerified() {
    this.setData({
      auditState: 'VERIFIED',
    })
    // 当从扫码页面跳转过来注册，当注册完后跳转到开锁页面
    // 当从我的行程页面跳转过来进行注册，则不会跳转到开锁页面
    if (this.redirectURL) {
      wx.redirectTo({
          url: this.redirectURL,
      })
    }
  },
})