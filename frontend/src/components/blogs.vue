<template>
  <div class="m-content-all">
    <h2>GX's Blog</h2>
    <h3>Index</h3>

    <div class="blogs">
      <blog-cell v-for="blog in blogsShow" :key="blog" :blog="blog"></blog-cell>
    </div>

    <el-pagination layout="prev,pager,next" small
      :total="this.blogs.length" :page-size="10"
      @current-change="changePage"
      v-if="this.blogs.length>10">
    </el-pagination>

  </div>
</template>
<script>
import { getBlogs } from "@/js/requests"

export default {
  name: 'Blogs',
  components: {
    'blog-cell': () => import('@/components/blog_cell.vue')
  },
  data () {
    return {
      date: new Date(),
      blogs: [],
      blogsShow: []
    }
  },
  watch: {
    '$route' (to, from) {
      console.log(to, from)
      this.listBlogs()
    },
  },
  created () {
    this.listBlogs()
  },
  methods: {
    listBlogs() {
      getBlogs().then(res => {
        if (res.data.code == 0 && res.data.data.length > 0) {
          this.blogs = res.data.data
          this.blogsShow = this.blogs.slice(0,10)
        }
      })
    },
    changePage(i) {
      this.blogsShow = this.blogs.slice((i-1)*10, i*10)
    },
  }
}
</script>

<style>
  h2{
    margin-top: 40px;
    margin-bottom: 20px;
  }
  h3{
    margin-bottom: 15px;
  }
  dd{
    height: 200px;
  }

  p>img{
    max-height: 200px;
    max-width: 200px;
  }
</style>
