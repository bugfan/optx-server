<template>
    <div>
      <el-input
        type="textarea"
        :rows="20"
        placeholder="请输入批量题目"
        v-model="textarea">
      </el-input>
      <el-button type="primary" @click="deal()" round>处理</el-button>
      <!-- 展示题目 -->
      <div v-show="show=='list'">
      <el-table :data="list" >
            <el-table-column align="center" label='ID' width="95">
              <template slot-scope="scope">
                {{scope.row.Qid}}
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
                <el-button type="primary" size="mini" icon="el-icon-edit" @click="piInsert()">批量插入题目</el-button>        
                <el-button type="primary" size="mini"  @click="resetInsert()">清空</el-button>        
        </div>
<!-- 修改狂 -->
        <div v-if="show=='dialog'">
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
            <el-button type="primary" @click="saveToList">保存</el-button>
            <el-button  @click="cancel">取消</el-button>
          </el-form-item>
        </el-form>
      </div>
  </div>
</template>

<script>
import api from '@/apis/insert'

export default {
  data() {
    return {
      list:[],
      textarea:'',
        show:'list',
      action:'add',
      form: {
          question: '',
          desc: '',  
          id: 0,
          answer: '',
          options:[''],
          //  上面的是每个题的属性
        },
    };
  },
  mounted: function() {
    // this.fetch()
  },
  methods: {
    saveToList () {
       this.list[this.form.id]={
         Answer:this.form.answer,
         Desc:this.form.desc,
         Options:this.form.options,
         Question:this.form.question
       }
       this.reset()
    },
    resetInsert() {
      this.list = []
    },
    piInsert () {
      if (this.list.length < 1){
        this.$message({
          type:"error",
          message:"还没有题目填入"
        })
      }else{
        var newList = []
        for (let i = 0;i<this.list.length;i++){
          if (this.list[i]){
            newList.push(this.list[i])
          }
        }
        
        api.create(newList).then(res => {
          console.log("res:",res.data)
          if (res.status==200){
            this.$message({
              type:"success",
              message:"题录入成功"
            })
            this.list=[]
          }
        }).error(res => {
          console.log("err:",res.data)
          this.$message({
            type:"error",
            message:res.data
          })
        })
      }
    },
    getAnswerIndex (t) {
     t=t.toUpperCase()
     if (t== 'A'){
       return 1
     }
     if (t== 'B'){
       return 2
     }
     if (t== 'C'){
       return 3
     }
     if (t== 'D'){
       return 4
     }
     if (t== 'E'){
       return 5
     }
      if (t== 'F'){
       return 6
     }
     return -1
    },
    editDialog(idx, data) {
      console.log("data:",idx,data)
      this.show='dialog'
      this.action='edit'
      this.form.answer=data.Answer
      this.form.question=data.Question
      this.form.desc=data.Desc
      this.form.id=idx
      this.form.options=data.Options
    },
    deleteDialog (idx, data){
       this.list[idx]=null
        let oldList = this.list
        this.list = []
        for (let i = 0;i<oldList.length;i++){
          if (oldList[i]){
            this.list.push(oldList[i])
          }
        }
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
    deal() {
      this.list=[]
    var segs = this.textarea.split(/[\d]+.[\s]+/)
    for (let i =1 ;i<segs.length;i++){
      if (segs[i].length<2){
        continue
      }
      // console.log("ti:",segs[i])
      let ti = segs[i].split('。')
      if (ti.length > 0){
        // 计算问题
        let ques = ti[0]
        let answer = ques.slice(ques.indexOf('（')+1,ques.indexOf('）')).replace(/\s+/g,"") //replace(/ /g,'')
        ques=ques.replace(new RegExp(answer,'g'),'')
        //计算答案选项
        let answs = ti[1].split('\r')
        let answers = []
        let answsItem = answs[0].split('）')
        if (answsItem.length==1){
          // 中文点
          let cans = answsItem[0].split(/[a-z]{1}|[A-Z]{1}、/)
          for (let k in cans){
            if (cans[k]=='↵'){
              continue
            }
            let cval = cans[k]
            let cval2=cval.replace(/^\s+|\s+$/g,"")
            if (cval2==''){
              continue
            }
            answers.push(cval2)
          }
        }else{
          // 括号包起来的
          for (let j = 1;j<answsItem.length;j++){
              let startInd = answsItem[j].indexOf(' ')
              let item = answsItem[j].slice(0,startInd).replace(/^\s+|\s+$/g,"")  //replace(/ /g,'')
              if (item.slice(item.length-1,item.length)=='（'){
                item=item.slice(0,item.length-1)
              }
              answers.push(item)
          }
        }
        
        let obj = {
          Qid:i,
          Question:ques,
          Options:answers,
          Desc:'',
          Answer:this.getAnswerIndex(answer),
        }
        // console.log("ti 22222:",obj)
        if (obj.Options.length<2){  //录入有问题，需要单独录入
          continue
        }
        this.list.push(obj)
      }
    }
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