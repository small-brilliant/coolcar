// index.ts
// 获取应用实例

import { getUserProfile } from "../../utils/util"

const app = getApp<IAppOption>()
const img = '/resources/car.jpg'

Page({
  data: {
    scale: 20,
    latitude: 26.099994,
    longitude: 113.324520,
    isOverLooking: true,
    is3D: false,
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
    wx.scanCode({
      success: ()=>{
        wx.navigateTo({
          url: '/pages/register/register'
        })
      },
      fail: console.error,
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
