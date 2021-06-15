import {get,post,setToken} from '../utils/request'


// 定义一个插件
export default{
    // 插件中必须包含install方法
    install: function(vue){
        // 给vue混入成员
        vue.mixin({
            methods:{
                $get(url,params){
                    return get(url,params)
                },
                $post(url,params){
                    return post(url,params)
                },
                $setToken(){
                    setToken()
                }
            },
        })
    }
}