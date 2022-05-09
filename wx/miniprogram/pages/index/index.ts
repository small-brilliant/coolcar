// index.ts
// 获取应用实例

import { IAppOption } from "../../appoption"
import { CarService } from "../../service/car"
import { ProfileService } from "../../service/profile"
import { rental } from "../../service/proto_gen/rental/rental_pb"
import { TripService } from "../../service/trip"
import { routing } from "../../utils/routing"

const app = getApp<IAppOption>()
const defaultAvatar="/resources/car.png"
const initialLat = 30
const initialLng = 120

interface Marker{
  iconPath: string,
  id: number,
  latitude: number,
  longitude: number,
  width: number,
  height: number,
}
Page({
  isPageShowing:false,
  socket: undefined as WechatMiniprogram.SocketTask | undefined,
  data: {
    scale: 16,
    latitude: initialLat,
    longitude: initialLng,
    isOverLooking: true,
    is3D: false,
    avatarUrl:'',
    markers: [] as Marker[],
    setting: {
      skew: 0,
      rotate: 0,
      showLocation: true,
      showScale: true,
      subKey: '',
      layerStyle: 1,
      enableZoom: true,
      enableScroll: true,
      enableRotate: false,
      showCompass: false,
      enable3D: false,
      enableOverlooking: false,
      enableSatellite: false,
      enableTraffic: false,
    },
  },
  onLoad(){
    if (app.globalData.userInfo) {
      this.setData({
        avatarUrl: app.globalData.userInfo.avatarUrl
      })
    }
  },
  onMyLocationTap(){
    wx.getLocation({
      type: 'gcj02',
      success: res=>{
        this.setData({
          latitude: res.latitude,
          longitude: res.longitude,
        })
      },
      fail: () =>{
        wx.showToast({
          icon: 'none',
          title: '前往设置页面授权',
        })
      }
    })
  },
  async onScanClicked(){
    const trips = await TripService.GetTrips(rental.v1.TripStatus.IN_PROGRESS)
    if ((trips.trips?.length || 0 )>0){
      await this.selectComponent('#tripModal').showModal()
      wx.navigateTo({
        url: routing.driving({
          trip_id:trips.trips![0].id!,
        })
      })
      return
    }
    //TODO: get car id from scan result
    wx.scanCode({
      success: async ()=>{
        const carId = '6278c168ec8f83dc2cf4f929'
        const lockURL = routing.lock({
          car_id: carId,
        })
        const profile = await ProfileService.getProfile()
        if ( profile.identityStatus === rental.v1.IdentityStatus.VERIFIED){
          wx.navigateTo({
            url: lockURL,
          })
        }else{
          await this.selectComponent('#licModal').showModal()
          wx.navigateTo({
            url: routing.register({
              redirectURL: lockURL,
            })
          })
        }
      },
      fail: console.error,
    })
  },
  onMyTrips(){
    wx.navigateTo({
      url:routing.mytrips()
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
 
  setupCarPosUpdater() {
    const map = wx.createMapContext("mapId")
    const markersByCarID = new Map<string, Marker>()
    let translationInProgress = false
    const endTranslation = () => {
      translationInProgress = false
    }
    this.socket = CarService.subscribe(car => {
      if (!car.id || translationInProgress || !this.isPageShowing) {
        console.log('dropped')
        return
      }
      const marker = markersByCarID.get(car.id)
      if (!marker) {
        // Insert new marker.
        const newMarker: Marker = {
          id: this.data.markers.length,
          iconPath: car.car?.driver?.avatarUrl || defaultAvatar,
          latitude: car.car?.position?.latitude || initialLat,
          longitude: car.car?.position?.longitude || initialLng,
          height: 20,
          width: 20,
        }
        markersByCarID.set(car.id, newMarker)
        this.data.markers.push(newMarker)
        translationInProgress = true
        this.setData({
          markers: this.data.markers,
        }, endTranslation)
        return
      }

      const newAvatar = car.car?.driver?.avatarUrl || defaultAvatar
      const newLat = car.car?.position?.latitude || initialLat
      const newLng = car.car?.position?.longitude || initialLng
      if (marker.iconPath !== newAvatar) {
        // Change iconPath and possibly position.
        marker.iconPath = newAvatar
        marker.latitude = newLat
        marker.longitude = newLng
        translationInProgress = true
        this.setData({
          markers: this.data.markers,
        }, endTranslation)
        return
      }

      if (marker.latitude !== newLat || marker.longitude !== newLng) {
        // Move marker.
        translationInProgress = true
        map.translateMarker({
          markerId: marker.id,
          destination: {
            latitude: newLat,
            longitude: newLng,
          },
          autoRotate: false,
          rotate: 0,
          duration: 900,
          animationEnd: endTranslation,
        })
      }
    })
  },
  onShow(){
    this.isPageShowing = true
    if(!this.socket){
      this.setupCarPosUpdater()
    }
  },
  onHide(){
    this.isPageShowing = false
    if (this.socket){
      this.socket.close({
        success: () =>{
          this.socket = undefined
        }
      })
    }
  },
})
