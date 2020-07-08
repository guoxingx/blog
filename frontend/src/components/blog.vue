<template>
  <div class="blog" style="margin-top: 50px;">
  <div class="blog-main" v-if="!notFound">
    <div class="html" id="html"></div>
    <div class="blog-main-footer">
      <h3 class="date">{{ this.blog.date.substring(0, 10) }}</h3>
    </div>
    <div id="commentsContainer"></div>
    <!-- <div id="commentsContainer"></div> -->
  </div>

  <div v-if="notFound">
    <h2>404</h2>
    <h3>Resource Not Found</h3>
    <el-link type="primary" @click="toIndex()">Back To index...</el-link>
  </div>

  </div>
</template>

<script>
import { getBlog, getHtml } from "@/js/requests"
import marked from 'marked'
import '@/assets/style/markdown.css'
import { gitmentConfig } from '@/js/gitmentConfig'
import Gitment from 'gitment'
import 'gitment/style/default.css'

export default {
  data () {
    return {
      blog: {},
      notFound: false
    }
  },
  created () {
    this.listBlog()
  },
  methods: {
    listBlog () {
      getBlog(this.$route.params.id).then(res => {
        if (res.data.code == 0) {
          this.blog = res.data.data

          getHtml(this.blog.path).then(res => {
            // this.html = res.data
            // alert((res.data))
            document.getElementById('html').innerHTML = marked(res.data)
          })
        } else {
          this.notFound = true
        }
      })
    },
    toIndex () {
      this.$router.push({ name: 'blogs' })
    },
    gitmentRender () {
      console.log(gitmentConfig)
      const gitment = new Gitment({
        id: this.$route.path,
        owner: gitmentConfig.owner,
        repo: gitmentConfig.repo,
        oauth: gitmentConfig.oauth
      })
      gitment.render('commentsContainer')
    }
  },
  watch: {
    '$route' (to, from) {
      console.log(to, from)
      this.gitmentRender()
      this.listBlog()
    }
  },
  beforeRouteEnter (to, from, next) {
    window.document.title = 'GXBlog'
    next(vm => { vm.gitmentRender() })
  },
  beforeRouteUpdate (to, from, next) {
    next()
  }
}
</script>

<style>
.gitment-comment-main {
  background: white;
}
.gitment-editor-main {
  background: white;
}
</style>
