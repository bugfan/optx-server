<template>
    <el-form ref="form" :model="form" label-width="80px">
      <el-form-item label="问题">
        <el-input v-model="form.question"></el-input>
      </el-form-item>
      <el-form-item label="答案">
        <el-input type="textarea" v-model="form.desc"></el-input>
      </el-form-item>
       <el-form-item  label="选项" required>
          <el-input v-for="(v, k) in form.options" v-model="form.options[k]"  :key="k">
            <el-button v-if="k===0" slot="append" type="primary" @click="form.options.push('')"><i class="el-icon-plus"></i></el-button>
            <el-button v-if="k!=0" slot="append" type="primary" @click="form.options.splice(k, 1);"><i class="el-icon-delete"></i></el-button>
          </el-input>
        </el-form-item>
      <el-form-item label="答案选项">
        <el-input v-model="form.answer"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">立即创建</el-button>
        <el-button>取消</el-button>
      </el-form-item>
    </el-form>
</template>

<script>

import api from '@/apis/option'

export default {
  data() {
    return {
      form: {
          question: '',
          desc: '',  
          answer: '',
          options:[''],
        },
    };
  },
  created: function() {

  },
  methods: {
    onSubmit() {
      api.create({
        question:this.form.question,
        desc:this.form.desc,
        answer:parseInt(this.form.answer),
        options:this.form.options,
        }).then(res =>{
          console.log(res.data)
      })  
    },
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#table {
  padding: 0 10px;
}

.add {
  border: 1px solid #ddd;
  margin: 10px 0;
  padding: 15px;
}

.header {
  padding: 10px;
  border-bottom: 1px solid #eee;
}
</style>