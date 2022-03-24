import { ProfileService } from "../../service/profile"
import { rental } from "../../service/proto_gen/rental/rental_pb"
import { routing } from "../../utils/routing"
function padString(n:number){
  return n<10? '0'+n.toFixed(0):n.toFixed(0)
}
function formatDate(millis: number){
  const dt = new Date(millis)
  const y = dt.getFullYear()
  const m = dt.getMonth() +1
  const d = dt.getDate()
  return `${padString(y)}-${padString(m)}-${padString(d)}`
}
// pages/register/register.ts
Page({
  redirectUrl:'',
  profileRefresher: 0,
  data: {
    licNo: '',
    name: '',
    genderIndex: 0,
    genders: ['未知', '男', '女'],
    birthDate: '1990-01-01',
    licImgURL: '',
    state: rental.v1.IdentityStatus[rental.v1.IdentityStatus.UNSUBMITTED],
  },
  /**
   * 用户上传驾驶证图片点击事件
   */
   onUploadLic() {
    console.log('onUploadLic')
    wx.chooseImage({
      success: res =>{
        if (res.tempFilePaths.length >0){
          this.setData({
            licImgURL:res.tempFilePaths[0]
          })
        }
      }
    })
  },
  onSubmit(){
    console.log(this.data.genderIndex)
    ProfileService.submitProfile({
      licNumber: this.data.licNo,
      name: this.data.name,
      gender: this.data.genderIndex,
      birthDateMillis: Date.parse(this.data.birthDate),
    }).then(p => {
      this.renderProfile(p)
      this.scheduleProfileRefresher()
    })
  },
  scheduleProfileRefresher(){
    this.profileRefresher = setInterval(()=>{
      ProfileService.getProfile().then(p =>{
        this.renderProfile(p)
        if(p.identityStatus!==rental.v1.IdentityStatus.PENDING){
          this.clearProfileRefresher()
        }
        if(p.identityStatus===rental.v1.IdentityStatus.VERIFIED){
          this.onLicVerified()
        }
      })
    },1000)
  },

  clearProfileRefresher(){
    if(this.profileRefresher){
      clearInterval(this.profileRefresher)
      this.profileRefresher=0
    }
  },
  onLicVerified(){
    if (this.redirectUrl){
      wx.redirectTo({
        url: this.redirectUrl,
      })
    }
    
  },
  onResubmit(){
    ProfileService.clearProfile().then(p =>this.renderProfile(p))
  },
  genderChange(e: any){
    this.setData({
      genderIndex: parseInt(e.detail.value)
    })
  },
  onBirthDateChange(e:any){
    this.setData({
      birthDate: e.detail.value
    })
  },

  renderProfile(p: rental.v1.IProfile){
    this.setData({
      licNo: p.identity?.licNumber || '',
      name: p.identity?.name || '',
      genderIndex: p.identity?.gender|| 0,
      birthDate: formatDate(p.identity?.birthDateMillis || 0),
      state: rental.v1.IdentityStatus[p.identityStatus||0],
    })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(opt: Record<'redirect',string>) {
    const o: routing.RegisterOpts = opt
    console.log(o.redirect)
    if (o.redirect){
      this.redirectUrl = decodeURIComponent(o.redirect)
      console.log(this.redirectUrl)
    }
    ProfileService.getProfile().then(p =>this.renderProfile(p))
  },
  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload() {
     this.clearProfileRefresher()
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