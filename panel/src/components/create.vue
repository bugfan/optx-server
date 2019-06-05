<template>
    <el-form ref="form" :model="form" label-width="80px">
      <el-form-item label="问题" required>
        <el-input v-model="form.question" placeholder="输入问题"></el-input>
      </el-form-item>
      <el-form-item label="答案描述" >
        <el-input type="textarea" placeholder="输入问题答案的描述或者解答" v-model="form.desc"></el-input>
      </el-form-item>
       <el-form-item  label="选项" required>
          <el-input v-for="(v, k) in form.options" v-model="form.options[k]" placeholder="添加问题，可以添加多个" :key="k">
            <el-button v-if="k===0" slot="append" type="primary" @click="form.options.push('')"><i class="el-icon-plus"></i></el-button>
            <el-button v-if="k!=0" slot="append" type="primary" @click="form.options.splice(k, 1);"><i class="el-icon-delete"></i></el-button>
          </el-input>
        </el-form-item>
      <el-form-item label="答案选项" required>
        <el-input v-model="form.answer" placeholder="答案输入数字,按照选择顺序从0开始" ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">立即创建</el-button>
        <el-button  @click="reset">取消</el-button>
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
    reset() {
      this.form.question=''
      this.form.desc=''
      this.form.answer=''
      this.form.options=['']
    },
    onSubmit() {
       if (this.form.question==''){
            this.$message({
              type: 'info',
              message: '问题为空!'
            });
            return
       }
       if (this.form.options.length <2){
            this.$message({
              type: 'info',
              message: '请至少添加两个选项!'
            });
            return
       }
        if (this.form.answer ==''){
             this.$message({
              type: 'info',
              message: '请输入答案!'
            });
            return
       }
      api.create({
        question:this.form.question,
        desc:this.form.desc,
        answer:parseInt(this.form.answer),
        options:this.form.options,
        }).then(res =>{
          console.log(res.data)
          if (res.sataus >300){
            this.$message({
              type: 'info',
              message: res.data,
            });
          }
          if (res.status ==201){
            this.reset()
             this.$message({
              type: 'success',
              message: '添加成功',
            });
          }
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