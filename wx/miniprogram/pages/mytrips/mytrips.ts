import { routing } from "../../utils/routing"

interface Trip {
    id: string
    start: string
    end: string
    duration: string
    fee: string
    distance: string
    status: string
    inProgress: boolean
}

interface MainItem {
  id: string
  navId: string
  navScrollId: string
  data: Trip 
}

interface NavItem {
  id: string
  mainId: string
  label: string
}

interface MainItemQueryResult {
  id: string
  top: number
  dataset: {
    navId: string
    navScrollId: string
  }
}

// const licStatusMap = new Map([
//   [rental.v1.IdentityStatus.UNSUBMITTED, '未认证'],
//   [rental.v1.IdentityStatus.PENDING, '未认证'],
//   [rental.v1.IdentityStatus.VERIFIED, '已认证'],
// ])

// const licStatusMap = new Map([
//   [rental.v1.IdentityStatus.UNSUBMITTED, '未认证'],
//   [rental.v1.IdentityStatus.PENDING, '未认证'],
//   [rental.v1.IdentityStatus.VERIFIED, '已认证'],
// ])

Page({
  scrollStates: {
    mainItems: [] as MainItemQueryResult[] 
  },
  data: {
    avatarURL: '',
    mainItems: [] as MainItem[],
    navItems: [] as NavItem[],
    mainScroll: '',
    //licStatus: licStatusMap.get(rental.v1.IdentityStatus.UNSUBMITTED),
    licStatus: '未知',
    tripsHeight: 0, 
    navCount: 0,    
    navSel: '',
    navScroll: '',

    promotionItems: [
      {
          img: 'https://img.mukewang.com/5f7301d80001fdee18720764.jpg',
          promotionID: 1,
      },            
      {
          img: 'https://img.mukewang.com/5f6805710001326c18720764.jpg',
          promotionID: 2,
      },
      {
          img: 'https://img.mukewang.com/5f6173b400013d4718720764.jpg',
          promotionID: 3,
      },
      {
          img: 'https://img.mukewang.com/5f7141ad0001b36418720764.jpg',
          promotionID: 4,
      },
    ],
  },

  async onLoad() {
    this.populateTrips()  // 生成随机行程数据
    const userInfo = await getApp().globalData.userInfo
    this.setData({
      avatarURL: userInfo.avatarUrl,
    })
  },

  // 生成随机行程数据
  populateTrips() {
    const mainItems: MainItem[] = []
    const navItems: NavItem[] = []
    let navSel = ''
    let prevNav = ''
    for (let i = 0; i < 100; i++) {
      const mainId = 'main-' + i
      const navId = 'nav-' + i
      const tripId = (10001 + i).toString()
      if (!prevNav) {
        prevNav = navId
      } 
      mainItems.push({
        id: mainId,
        navId: navId,
        navScrollId: prevNav,
        data: {
          id: tripId,
          start: '北京故宫',
          end: '鸟巢',
          duration: '0时44分',
          fee: '56.00元',
          distance: '27.0公里',
          status: '已完成',
          inProgress: false,
        }
      })
      navItems.push({
        id: navId,
        mainId: mainId,
        label: tripId,
      })
      if (i === 0) {
        navSel = navId
      }
      prevNav = navId
    }
    this.setData({
      mainItems,
      navItems,
      navSel, 
    }, () => {
      this.prepareScrollStates()
    })
  },

  prepareScrollStates() {
    wx.createSelectorQuery().selectAll('.main-item')
      .fields({
        id: true,
        dataset: true,
        rect: true,
      }).exec(res => {
        this.scrollStates.mainItems = res[0]
      })
  },


  // 页面渲染完成后获取垂直导航组件的高度
  // 等于手机窗口高度 - 顶部轮播组件的高度 
  onReady() {
    wx.createSelectorQuery().select('#heading')
      .boundingClientRect(rect => {
        const height = wx.getSystemInfoSync().windowHeight - rect.height
        this.setData({
          tripsHeight: height,
          navCount: Math.round(height/50),
        })
      }).exec()
  },

  // 点击轮播图片后的处理函数
  onPromotionItemTap(e: any) {
    // 获取所点击的图片ID
    const promotionID = e.currentTarget.dataset.promotionId
    if (promotionID) {
      console.log(promotionID)
      // 跳转到图片对应的促销页面
    }
  },

  // 获取用户头像
  onChooseAvatar(e: any) {
    this.setData({
      avatarURL: e.detail.avatarUrl,
    }),
    // 将获取的avatarUrl保存到全局变量中
    getApp().globalData.userInfo.avatarUrl = e.detail.avatarUrl
  },

  // 跳转到注册页面
  onRegisterTap() {
    wx.navigateTo({
      // 无参跳转，则注册完后不会进入开锁页面
      //url: '/pages/register/register',
      url: routing.register(),
    })
  },

  // 点击左边导航栏时，右边行程列表跟着滚动
  onNavItemTap(e: any) {
    const mainId: string = e.currentTarget?.dataset?.mainId
    const navId: string = e.currentTarget?.id
    if (mainId && navId) {
      this.setData({
        mainScroll: mainId,
        navSel: navId,
      })
    }
  },

  // 右边滑动时，左边导航栏跟着滑动
  onMainScroll(e: any) {
    const top: number = e.currentTarget?.offsetTop + e.detail?.scrollTop
    if (top === undefined) { return }

    const selItem = this.scrollStates.mainItems.find(v => v.top >= top)
    if (!selItem) { return }
    
    this.setData({
      navSel: selItem.dataset.navId,
      navScroll: selItem.dataset.navScrollId,
    })
  },

  // 点击正在进行中的行程记录跳转到行程页面
  onMianItemTap(e: any) {
    if (!e.currentTarget.dataset.tripInProgress) {
        return
    }
    const tripId = e.currentTarget.dataset.tripId
    if (tripId) {
        wx.redirectTo({
            url: routing.drving({
                trip_id: tripId,
            }),
        })
    }
  },

})