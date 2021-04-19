<template>
  <div id="app">
    <!-- <datav /> -->
    <router-view></router-view>
  </div>
</template>

<script>
import { checkDB } from "@/api/initdb"
// import datav from './view/datav/index.vue'
import { getPVS_Base_LineList } from '@/api/PVS_Base_Line'

export default {
  name: 'app',
  async created(){
    await checkDB()
 
 // 不能在这里调用API, 因为有可能还未登录, 无法访问
            // const table = await getPVS_Base_LineList({ page: 1, pageSize: 100 })
            // if (table.code == 0) {
            //     this.$store.commit('update', ['lines', table.data.list]);
            // }
  },

  components: {
    // datav
  },
}
</script>

<style lang="scss">
// 引入初始化样式
@import '@/style/main.scss';
@import '@/style/base.scss';
@import '@/style/mobile.scss';
#app {
  background: #eee;
  height: 100vh;
  overflow: hidden;
}
</style>
