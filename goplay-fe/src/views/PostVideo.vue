
<template>
  <div class="post-video">
    <h2>欢迎投稿：</h2>
    <el-form ref="form" :model="form" label-width="80px">
      <el-form-item label="标题">
        <el-input v-model="form.title"></el-input>
      </el-form-item>

      <el-form-item label="描述">
        <el-input type="textarea" v-model="form.info"></el-input>
      </el-form-item>

      <el-form-item>
        <el-button type="primary">立即创建</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>


<script>
export default {
  name: "PostVideo",
  data() {
    return {
      form: {
        title: "",
        info: ""
      }
    };
  },
  methods: {},
  components: {}
};
</script>