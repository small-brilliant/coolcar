import { ProfileService } from "../../service/profile"
import { rental } from "../../service/proto_gen/rental/rental_pb"
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
const updateIntervalSec = 5
const initialLat = 30
const initialLng = 120
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
    fee: '0.00',
    markers: [
      {
          iconPath: "/resources/car.png",
          id: 0,
          latitude: 39.92,
          longitude: 116.46,
          width: 20,
          height: 20,
      },
  ],
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(opt: Record<'trip_id',string>) {
    const o: routing.DrivingOpts = opt
    this.tripID = o.trip_id
    TripService.GetTrip(o.trip_id).then(console.log)
    this.setupLocationUpdator(),
    this.setupTimer(o.trip_id)
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
  async setupTimer(tripID: string){
    const trip = await TripService.GetTrip(tripID)
    if (trip.status !== rental.v1.TripStatus.IN_PROGRESS) {
      console.error('trip not in progress')
      return
    }
    let secSinceLastUpdate = 0
    let lastUpdateDurationSec = trip.current!.timestampSec! - trip.start!.timestampSec!
    const toLocation = (trip: rental.v1.ITrip) => ({
      latitude: trip.current?.location?.latitude || initialLat,
      longitude: trip.current?.location?.longitude || initialLng,
    })
    const location = toLocation(trip)
    this.data.markers[0].latitude = location.latitude
    this.data.markers[0].longitude = location.longitude
    this.setData({
        elapsed: formatDuration(lastUpdateDurationSec),
        fee: formatFee(trip.current!.feeCent!),
        location,
        markers: this.data.markers,
    })
    this.timer = setInterval(() => {
      secSinceLastUpdate++
      if (secSinceLastUpdate % updateIntervalSec === 0) {
          TripService.GetTrip(tripID).then(trip => {
              lastUpdateDurationSec = trip.current!.timestampSec! - trip.start!.timestampSec!
              secSinceLastUpdate = 0
              const location = toLocation(trip)
              this.data.markers[0].latitude = location.latitude
              this.data.markers[0].longitude = location.longitude
              this.setData({
                  fee: formatFee(trip.current!.feeCent!),
                  location,
                  markers: this.data.markers,
              })
          }).catch(console.error)
      }
      this.setData({
          elapsed: formatDuration(lastUpdateDurationSec + secSinceLastUpdate),
      })
  }, 1000)
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