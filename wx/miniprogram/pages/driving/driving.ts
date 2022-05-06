import { ProfileService } from "../../service/profile"
import { TripService } from "../../service/trip"
import { routing } from "../../utils/routing"


const centPerSec = 0.7

function formatDuration(sec: number){
  const padString = (n:number)=>
    n<10? '0'+n.toFixed(0) : n.toFixed(0)
  
  const h = Math.floor(sec/3600)
  sec -= h*3600
  const m = Math.floor(sec/60)
  sec -= m*60
  const s = Math.floor(sec)
  return `${padString(h)}:${padString(m)}:${padString(s)}`  
}
function formatFee(cents: number){
  return (cents / 100).toFixed(2)
}


Page({

  /**
   * 页面的初始数据
   */
  tripID: '',
  timer: undefined as undefined|number,
  data: {
    location: {
      latitude: 39.92,
      longitude: 116.46,
    },
    scale: 14,
    elapsed: '00:00:00',
    fee: '0.00'
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(opt: Record<'trip_id',string>) {
    const o: routing.DrivingOpts = opt
    this.tripID = o.trip_id
    TripService.GetTrip(o.trip_id).then(console.log)
    this.setupLocationUpdator(),
    this.setupTimer()
  },
  onUnload(){
    wx.stopLocationUpdate()
    if (this.timer)[
      clearInterval(this.timer)
    ]
  },
  setupLocationUpdator(){
    wx.startLocationUpdate({
      fail: console.error
    })
    wx.onLocationChange(loc =>{
      console.log("location: ",loc)
      this.setData({
        latitude: loc.latitude,
        longitude: loc.longitude,
      })
    })
  },

  /**
   * 计费
   */
  setupTimer(){
    let elapsedSec = 0
    let cents = 0
    this.timer = setInterval(()=>{
      elapsedSec++
      cents += centPerSec
      this.setData({
        elapsed: formatDuration(elapsedSec),
        fee: formatFee(cents),
      })
    },1000)
  },
  EndTrip(){
    TripService.finishTrip(this.tripID).then(() => {
      wx.redirectTo({
          url: routing.mytrips(),
      })
    }).catch(err => {
        console.error(err)
        wx.showToast({
            title: '结束行程失败',
            icon: 'none',
        })
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