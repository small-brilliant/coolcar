import { routing } from "../../utils/routing"

interface Trip {
  id: string
  start: string
  end: string
  duration: string
  fee: string
  distance: string
  status: string
}
interface MainItem{
  id: string
  navId: string
  navScrollId: string
  data: Trip
}
interface NavItem{
  id: string
  mainId: string
  label: string
}
interface MainItemQueryResult{
  id: string
  top: number
  dataset:{
    navId: string
    navScrollId: string
  }
}
// pages/mytrips/mytrips.ts
Page({

  /**
   * 页面的初始数据
   */
  scrollStates: {
    mainItems: [] as MainItemQueryResult[],
  },
  data: {
    indicatorDots: true,
    autoplay:true,
    interval:5000,
    duration:500,
    circular:false,
    vertical: false,
    mainScroll: "",
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
    ],
    mainItems: [] as MainItem[],
    navItems: [] as NavItem[],
    navCount: 0,
    navSel:'',
    navScroll: '',
    tripsHeight: 0,
  },

  onPromotoinItemTap(e: any){
    const promotionId = e.currentTarget.dataset.promotionId
    if (promotionId){

    }
  },
  onRegisterTap(){
     wx.navigateTo({
       url:routing.register()
     })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad() {
    this.populateTrips()
  },
  populateTrips(){
    const mainItems: MainItem[] = []
    const navItems: NavItem[] = []
    let navSel= ''
    let prevNav = ''
    for (let i = 0; i < 100 ; i++){
      if (!prevNav) {
        prevNav = 'nav-'+i
      }
        mainItems.push({
            id:'main-'+i,
            navId:'nav-'+i,
            navScrollId: prevNav,
            data: {
                id: (10001+i).toString(),
                start:'东方明珠',
                end: '迪士尼',
                distance: '22.0公里',
                duration:'0时52分',
                fee: '120.00元',
                status: '已完成',
    		    },
        })
        navItems.push({
            id: 'nav-'+i,
    		    mainId: 'main-'+i,
    		    label: (10001+i).toString(),
        })
        if (i===0){
          navSel = 'nav-'+i
        }
        prevNav = 'nav-'+i
    }
    this.setData({
        mainItems,
        navItems,
        navSel,
    },() =>{
      this.prepareScrollstatus()
    })
  },
  prepareScrollstatus(){
    wx.createSelectorQuery().selectAll('.main_items').fields({
      id: true,
      dataset:true,
      rect: true,
    }).exec(res =>{
      console.log(res)
      this.scrollStates.mainItems = res[0]
    })
  },
  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {
    wx.createSelectorQuery().select("#heading").boundingClientRect(rect =>{
      const height = wx.getSystemInfoSync().windowHeight - rect.height
      this.setData({
        tripsHeight: height,
        navCount: Math.round(height/50),
      })
    }).exec()
  },
  onNavItemTap(e: any){
    console.log("diandaol")
    const mainId: string = e.currentTarget?.dataset?.mainId
    const navId: string = e.currentTarget?.id
    if (mainId && navId){
      this.setData({
        mainScroll: mainId,
        navSel: navId,
      })
    }
  },
  onMainScroll(e: any){
    console.log(e)
    const top: number = e.currentTarget?.offsetTop + e.detail?.scrollTop
    if (top === undefined){
      return
    }
    const selItem = this.scrollStates.mainItems.find(v => v.top>=top)
    if (selItem === undefined){
      return
    }
    this.setData({
      navSel:selItem.dataset.navId,
      navScroll:selItem.dataset.navScrollId,
    })
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