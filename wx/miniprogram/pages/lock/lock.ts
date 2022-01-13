// pages/lock/lock.ts
Page({
  /**
   * 页面的初始数据
   */
  data: {
    avatarUrl:'',
  },
  onLoad(opt){
    console.log("current carId:",opt.car_id)
    const app = getApp<IAppOption>()
    console.log(app.globalData.userInfo)
    if (app.globalData.userInfo) {
      this.setData({
        avatarUrl: app.globalData.userInfo.avatarUrl
      })
    }
  },
  onUnlockTap(){
    wx.getLocation({
      type: 'gcj02',
      success: loc =>{
        console.log('starting a trip',{
          location:{
            latitude:loc.latitude,
            longitude:loc.longitude
          }
        })
        const tripId = 'trip456'
        wx.showLoading({
          title:'开锁中',
          mask: true,
        })

        setTimeout(() => {
          wx.redirectTo({
            url:`/pages/driving/driving?trip-id=${tripId}`,
            complete: ()=>{
              wx.hideLoading()
            }
          })
        },2000)

      },
      fail: ()=>{
        wx.showToast({
          icon:"none",
          title:'前往右上角设置界面授权获取位置信息',
        })
      }
    })
  },
  onGetUserInfo(){
    const app = getApp<IAppOption>()
    console.log('onGetUserInfo')
    wx.getUserProfile({
      desc: '展示用户信息', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
      success: (res) => {
        console.log(res)
        app.globalData.userInfo = res.userInfo
        this.setData({
          avatarUrl: res.userInfo.avatarUrl
        })
        wx.setStorageSync('storage_info', res.userInfo)
      },
      fail:(res)=>{
        // debugger
        console.log(res)
      }
    })
  },
  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow() {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide() {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload() {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh() {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom() {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage() {

  }
})