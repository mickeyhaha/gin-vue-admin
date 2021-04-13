import { store } from '@/store/index'
//  获取字典方法 使用示例 getDict('sex').then(res)  或者 async函数下 const res = await getDict('sex')
export const getLines = async () => {
    await store.dispatch("lines/getLines")
    return store.getters["lines/getLines"]
}