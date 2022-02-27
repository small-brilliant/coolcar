import camelcaseKeys from "camelcase-keys"
import { IAppOption } from "./appoption"
import { coolar } from "./service/proto_gen/trip_pb"

// app.ts
App<IAppOption>({
  globalData: {
    userInfo: wx.getStorageSync('storage_info') || undefined
  },
  onLaunch() {
    wx.request({
      url:"http://localhost:8080/trip/123",
      method: "GET",
      success: res =>{
        const getTripResponse = coolar.GetTripResponse.fromObject(camelcaseKeys(res.data as object,{
          deep:true,
        }))
        console.log(getTripResponse)
        console.log("status is", coolar.TripStatus[getTripResponse.trip?.status!])
      },
      fail: console.log,
    })
    // 展示本地存储能力
    const logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)
    // 登录
    wx.login({
      success: res => {
        console.log(res.code)
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
      }
    })
  }
})