// index.ts
// 获取应用实例

import { IAppOption } from "../../appoption"
import { ProfileService } from "../../service/profile"
import { rental } from "../../service/proto_gen/rental/rental_pb"
import { TripService } from "../../service/trip"
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
        const carId = 'car123'
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
  },
})
