<template>
  <div id="app">
    <!-- <datav /> -->
    <router-view></router-view>
  </div>
</template>

<script>
import { checkDB } from "@/api/initdb"
import datav from './view/datav/index.vue'

export default {
  name: 'app',
  async created(){
    await checkDB()
  },

  components: {
    datav
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
