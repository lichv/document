
<template>
	<div class="markdown-display">
		<div class="markdown-header">
			<el-button @click="drawer = true" type="primary" class="markdown-header-handle-btn">管理图片</el-button>
		</div>
		<div class="markdown-content">
			<div class="markdown-side">
				<el-menu router>
					<template v-for="item in items">
						<el-menu-item :index="'/item/'+item.filename">{{item.filename}}</el-menu-item>
					</template>
					
				</el-menu>
			</div>
			<div class="markdown-main">
				<div v-html="html" class="markdown-body"></div>
			</div>
		</div>
		<el-drawer v-model="drawer" :direction="direction" size="60%">
			<el-manage-oss></el-manage-oss>
		</el-drawer>
	</div>
</template>
<script>
	import {watch} from 'vue'
	import api from '/@/api';
	import marked from 'marked'
	export default {
		name: 'ElMarkdownDisplay',
		data() {
			return {
				drawer: false,
				direction: 'rtl',
				client:null,
				imageUrl: '',
				items:[],
				html:'',
			}
		},
		props: {
			stsUrl: {
				type: String,
				default: "/api/aliyun/sts"
			},
			region: {
				type: String,
				default: "oss-cn-shanghai"
			},
			bucket: {
				type: String,
				default: "easytc"
			},
			rootname: {
				type: String,
				default: "aegicare"
			},
			dirname: {
				type: String,
				default: "picture"
			},
		},
		watch:{
			$router:{
				handler(val,oldval){
					console.log('val',val)
					console.log('oldval',oldval)
				},
			}
		},

		mounted() {
			let _this = this;
			this.getFiles()
			this.readFile()
			this.$watch(
				() => this.$route.params,
				(toParams, previousParams) => {

					console.log('toParams',toParams)
					console.log('previousParams',previousParams)
					console.log('this.$route.params.pathMatch',_this.$route.params.pathMatch)
					_this.readFile()
				}
				)
		},
		methods:{
			getFiles:function(){
				let _this = this;
				api.getMarkdownFiles.send()
				.then(result => {
					console.log('result',result)
					if (result.state==2000) {
						_this.items = result.data
					}
				})
				.catch(e => {
					console.log(e)
				})
			},
			readFile:function(){
				let _this = this;
				var pathMatch = this.$route.params.pathMatch
				console.log('pathMatch',pathMatch)
				if (typeof(pathMatch)!="undefined") {
					var path = pathMatch.join('/')
					api.readMarkdown.send({"path":path})
					.then(result => {
						console.log('result',result)
						if (result.state==2000) {
							_this.html = marked(result.data)
						}
					})
					.catch(e => {
						console.log(e)
					})
				}else{
					console.log('没进入')
				}
			}
		},
	}
</script>
<style>
.markdown-display{
	height: 100%;
	width: 100%;
	padding: 0;
	margin: 0;
}
.markdown-header{
	height: 50px;
	border-bottom: 1px solid #666;
}
.markdown-header-handle-btn{
	margin-right: 16px;
	float: right;
	margin-top: 8px;
}
.markdown-content{
	height: calc(100% - 51px);
	display: flex;
}
.markdown-side{
	width: 320px;
	height: 100%;
	overflow-y: scroll;
	background-color: #fff;
}
.markdown-main{
	padding: 20px;
	height: 100%;
	overflow-y: scroll;
	flex: 1;
}
</style>