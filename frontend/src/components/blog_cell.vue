<template>
  <div class="blog-unit">
    <div>
      <div style="float: left; margin-right: 30px;">
        <img :src="blog.cover" @click="toBlog(1)">
      </div>
      <div style="text-align: left;">
        <div>
        <el-link type="primary" :underline="false" @click="toBlog(1)">
          <h3>{{ blog.title }}</h3>
        </el-link>
        </div>
        <div style="">
          <p>{{ blog.summary }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ['blog'],
  methods: {
    toBlog () {
      this.$router.push({ name: 'blog', params: { id: this.blog.id } })
    }
  }
}
</script>

<style>
.blog-unit {
  display: inline-block;
  margin-top: 30px;
  width: 95%;
  /*height: 200px;*/
  /*max-height: 200px;*/

  padding-top: 20px;
  border-bottom: 0.5px solid #BACCE3;
}

.blog-unit img {
  margin: 0 auto;
  max-height: 120px;
  margin-bottom: 5px;
  /*width: 300px;*/
  /*max-width: 100%;*/
  /*max-width: 800px;*/
  vertical-align: middle;
}

.blog-unit .tag {
  margin-bottom: 10px;
  margin-left: 10px;
}

.blog-unit a {
  text-decoration: none;
}

.blog-unit-left {
  float: left;
}
.blog-unit-text {
  text-align: left;
}

</style>
