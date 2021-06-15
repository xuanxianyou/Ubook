// 请求相关的方法
import axios from  'axios'
// 导入加载进度条
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'

// 初始化axios对象
var spider = axios.create({
    // 接口URL
    baseURL : 'http://localhost:9000',
    // 超时时间
    timeout: 5000,
})

// 创建一个GET请求
let get = async function(url,params){
    let {data} = await spider.get(url,{params});
    return data;
}
// 创建一个post请求
let post = async function(url,params){
    let {data} = await spider.post(url,params);
    return data;
}
// 添加
let setToken = function(){
    axios.defaults.headers.common['token'] = sessionStorage.getItem('token')
}
// 添加请求拦截器
spider.interceptors.request.use(
    function(config){
        // 发送请求之前做的事情
        NProgress.start()
        return config;
    },function(error){
        // 请求错误处理
        NProgress.done()
        return Promise.reject(error);
});
// 添加响应拦截器
spider.interceptors.response.use(
    function(response){
        // 响应数据处理
        NProgress.done()
        return response;
    },function(error){
        // 响应错误处理
        NProgress.done()
        return Promise.reject(error);
})

export{
    get,
    post,
    setToken,
}