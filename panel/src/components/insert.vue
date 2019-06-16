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
</template>

<script>
import api from '@/apis/option'

export default {
  data() {
    return {
      list:[],
      textarea:'',
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
    this.fetch()
  },
  methods: {
    fetch(){
      
    },
    getAnswerIndex (t) {
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
    deal() {
      this.list=[]
    var segs = this.textarea.split('.')
    for (let i =1 ;i<segs.length;i++){
      // console.log("ti:",segs[i])
      let ti = segs[i].split('。')
      if (ti.length > 0){
        let ques = ti[0]
        
        let answs = ti[1].split('\r')
        let answsItem = answs[0].split('）')
        let answers = []
        for (let j = 1;j<answsItem.length;j++){
            let startInd = answsItem[j].indexOf(' ')
            let item = answsItem[j].slice(0,startInd).replace(/^\s+|\s+$/g,"")  //replace(/ /g,'')
            answers.push(item)
        }
        let answer = ques.slice(ques.indexOf('（')+1,ques.indexOf('）')).replace(/\s+/g,"") //replace(/ /g,'')
        ques=ques.replace(new RegExp(answer,'g'),'')
        let obj = {
          Question:ques,
          Options:answers,
          Desc:'',
          Answer:this.getAnswerIndex(answer),
        }
        // console.log("ti 22222:",obj)
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