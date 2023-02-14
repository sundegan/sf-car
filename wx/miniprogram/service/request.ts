import camelcaseKeys from "camelcase-keys"
import { auth } from "./proto_gen/auth/auth_pb"

export namespace sfcar {
    const serverAddr = 'http://localhost:8080'
    const AUTH_ERR = 'AUTH_ERR'

    const authData = {
        token: '',
        expireMs: 0,
    }

    // 定义请求参数
    export interface RequestOption<REQ, RES> {
        method: 'GET' | 'PUT' | 'POST' | 'DELETE'
        path: string
        data?: REQ
        respMarshaller: (r: object) => RES
    }

    
    export interface AuthOption {
        attachAuthHeader: boolean   // 用来控制是否携带AuthHeader
        retryOnAuthError: boolean   // 用来标记是否进行重试
    }

    // 将发送请求改写成Promise
    // 发送请求，返回RES类型的响应
    function sendRequest<REQ, RES>(opt: RequestOption<REQ, RES>, a: AuthOption): Promise<RES> {
        return new Promise((resolve, reject) => {
            // 定义请求头部
            const header:Record<string, any> = {}
            // 客户端进行Token验证,确定需要携带Token,且Token存在及没有过期,则给请求附上Token
            // 如果确定需要携带Token,但是Token无效,则报错
            if (a.attachAuthHeader) {
                if (authData.token && authData.expireMs >= Date.now()) {
                    header.authorization = 'Bearer ' + authData.token
                } else {
                    reject(AUTH_ERR)
                    return
                }
            } 
            wx.request({
                url: serverAddr + opt.path,
                method: opt.method,
                data: opt.data as WechatMiniprogram.IAnyObject,
                header,
                success: res => {
                    if (res.statusCode === 401) {
                        reject(AUTH_ERR)
                    } else if (res.statusCode >=  400) {
                        reject(res)
                    } else {
                        resolve(opt.respMarshaller(
                            camelcaseKeys(res.data as object, {
                                deep: true,
                            })
                        ))
                    }
                },
                fail: reject,
            })
        })
    }

    export async function sendRequestWithAuthRetry<REQ, RES>(opt: RequestOption<REQ, RES>, a?: AuthOption): Promise<RES> {
       // 如果没有a参数默认需要携带Token和进行重试
        const authOpt = a || {
            attachAuthHeader: true,
            retryOnAuthError: true,
       }
        // 客户端尝试进行登录重试
        try {
            await login()
            return sendRequest(opt, authOpt)
       } catch(err) {
            if (err === AUTH_ERR && authOpt.retryOnAuthError) {
                // 清除登录状态
                authData.token = ''
                authData.expireMs = 0
                // 进行登录重试
                return sendRequestWithAuthRetry(opt, {
                    attachAuthHeader: authOpt.attachAuthHeader,
                    retryOnAuthError: false,    // 进行重试后如果失败则不进行重试
                })
            } else {
                throw err
            }
       }
    }

    // 登录请求
    export async function login() {
        // 客户端对Token进行验证,如果Token有效则直接发送业务请求到服务器
        if (authData.token && authData.expireMs >= Date.now()) {
            
            return 
        }
        // 如果Token无效,则先进行登录
        const wxResp = await wxLogin()
        const reqTimeMs = Date.now()
        const resp = await sendRequest<auth.v1.ILoginRequest, auth.v1.ILoginResponse> ({
            method: 'POST',
            path: '/v1/auth/login',
            data: {
                code: wxResp.code,
            },
            respMarshaller: auth.v1.LoginResponse.fromObject,
        }, {
            attachAuthHeader: false,
            retryOnAuthError: false,
        })
        // 保存Token在内存中
        authData.token = resp.accessToken!
        authData.expireMs = reqTimeMs + resp.expiresIn! * 1000
    }

    // 将微信登录改写成Promise
    function wxLogin(): Promise<WechatMiniprogram.LoginSuccessCallbackResult> {
        return new Promise((resolve, reject) => {
            wx.login({
                success: resolve,
                fail: reject,
            })
        })
    }

    export interface UploadFileOpts {
        localPath: string
        url: string
    }
    
    // 驾照图片上传服务
    export function uploadfile(o: UploadFileOpts): Promise<void> {
        const data = wx.getFileSystemManager().readFileSync(o.localPath)
        return new Promise((resolve, reject) => {
            wx.request({
                method: 'PUT',
                url: o.url,
                data: data,
                success: res => {
                    if (res.statusCode >= 400) {
                        reject(res)
                    } else {
                        resolve()
                    }
                },
                fail: reject,
            })
        })
    }
}