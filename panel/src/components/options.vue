<template>
    <div>
      <div v-show="show=='list'">
         <el-button class="filter-item" @click="newDialog()" type="primary" icon="el-icon-edit">添加</el-button>
           <br/>
           <el-select v-model="kind" placeholder="请选择题型">
            <el-option
              v-for="item in kinds"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
          <el-table :data="list" >
            <el-table-column align="center" label='ID' width="95">
              <template slot-scope="scope">
                {{scope.row.ID}}
              </template>
            </el-table-column>
            <el-table-column label="问题">
              <template slot-scope="scope">
                {{scope.row.Question}}
              </template>
            </el-table-column>
            <el-table-column label="选项">
              <template slot-scope="scope">
                {{scope.row.Options}}
              </template>
            </el-table-column>
            <el-table-column label="描述">
              <template slot-scope="scope">
                {{scope.row.Desc}}
              </template>
            </el-table-column>
            <el-table-column label="答案">
              <template slot-scope="scope">
                {{scope.row.Answer}}
              </template>
            </el-table-column>
            <el-table-column align="center" label="动作" width="270" class-name="small-padding fixed-width">
              <template slot-scope="scope">
                <el-button type="primary" size="mini" icon="el-icon-edit" @click="editDialog(scope.$index, scope.row)">编辑</el-button>
                <el-button type="danger" size="mini" icon="el-icon-delete" circle @click="deleteDialog(scope.$index,scope.row)"></el-button>
              </template>
            </el-table-column>
          </el-table>
      </div>
      <div v-if="show=='dialog'">
         <el-form ref="form" :model="form" label-width="80px">
          <el-form-item label="问题" required>
            <el-select v-model="kind" placeholder="请选择题型">
            <el-option
              v-for="item in kinds"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
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
            <el-button type="primary" @click="onSubmit">保存</el-button>
            <el-button  @click="cancel">取消</el-button>
          </el-form-item>
        </el-form>
      </div>
  </div>
</template>

<script>
import api from '@/apis/option'

export default {
  data() {
    return {
      list:[],
      show:'list',
      action:'add',
      form: {
          question: '',
          desc: '',  
          id: 0,
          answer: '',
          options:[''],
        },
      kind:0,
      kinds: [{
        value: 1,
        label: '判断题'
      }, {
        value: 0,
        label: '选择题'
      }],
    };
  },
  created: function() {
    this.fetch()
  },
  mounted() {
    this.fetch()
  },
  watch: {
    kind(){
      this.fetch()
    }
  },
  methods: {
    fetch(){
      api.all({kind:this.kind}).then(res => {
        this.list=res.data
        // console.log(this.list)
      })
    },
    editDialog(idx, data) {
      this.show='dialog'
      this.action='edit'
      this.form.answer=data.Answer
      this.form.question=data.Question
      this.form.desc=data.Desc
      this.form.id=data.ID
      this.form.options=data.Options
    },
    newDialog(){
      this.reset()
      this.show='dialog'
    },
    cancel() {
      this.show='list'
      // this.fetch()
      this.reset()
    },
    reset() {
      this.action='add'
      this.show='list'
      this.form.question=''
      this.form.id=0
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
      if (this.action=='add'){
          api.create({
            question:this.form.question,
            desc:this.form.desc,
            kind:this.kind,
            answer:parseInt(this.form.answer),
            options:this.form.options,
            }).then(res =>{
              // console.log(res.data)
              if (res.sataus >300){
                this.$message({
                  type: 'info',
                  message: res.data,
                });
              }
              if (res.status ==201){
                this.fetch()
                this.reset()
                this.$message({
                  type: 'success',
                  message: '添加成功',
                });
              }
          })  
      }else{
        api.update(this.form.id,{
            question:this.form.question,
            desc:this.form.desc,
            kind:this.kind,
            answer:parseInt(this.form.answer),
            options:this.form.options,
            }).then(res =>{
              // console.log(res.data)
              if (res.sataus >300){
                this.$message({
                  type: 'info',
                  message: res.data,
                });
              }
              if (res.status ==200){
                this.reset()
                this.$message({
                  type: 'success',
                  message: '保存成功',
                });
              }
          })  
      }
    },
     deleteDialog(index,data) {
      this.$confirm('此操作将删除该条目, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(()=>{
        return api.del(data.ID)
      }).then(() => {
        this.fetch()
        this.$message({
          type: 'success',
          message: '删除成功!'
        });
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });
      });
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

.add h2 {
  padding: 10px 0px;
  font-size: 12px;
}

.add span {
  margin-right: 5px;
}

.header {
  padding: 10px;
  border-bottom: 1px solid #eee;
}
</style>