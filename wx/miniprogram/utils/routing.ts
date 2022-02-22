export namespace routing {
    // driving
    export interface DrivingOpts{
        trip_id: string
    }
    export function driving(o: DrivingOpts){
        return `/pages/driving/driving?trip_id=${o.trip_id}`
    }
    // Lock
    export interface LockOpts{
        car_id: string
    }
    export function lock(o:LockOpts){
        return `/pages/lock/lock?car_id=${o.car_id}`
    }
    // register
    export interface RegisterOpts{
        redirect?: string
    }
    export interface RegisterParams{
        redirectURL: string
    }
    export function register(p?:RegisterParams){
        // `/pages/register/register?redirectUrl=${encodeURIComponent(redirectUrl)}`
        const page= '/pages/register/register'
        if(!p){
            return page
        }
        return `${page}?redirect=${encodeURIComponent(p.redirectURL)}`
    }
    // '/pages/mytrips/mytrips'
    export function mytrips(){
        return '/pages/mytrips/mytrips'
    }
}