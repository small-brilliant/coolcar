import { IAppOption } from "../../appoption"
import { CarService } from "../../service/car"
import { car } from "../../service/proto_gen/car/car_pb"
import { TripService } from "../../service/trip"
import { routing } from "../../utils/routing"

// pages/lock/lock.ts
Page({
  /**
   * 页面的初始数据
   */
  carID: "",
  carRefresher: 0,
  data: {
    avatarUrl:'',
  },
  onLoad(opt: Record<'car_id',string>){
    const o: routing.LockOpts = opt
    this.carID = o.car_id
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
      success: async loc => {
        if (!this.carID) {
            console.error('no carID specified')
            return
        }
        const trip = await TripService.CreateTrip({
          carId: this.carID,
          start: {
            latitude:loc.latitude,
            longitude:loc.longitude,
          },
          avatarUrl: this.data.avatarUrl,
        })
        wx.showLoading({
          title:'开锁中',
          mask: true,
        })
        this.carRefresher = setInterval(async() =>{
          const c = await CarService.getCar(this.carID)
          if(c.status === car.v1.CarStatus.UNLOCKED){
            this.clearRefresher()
            wx.redirectTo({
              // url:`/pages/driving/driving?trip-id=${tripId}`,
              url: routing.driving({
                trip_id: trip.id!
              }),
              complete: ()=>{
                wx.hideLoading()
              }
            })
          }
        },2000)
        if (!trip.id){
          console.error(trip)
          console.error("no tripID in response",trip)
          return
        }
      },
      fail: ()=>{
        wx.showToast({
          icon:"none",
          title:'前往右上角设置界面授权获取位置信息',
        })
      }
    })
  },
  clearRefresher(){
    if (this.carRefresher){
      clearInterval(this.carRefresher)
      this.carRefresher = 0
    }
  },
  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload() {
    this.clearRefresher()
    wx.hideLoading
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