<template>
  <el-container class="layout-container">
    <el-header class="layout-header" height="50px">
      <div class="header-left">
        <img class="header-logo" src="https://www.aegicare.com/static/new/images/logo_new.png">
      </div>
      <div class="header-right">
        
      </div>
    </el-header>
      <el-container class="layout-body">
          <el-aside class="layout-aside">
            <el-menu :default-active="'README'" mode="vertical" @select="handleSelect">
                <el-menu-item v-for='(value,key) in mdFiles' :index='key' :key='key'>{{key}}</el-menu-item>
            </el-menu>
          </el-aside>
          <el-main class="layout-main">
              <div v-html="replacedMd"></div>
          </el-main>
      </el-container>
  </el-container>
</template>

<script>

const fileContext = require.context('@/../docs', true, /.md$/)
console.log(fileContext)
var mdFiles = {}
fileContext.keys().forEach(key=>{
  console.log('__inner',key)
    mdFiles[key.replace('./','').replace('.md','')] = fileContext(key)
})
export default {
    data(){return{
        md:'',
        url_mapping_table:{
            'backend_admin_url':'http://119.23.182.57:8020/admin',
        },
        mdFiles: mdFiles,
    }},
    methods:{
        handleSelect(index, indexPath){
            this.md = this.mdFiles[index]
        }
    },
    computed:{
        replacedMd(){
            let resultMd = this.md
            return resultMd
        }
    },
    created(){
        this.md = this.mdFiles['README']
    }
}
</script>

<style>
.layout-header{
  border-bottom: 1px solid #dedede;
  background-color: #0a1e3c;
}
.header-logo{
  height: 48px;
}
table{
  width: 100%;
}
tr, td {
      border: 1px solid #eee;
}
</style>