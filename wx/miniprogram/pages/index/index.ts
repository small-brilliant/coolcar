// index.ts
// 获取应用实例

import { routing } from "../../utils/routing"

const app = getApp<IAppOption>()
const img = '/resources/car.jpg'

Page({
  data: {
    scale: 16,
    latitude: 39.92,
    longitude: 116.46,
    isOverLooking: true,
    is3D: false,
    avatarUrl:'',
    markers:[
      {
        id: 0,
        iconPath: img,
        width: 50,
        height: 50,
        latitude: 26.099994,
        longitude: 113.324520,
      },
      {
        id: 1,
        iconPath: img,
        width: 50,
        height: 50,
        latitude: 23.099994,
        longitude: 113.324520,
      },
      {
        id: 2,
        iconPath: img,
        width: 20,
        height: 20,
        latitude: 29.756825521115363,
        longitude: 121.87222114786053,
      }
    ],
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
  onScanClicked(){
    //TODO: get car id from scan result
    const carId = 'car123'
    const redirectUrl = routing.lock({
      car_id: carId,
    })
    console.log(redirectUrl)

    wx.scanCode({
      success: ()=>{
        wx.navigateTo({
          url: routing.register({
            redirectURL: redirectUrl,
          })
        })
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
  moveCars() {
    const map = wx.createMapContext("mapId")
    const dest = {
      latitude: 26.099994,
      longitude: 113.324520,
    }
    const movecar = () =>{
      dest.latitude += 0.1,
      dest.longitude += 0.1,
      map.translateMarker({
        destination: {
          latitude: dest.latitude,
          longitude: dest.longitude,
        },
        markerId: 0,
        autoRotate: false,
        rotate: 0,
        duration: 5000,
        animationEnd: movecar,
      })
    }
    movecar()
  }
})
