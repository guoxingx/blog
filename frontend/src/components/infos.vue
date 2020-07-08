<template>

  <div>
    <div class="info">
      <h3 class="links-head">Links</h3>
      <div><el-link type="primary" href="https://github.com/guoxingx">My Github</el-link></div>
      <div><el-link type="primary" href="mailto:3064793912@qq.com">Email me</el-link></div>

      <h3 class="links-head">Blog Archive</h3>
      <div>
        <el-link type="primary" @click="toIndex()">To index</el-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {

    }
  },
  methods: {
    toIndex () {
      this.$router.push({ name: 'blogs' })
    }
  }
}
</script>

<style>
.info {
  margin-top: 60px;
}
.links-head {
  margin-top: 20px;
}
</style>
