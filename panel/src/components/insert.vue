<template>
    <div>
      
  </div>
</template>

<script>
import api from '@/apis/option'

export default {
  data() {
    return {
      list:[],
      form: {
          question: '',
          desc: '',  
          id: 0,
          answer: '',
          options:[''],
        },
    };
  },
  mounted: function() {
    this.fetch()
  },
  methods: {
    fetch(){
      
    }
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