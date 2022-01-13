// pages/mytrips/mytrips.ts
Page({

  /**
   * 页面的初始数据
   */
  data: {
    indicatorDots: true,
    autoplay:true,
    interval:5000,
    duration:500,
    circular:false,
    vertical: false,
    promotionItems: [
      {
        promotionID:1,
        img:'https://img1.sycdn.imooc.com/61cd26c80001889817920764.jpg',
      },
      {
        promotionID:2,
        img:'https://img2.sycdn.imooc.com/61db97190001c18817920764.jpg',
      },
      {
        promotionID:3,
        img:'https://img2.sycdn.imooc.com/61df86b10001e3d217920764.jpg',
      }
    ]
  },

  onPromotoinItemTap(e: any){
    const promotionId = e.currentTarget.dataset.promotionId
    if (promotionId){

    }
  },
  onRegisterTap(){
     wx.navigateTo({
       url:'/pages/register/register'
     })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad() {

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