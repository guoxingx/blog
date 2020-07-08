<template>
  <section class="g-layout-all">
      <el-row>
        <el-col :xs="0" :sm="2" :md="4" :lg="4" style="color: white">
          fixed
        </el-col>

        <el-col :xs="24" :sm="20" :md="16" :lg="16">
          <el-row>
            <el-col :xs="24" :sm="20">
              <section>
                <transition name="fade" mode="out-in">
                  <router-view :key="key" />
                </transition>
              </section>
            </el-col>

            <el-col :xs="24" :sm="4">
              <section>
                <infos/>
              </section>
            </el-col>
          </el-row>
        </el-col>
      </el-row>

  </section>
</template>
<script>
export default {
  name: 'Layout',
  components: {
    'infos': () => import('@/components/infos.vue')
  },
  computed: {
    key () {
      return this.$route.fullPath
    }
  },
  created () {

  },
  methods: {

  }
}
</script>
