// pages/register/register.ts
Page({
  redirectUrl:'',
  data: {
    genders: ['未知','男','女','其他'],
    genderIndex: 0,
    birthDate:'1990-01-01',
    LicImgURL: '',
    state: 'UNSUBMITTED' as 'UNSUBMITTED' | 'PENDING' | 'VERIFIED',
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
            LicImgURL:res.tempFilePaths[0]
          })
        }
      }
    })
  },
  onSubmit(){
    this.setData({
      state: "PENDING",
    })
    setTimeout(() => {
      this.onLicVerified()
    },3000);
  },
  onLicVerified(){
    this.setData({
      state: "VERIFIED"
    })
    if (this.redirectUrl){
      wx.redirectTo({
        url: this.redirectUrl,
      })
    }
    
  },
  onResubmit(){
    this.setData({
      state: "UNSUBMITTED",
      LicImgURL:'',
    })
  },
  genderChange(e: any){
    this.setData({
      genderIndex: e.detail.value
    })
  },
  onBirthDateChange(e:any){
    this.setData({
      birthDate: e.detail.value
    })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(opt) {
    if (opt.redirectUrl){
      this.redirectUrl = decodeURIComponent(opt.redirectUrl)
    }
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