<template>
  <div>
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
    };
  },
  created: function() {
    this.fetch()
  },
  methods: {
    fetch(){
      api.all().then(res => {
        this.list=res.data
        console.log(this.list)
      })
    },
    editDialog(idx, data) {

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